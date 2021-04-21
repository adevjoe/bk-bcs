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

package configuration

import "fmt"

type AggregationBcsStorageInfo struct {
	bcsStoragePodUrlBase string
	bcsStorageToken      string
}

func (asi *AggregationBcsStorageInfo) SetBcsStorageInfo(acm *AggregationConfigMapInfo) {
	asi.bcsStoragePodUrlBase = fmt.Sprintf("%s/%s", acm.GetBcsStorageAddress(), acm.GetBcsStoragePodUri())
	asi.bcsStorageToken = acm.bcsStorageToken
}

func (asi *AggregationBcsStorageInfo) GetBcsStorageToken() string {
	return asi.bcsStorageToken
}

func (asi *AggregationBcsStorageInfo) GetBcsStoragePodUrlBase() string {
	return asi.bcsStoragePodUrlBase
}
