/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package manager

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-webconsole/console/types"

	"github.com/gorilla/websocket"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/remotecommand"
)

// GuideMessages is a response string
var GuideMessages = []string{
	"###########################################################################################\r\n",
	"#                                 Welcome To BCS Console                                  #\r\n",
	"###########################################################################################\r\n",
}

//DefaultCommand 默认命令, 可以优先使用bash, 如果没有, 回退到sh
var DefaultCommand = []string{
	"/bin/sh",
	"-c",
	"TERM=xterm-256color; export TERM; [ -x /bin/bash ] && (" +
		"[ -x /usr/bin/script ] && /usr/bin/script -q -c \"/bin/bash\" /dev/null || exec /bin/bash) || exec /bin/sh",
}

const (
	writeWait  = 10 * time.Second
	pongWait   = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10

	// InputLineBreaker 输入分行标识
	InputLineBreaker = "\r"
	// OutputLineBreaker 输出分行标识
	OutputLineBreaker = "\r\n"

	// AnsiEscape bash 颜色标识
	AnsiEscape = "r\"\\x1B\\[[0-?]*[ -/]*[@-~]\""

	queueName = "bcs_web_console_record"
	tags      = "bcs-web-console"

	StdinChannel  = "0"
	StdoutChannel = "1"
	StderrChannel = "2"
	ErrorChannel  = "3"
	ResizeChannel = "4"

	// 审计上报、ws连接监测时间间隔
	recordInterval = 10
)

type errMsg struct {
	Msg string `json:"msg,omitempty"`
}

// WsMessage websocket消息
type WsMessage struct {
	MessageType int
	Data        types.XtermMessage
}

// ssh流式处理器
type streamHandler struct {
	wsConn      *wsConn
	resizeEvent chan remotecommand.TerminalSize
}

type wsConn struct {
	conn          *websocket.Conn
	inChan        chan *WsMessage // 读取队列
	outChan       chan *WsMessage // 发送队列
	mutex         sync.Mutex      // 避免重复关闭管道
	isClosed      bool
	closeChan     chan struct{} // 关闭通知
	ConnTime      time.Time     // 连接时间
	LastInputTime time.Time
	PodName       string //
	ConfigMapName string
	Username      string //
	SessionID     string
	Project       string
	Cluster       string
	InputRecord   string // 输入
	OutputRecord  string // 输出
	Context       interface{}
}

func genWsConn(conn *websocket.Conn, conf types.WebSocketConfig) *wsConn {
	configMapName := getConfigMapName(conf.ClusterID, conf.ProjectsID)
	podName := getPodName(conf.ClusterID, conf.ProjectsID)
	return &wsConn{
		conn:          conn,
		inChan:        make(chan *WsMessage, 1000),
		outChan:       make(chan *WsMessage, 1000),
		isClosed:      false,
		closeChan:     make(chan struct{}),
		ConnTime:      time.Now(),
		LastInputTime: time.Now(),
		PodName:       podName,
		ConfigMapName: configMapName,
		Username:      conf.User,
		SessionID:     conf.SessionID,
		Project:       conf.ProjectsID,
		Cluster:       conf.ClusterID,
	}
}

// 读取协程
func (c *wsConn) wsReadLoop() {
	defer c.wsClose()
	for {
		// 读一条message
		msgType, data, err := c.conn.ReadMessage()
		if err != nil {
			return
		}

		// 解析base64数据
		dataDec, err := base64.StdEncoding.DecodeString(string(data[1:]))
		if err != nil {
			continue
		}

		// 解析数据
		wsMessage := WsMessage{
			MessageType: msgType,
		}
		xtermMsg := types.XtermMessage{}
		if string(data[0]) == ResizeChannel {
			wsMessage.Data.MsgType = "resize"
			err = json.Unmarshal(dataDec, &xtermMsg)
			if err != nil {
				continue
			}
			wsMessage.Data.Rows = xtermMsg.Rows
			wsMessage.Data.Cols = xtermMsg.Cols
		} else {
			wsMessage.Data.MsgType = "input"
			wsMessage.Data.Input = string(dataDec)
			// 把输入存起来
			c.InputRecord += string(dataDec)
		}

		// 更新ws时间
		c.LastInputTime = time.Now()

		// 放入请求队列
		c.inChan <- &wsMessage
	}
}

// 发送协程
func (c *wsConn) wsWriteLoop() {
	// 服务端返回给页面的数据
	for {
		select {
		// 取一个应答
		case msg := <-c.outChan:
			// 写给web websocket
			output := base64.StdEncoding.EncodeToString([]byte(msg.Data.Output))
			if err := c.conn.WriteMessage(msg.MessageType, []byte(output)); err != nil {
				break
			}
			c.OutputRecord += msg.Data.Output
		case <-c.closeChan:
			c.wsClose()
		}
	}
}

// 关闭连接
func (c *wsConn) wsClose() {
	c.conn.Close()
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if !c.isClosed {
		c.isClosed = true
		close(c.closeChan)
	}
}

