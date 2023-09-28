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
 */

package utils

import (
	"encoding/json"
	"fmt"
	"reflect"

	structpb "github.com/golang/protobuf/ptypes/struct"
	"google.golang.org/protobuf/encoding/protojson"
	spb "google.golang.org/protobuf/types/known/structpb"
)

// MapToProtobufStruct converts a map to a protobuf struct.
func MapToProtobufStruct(m map[string]interface{}) (*structpb.Struct, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	s := &structpb.Struct{}
	err = protojson.Unmarshal(b, s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// ProtobufStructToMap converts a protobuf struct to a map.
func ProtobufStructToMap(s *structpb.Struct) (map[string]interface{}, error) {
	b, err := protojson.Marshal(s)
	if err != nil {
		return nil, err
	}
	m := make(map[string]interface{})
	err = json.Unmarshal(b, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// StructToProtobufStruct converts a struct to a protobuf struct.
func StructToProtobufStruct(s interface{}) (*structpb.Struct, error) {
	return MapToProtobufStruct(s.(map[string]interface{}))
}

// ProtobufStructToStruct converts a protobuf struct to a struct.
func ProtobufStructToStruct(s *structpb.Struct) (interface{}, error) {
	return ProtobufStructToMap(s)
}

// MarshalInterfaceToListValue trans interface to ListValue
func MarshalInterfaceToListValue(data interface{}) (*structpb.ListValue, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	l := &structpb.ListValue{}
	err = l.UnmarshalJSON(b)
	if err != nil {
		return nil, err
	}
	return l, nil
}

// MarshalInterfaceToValue trans interface to Struct
func MarshalInterfaceToValue(data interface{}) (*structpb.Struct, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	l := &structpb.Struct{}
	err = l.UnmarshalJSON(b)
	if err != nil {
		return nil, err
	}
	return l, nil
}

// MarshalInterfacesToListValue trans interface slice to ListValue
func MarshalInterfacesToListValue(objs interface{}) (*structpb.ListValue, error) {
	b, err := json.Marshal(objs)
	if err != nil {
		return nil, err
	}
	l := &structpb.ListValue{}
	err = l.UnmarshalJSON(b)
	if err != nil {
		return nil, err
	}
	return l, nil
}

// Map2pbStruct convert map to pbstruct
// ref: https://devnote.pro/posts/10000050901242
func Map2pbStruct(m map[string]interface{}) *spb.Struct {
	size := len(m)
	if size == 0 {
		return nil
	}
	fields := make(map[string]*spb.Value, size)
	for k, v := range m {
		fields[k] = InterfaceToValue(v)
	}
	return &spb.Struct{
		Fields: fields,
	}
}

// InterfaceToValue convert interface to value
func InterfaceToValue(v interface{}) *spb.Value {
	switch v := v.(type) {
	case nil:
		return nil
	case bool:
		return &spb.Value{
			Kind: &spb.Value_BoolValue{
				BoolValue: v,
			},
		}
	case int:
		return &spb.Value{
			Kind: &spb.Value_NumberValue{
				NumberValue: float64(v),
			},
		}
	case int8:
		return &spb.Value{
			Kind: &spb.Value_NumberValue{
				NumberValue: float64(v),
			},
		}
	case int32:
		return &spb.Value{
			Kind: &spb.Value_NumberValue{
				NumberValue: float64(v),
			},
		}
	case int64:
		return &spb.Value{
			Kind: &spb.Value_NumberValue{
				NumberValue: float64(v),
			},
		}
	case uint:
		return &spb.Value{
			Kind: &spb.Value_NumberValue{
				NumberValue: float64(v),
			},
		}
	case uint8:
		return &spb.Value{
			Kind: &spb.Value_NumberValue{
				NumberValue: float64(v),
			},
		}
	case uint32:
		return &spb.Value{
			Kind: &spb.Value_NumberValue{
				NumberValue: float64(v),
			},
		}
	case uint64:
		return &spb.Value{
			Kind: &spb.Value_NumberValue{
				NumberValue: float64(v),
			},
		}
	case float32:
		return &spb.Value{
			Kind: &spb.Value_NumberValue{
				NumberValue: float64(v),
			},
		}
	case float64:
		return &spb.Value{
			Kind: &spb.Value_NumberValue{
				NumberValue: v,
			},
		}
	case string:
		return &spb.Value{
			Kind: &spb.Value_StringValue{
				StringValue: v,
			},
		}
	case error:
		return &spb.Value{
			Kind: &spb.Value_StringValue{
				StringValue: v.Error(),
			},
		}
	default:
		// 回退为其他类型
		return toValue(reflect.ValueOf(v))
	}
}

func toValue(v reflect.Value) *spb.Value { // nolint
	switch v.Kind() {
	case reflect.Bool:
		return &spb.Value{
			Kind: &spb.Value_BoolValue{
				BoolValue: v.Bool(),
			},
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return &spb.Value{
			Kind: &spb.Value_NumberValue{
				NumberValue: float64(v.Int()),
			},
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return &spb.Value{
			Kind: &spb.Value_NumberValue{
				NumberValue: float64(v.Uint()),
			},
		}
	case reflect.Float32, reflect.Float64:
		return &spb.Value{
			Kind: &spb.Value_NumberValue{
				NumberValue: v.Float(),
			},
		}
	case reflect.Ptr:
		if v.IsNil() {
			return nil
		}
		return toValue(reflect.Indirect(v))
	case reflect.Array, reflect.Slice:
		size := v.Len()
		if size == 0 {
			return nil
		}
		values := make([]*spb.Value, size)
		for i := 0; i < size; i++ {
			values[i] = toValue(v.Index(i))
		}
		return &spb.Value{
			Kind: &spb.Value_ListValue{
				ListValue: &spb.ListValue{
					Values: values,
				},
			},
		}
	case reflect.Struct:
		t := v.Type()
		size := v.NumField()
		if size == 0 {
			return nil
		}
		fields := make(map[string]*spb.Value, size)
		for i := 0; i < size; i++ {
			name := t.Field(i).Name
			if len(name) > 0 && 'A' <= name[0] && name[0] <= 'Z' {
				fields[name] = toValue(v.Field(i))
			}
		}
		if len(fields) == 0 {
			return nil
		}
		return &spb.Value{
			Kind: &spb.Value_StructValue{
				StructValue: &spb.Struct{
					Fields: fields,
				},
			},
		}
	case reflect.Map:
		keys := v.MapKeys()
		if len(keys) == 0 {
			return nil
		}
		fields := make(map[string]*spb.Value, len(keys))
		for _, k := range keys {
			if k.Kind() == reflect.String {
				fields[k.String()] = toValue(v.MapIndex(k))
			}
		}
		if len(fields) == 0 {
			return nil
		}
		return &spb.Value{
			Kind: &spb.Value_StructValue{
				StructValue: &spb.Struct{
					Fields: fields,
				},
			},
		}
	default:
		// 最后排序
		return &spb.Value{
			Kind: &spb.Value_StringValue{
				StringValue: fmt.Sprint(v),
			},
		}
	}
}
