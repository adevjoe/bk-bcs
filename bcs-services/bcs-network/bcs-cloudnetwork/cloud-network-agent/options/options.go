/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.,
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under,
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package options

import (
	"bk-bcs/bcs-common/common/blog"
	"bk-bcs/bcs-common/common/conf"
)

const (
	// ServiceRegistryKubernetes service discovery for k8s
	ServiceRegistryKubernetes = "kubernetes"
	// ServiceRegistryMesos service discovery for mesos
	ServiceRegistryMesos = "mesos"
)

// NetworkOption the option of bcs elastic network interface controller
type NetworkOption struct {
	conf.ServiceConfig
	conf.ServerOnlyCertConfig
	conf.FileConfig
	conf.MetricConfig
	conf.LogConfig
	conf.ProcessConfig

	Cluster              string `json:"cluster" value:"" usage:"cluster for bcs"`
	ServiceRegistry      string `json:"serviceRegistry" value:"kubernetes" usage:"registry for service discovery; [kubernetes, mesos]"`
	Kubeconfig           string `json:"kubeconfig" value:"" usage:"kubeconfig for kube-apiserver, Only required if out-of-cluster."`
	KubeResyncPeriod     int    `json:"kubeResyncPeried" value:"300" usage:"resync interval for informer factory in seconds; (default 300)"`
	KubeCacheSyncTimeout int    `json:"kubeCacheSyncTimeout" value:"10" usage:"wait for kube cache sync timeout in seconds; (default 10)"`
	CheckInterval        int    `json:"checkInterval" value:"300" usage:"interval for checking ip rules and route tables"`

	SecurityGroups string `json:"securityGroups" value:"" usage:"vpc security groups used by elastic network interface, e.g. \"vpc-001:sg-111;vpc-002:sg-222\""`
	Subnets        string `json:"subnets" value:"" usage:"the ips of elastic network interface come from these subnets, if no subnet specified means using the subnet which the node belongs to, e.g. \"vpc-001:subnet-111,subnet-112;vpc-002:subnet-222\""`
	EniNum         int    `json:"eniNum" value:"0" usage:"the number of elastic network interface for each node; default is 0, means apply for as many eni as possible"`
	IPNumPerEni    int    `json:"ipNumPerEni" value:"0" usage:"the number of ip for each eni; default is 0, means apply for as many ip as possible"`
	Ifaces         string `json:"ifaces" value:"eth1" usage:"use ip of these network interfaces as node identity, split with comma or semicolon"`

	Debug bool `json:"debug" value:"false" usage:"open pprof"`
}

// New new option
func New() *NetworkOption {
	return &NetworkOption{}
}

// Parse parse options
func Parse(opt *NetworkOption) {
	conf.Parse(opt)

	// validate config
	if opt.ServiceRegistry != ServiceRegistryKubernetes && opt.ServiceRegistry != ServiceRegistryMesos {
		blog.Fatal("registry for service discovery, available values [kubernetes, mesos]")
	}
	if len(opt.Kubeconfig) == 0 {
		blog.Fatal("kubeconfig cannot be empty")
	}
}