// 发送返回消息到协程
func (c *wsConn) wsWrite(messageType int, data []byte) (err error) {

	select {
	case c.outChan <- &WsMessage{
		MessageType: messageType,
		Data: types.XtermMessage{
			Output: string(data),
		},
	}:

	case <-c.closeChan:
		err = errors.New("WsWrite websocket closed")
		break
	}
	return
}

func (c *wsConn) WsRead() (msg *WsMessage, err error) {

	select {
	case msg = <-c.inChan:
		return
	case <-c.closeChan:
		err = errors.New("WsRead websocket closed")
		break
	}

	return
}

func (c *wsConn) periodicTick(period time.Duration) {

	go wait.NonSlidingUntil(c.tickTimeout, period*time.Second, c.closeChan)
}

// 主动停止掉session
func (c *wsConn) tickTimeout() {
	nowTime := time.Now()
	idleTime := nowTime.Sub(c.LastInputTime).Seconds()
	if idleTime > TickTimeout {
		// BCS Console 已经分钟无操作
		msg := fmt.Sprintf("BCS Console 已经 %d 分钟无操作", TickTimeout/60)
		blog.Info("tick timeout, close session %s, idle time, %.2f", c.PodName, idleTime)
		c.inChan <- &WsMessage{
			MessageType: websocket.TextMessage,
			Data:        types.XtermMessage{Output: string(msg)},
		}
		c.wsClose()
		return
	}
	loginTime := nowTime.Sub(c.ConnTime).Seconds()
	if loginTime > LoginTimeout {
		// BCS Console 使用已经超过{}小时，请重新登录
		msg := fmt.Sprintf("BCS Console 使用已经超过 %d 小时，请重新登录", LoginTimeout/60)
		blog.Info("tick timeout, close session %s, login time, %.2f", c.PodName, loginTime)
		c.wsClose()
		c.inChan <- &WsMessage{
			MessageType: websocket.TextMessage,
			Data:        types.XtermMessage{Output: string(msg)},
		}
		c.wsClose()
		return
	}

}

// ResponseJSON response to client
func ResponseJSON(w http.ResponseWriter, status int, v interface{}) error {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(status)
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	_, err = w.Write(data)
	return err
}

//Next executor回调获取web是否resize
func (handler *streamHandler) Next() (size *remotecommand.TerminalSize) {
	ret := <-handler.resizeEvent
	size = &ret
	return
}

// executor回调读取web端的输入
func (handler *streamHandler) Read(p []byte) (size int, err error) {

	// 读web发来的输入
	msg, err := handler.wsConn.WsRead()
	if err != nil {
		handler.wsConn.wsClose()
		return
	}
	if msg.Data.MsgType == "reset" {
		// 放到channel里，等remotecommand executor调用Next取走
		handler.resizeEvent <- remotecommand.TerminalSize{Width: msg.Data.Cols, Height: msg.Data.Rows}
	}

	size = len(msg.Data.Input)
	copy(p, msg.Data.Input)
	return
}

// executor回调向web端输出
func (handler *streamHandler) Write(p []byte) (size int, err error) {
	// 产生副本
	copyData := make([]byte, len(p))
	copy(copyData, p)
	size = len(p)
	err = handler.wsConn.wsWrite(websocket.TextMessage, copyData)
	return
}

// 周期上报操作记录
func (c *wsConn) periodicRecord() *types.AuditRecord {

	inputRecord := c.flushInputRecord()
	outputRecord := c.flushOutputRecord()

	// 如果输入输出为空则取消此次上报
	if len(inputRecord) == 0 && len(outputRecord) == 0 {
		return nil
	}

	data := types.AuditRecord{
		InputRecord:  inputRecord,
		OutputRecord: outputRecord,
		SessionID:    c.SessionID,
		Context:      nil,
		ProjectID:    c.Project,
		ClusterID:    c.Cluster,
		UserPodName:  c.PodName,
		Username:     c.Username,
	}

	return &data

}

// 获取输入记录
func (c *wsConn) flushInputRecord() string {

	if c.InputRecord == "" {
		return ""
	}

	lineMsg := strings.Split(c.InputRecord, InputLineBreaker)
	var record, cmd string
	for _, s := range lineMsg {
		cmd = cleanBashEscape(s)
		if cmd == "" {
			continue
		}
		record += "\r\n" + fmt.Sprintf("%s: %s",
			time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05.06"), cmd)
		blog.Debug(record)
	}

	c.InputRecord = ""

	return record
}

// 去除bash转义字符
func cleanBashEscape(text string) string {
	// 删除转移字符
	re, err := regexp.Compile(AnsiEscape)
	if err != nil {
		return ""
	}
	text = re.ReplaceAllString(text, "")

	return text
}

// 获取输出记录
func (c *wsConn) flushOutputRecord() string {
	if c.OutputRecord == "" {
		return ""
	}

	lineMsg := strings.Split(c.OutputRecord, OutputLineBreaker)
	var record, cmd string
	for _, s := range lineMsg {
		cmd = cleanBashEscape(s)
		if cmd == "" {
			continue
		}
		record += "\r\n" + fmt.Sprintf("%s: %s",
			time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05.06"), cmd)
		blog.Debug(record)
	}
	c.OutputRecord = ""

	return record
}
