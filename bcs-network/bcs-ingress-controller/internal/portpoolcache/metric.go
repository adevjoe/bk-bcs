/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under,
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package portpoolcache

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

const metricCollectInterval = 30 * time.Second

var (
	portPoolCapacityMetric = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "bkbcs_ingressctrl",
		Subsystem: "cache",
		Name:      "poolitem_portitem_total",
		Help:      "port pool capacty",
	}, []string{"poolkey", "itemname", "protocol"})
	portPoolAllocatedMetric = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "bkbcs_ingressctrl",
		Subsystem: "cache",
		Name:      "poolitem_portitem_used",
		Help:      "port pool capacty",
	}, []string{"poolkey", "itemname", "protocol"})
)

func init() {
	metrics.Registry.MustRegister(portPoolCapacityMetric)
	metrics.Registry.MustRegister(portPoolAllocatedMetric)
}
