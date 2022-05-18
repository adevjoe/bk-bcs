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

package runtimex

import (
	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-helm-manager/internal/common/headerkey"
)

// CustomHeaderMatcher for http header
func CustomHeaderMatcher(key string) (string, bool) {
	switch key {
	case headerkey.RequestIDKey:
		return headerkey.RequestIDKey, true
	case headerkey.UsernameKey:
		return headerkey.UsernameKey, true
	default:
		return runtime.DefaultHeaderMatcher(key)
	}
}
