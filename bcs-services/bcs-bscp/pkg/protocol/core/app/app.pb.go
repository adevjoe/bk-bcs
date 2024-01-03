// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: app.proto

package pbapp

import (
	base "github.com/TencentBlueKing/bk-bcs/bcs-services/bcs-bscp/pkg/protocol/core/base"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// App source resource reference: pkg/dal/table/app.go
type App struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            uint32         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	BizId         uint32         `protobuf:"varint,2,opt,name=biz_id,json=bizId,proto3" json:"biz_id,omitempty"` // Deprecated: use space instead
	SpaceId       string         `protobuf:"bytes,3,opt,name=space_id,json=spaceId,proto3" json:"space_id,omitempty"`
	SpaceTypeId   string         `protobuf:"bytes,4,opt,name=space_type_id,json=spaceTypeId,proto3" json:"space_type_id,omitempty"`
	SpaceName     string         `protobuf:"bytes,5,opt,name=space_name,json=spaceName,proto3" json:"space_name,omitempty"`
	SpaceTypeName string         `protobuf:"bytes,6,opt,name=space_type_name,json=spaceTypeName,proto3" json:"space_type_name,omitempty"`
	Spec          *AppSpec       `protobuf:"bytes,7,opt,name=spec,proto3" json:"spec,omitempty"`
	Revision      *base.Revision `protobuf:"bytes,8,opt,name=revision,proto3" json:"revision,omitempty"`
}

func (x *App) Reset() {
	*x = App{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *App) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*App) ProtoMessage() {}

func (x *App) ProtoReflect() protoreflect.Message {
	mi := &file_app_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use App.ProtoReflect.Descriptor instead.
func (*App) Descriptor() ([]byte, []int) {
	return file_app_proto_rawDescGZIP(), []int{0}
}

func (x *App) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *App) GetBizId() uint32 {
	if x != nil {
		return x.BizId
	}
	return 0
}

func (x *App) GetSpaceId() string {
	if x != nil {
		return x.SpaceId
	}
	return ""
}

func (x *App) GetSpaceTypeId() string {
	if x != nil {
		return x.SpaceTypeId
	}
	return ""
}

func (x *App) GetSpaceName() string {
	if x != nil {
		return x.SpaceName
	}
	return ""
}

func (x *App) GetSpaceTypeName() string {
	if x != nil {
		return x.SpaceTypeName
	}
	return ""
}

func (x *App) GetSpec() *AppSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *App) GetRevision() *base.Revision {
	if x != nil {
		return x.Revision
	}
	return nil
}

// AppSpec source resource reference: pkg/dal/table/app.go
type AppSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ConfigType string  `protobuf:"bytes,2,opt,name=config_type,json=configType,proto3" json:"config_type,omitempty"` // config_type is enum type
	Mode       string  `protobuf:"bytes,4,opt,name=mode,proto3" json:"mode,omitempty"`                               // mode is enum type.
	Memo       string  `protobuf:"bytes,5,opt,name=memo,proto3" json:"memo,omitempty"`
	Reload     *Reload `protobuf:"bytes,6,opt,name=reload,proto3" json:"reload,omitempty"`
	Alias      string  `protobuf:"bytes,7,opt,name=alias,proto3" json:"alias,omitempty"`
	DataType   string  `protobuf:"bytes,8,opt,name=data_type,json=dataType,proto3" json:"data_type,omitempty"` // data_type is enum type.
}

func (x *AppSpec) Reset() {
	*x = AppSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppSpec) ProtoMessage() {}

func (x *AppSpec) ProtoReflect() protoreflect.Message {
	mi := &file_app_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppSpec.ProtoReflect.Descriptor instead.
func (*AppSpec) Descriptor() ([]byte, []int) {
	return file_app_proto_rawDescGZIP(), []int{1}
}

func (x *AppSpec) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AppSpec) GetConfigType() string {
	if x != nil {
		return x.ConfigType
	}
	return ""
}

func (x *AppSpec) GetMode() string {
	if x != nil {
		return x.Mode
	}
	return ""
}

func (x *AppSpec) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *AppSpec) GetReload() *Reload {
	if x != nil {
		return x.Reload
	}
	return nil
}

func (x *AppSpec) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

func (x *AppSpec) GetDataType() string {
	if x != nil {
		return x.DataType
	}
	return ""
}

// ReloadSpec source resource reference: pkg/dal/table/app.go
type Reload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReloadType     string          `protobuf:"bytes,1,opt,name=reload_type,json=reloadType,proto3" json:"reload_type,omitempty"` // reload_type is enum type
	FileReloadSpec *FileReloadSpec `protobuf:"bytes,2,opt,name=file_reload_spec,json=fileReloadSpec,proto3" json:"file_reload_spec,omitempty"`
}

func (x *Reload) Reset() {
	*x = Reload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reload) ProtoMessage() {}

func (x *Reload) ProtoReflect() protoreflect.Message {
	mi := &file_app_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reload.ProtoReflect.Descriptor instead.
func (*Reload) Descriptor() ([]byte, []int) {
	return file_app_proto_rawDescGZIP(), []int{2}
}

func (x *Reload) GetReloadType() string {
	if x != nil {
		return x.ReloadType
	}
	return ""
}

func (x *Reload) GetFileReloadSpec() *FileReloadSpec {
	if x != nil {
		return x.FileReloadSpec
	}
	return nil
}

// FileReloadSpec source resource reference: pkg/dal/table/app.go
type FileReloadSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReloadFilePath string `protobuf:"bytes,1,opt,name=reload_file_path,json=reloadFilePath,proto3" json:"reload_file_path,omitempty"`
}

func (x *FileReloadSpec) Reset() {
	*x = FileReloadSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileReloadSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileReloadSpec) ProtoMessage() {}

func (x *FileReloadSpec) ProtoReflect() protoreflect.Message {
	mi := &file_app_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileReloadSpec.ProtoReflect.Descriptor instead.
func (*FileReloadSpec) Descriptor() ([]byte, []int) {
	return file_app_proto_rawDescGZIP(), []int{3}
}

func (x *FileReloadSpec) GetReloadFilePath() string {
	if x != nil {
		return x.ReloadFilePath
	}
	return ""
}

var File_app_proto protoreflect.FileDescriptor

var file_app_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x62, 0x61,
	0x70, 0x70, 0x1a, 0x21, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x02, 0x0a, 0x03, 0x41, 0x70, 0x70, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x15, 0x0a,
	0x06, 0x62, 0x69, 0x7a, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x62,
	0x69, 0x7a, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12,
	0x22, 0x0a, 0x0d, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x70, 0x61, 0x63, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x73, 0x70,
	0x65, 0x63, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x61, 0x70, 0x70,
	0x2e, 0x41, 0x70, 0x70, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x2c,
	0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x70, 0x62, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xc0, 0x01, 0x0a,
	0x07, 0x41, 0x70, 0x70, 0x53, 0x70, 0x65, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x6f, 0x64,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x12, 0x25, 0x0a, 0x06, 0x72, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x61, 0x70, 0x70, 0x2e, 0x52, 0x65,
	0x6c, 0x6f, 0x61, 0x64, 0x52, 0x06, 0x72, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6c, 0x69,
	0x61, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x22,
	0x6a, 0x0a, 0x06, 0x52, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x6c,
	0x6f, 0x61, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x72, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3f, 0x0a, 0x10, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x72, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x62, 0x61, 0x70, 0x70, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x70, 0x65, 0x63, 0x52, 0x0e, 0x66, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x70, 0x65, 0x63, 0x22, 0x3a, 0x0a, 0x0e, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x70, 0x65, 0x63, 0x12, 0x28, 0x0a,
	0x10, 0x72, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x46,
	0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x42, 0x55, 0x5a, 0x53, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x42, 0x6c, 0x75,
	0x65, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x62, 0x6b, 0x2d, 0x62, 0x63, 0x73, 0x2f, 0x62, 0x63, 0x73,
	0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x63, 0x73, 0x2d, 0x62, 0x73,
	0x63, 0x70, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x3b, 0x70, 0x62, 0x61, 0x70, 0x70, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_proto_rawDescOnce sync.Once
	file_app_proto_rawDescData = file_app_proto_rawDesc
)

func file_app_proto_rawDescGZIP() []byte {
	file_app_proto_rawDescOnce.Do(func() {
		file_app_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_proto_rawDescData)
	})
	return file_app_proto_rawDescData
}

var file_app_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_app_proto_goTypes = []interface{}{
	(*App)(nil),            // 0: pbapp.App
	(*AppSpec)(nil),        // 1: pbapp.AppSpec
	(*Reload)(nil),         // 2: pbapp.Reload
	(*FileReloadSpec)(nil), // 3: pbapp.FileReloadSpec
	(*base.Revision)(nil),  // 4: pbbase.Revision
}
var file_app_proto_depIdxs = []int32{
	1, // 0: pbapp.App.spec:type_name -> pbapp.AppSpec
	4, // 1: pbapp.App.revision:type_name -> pbbase.Revision
	2, // 2: pbapp.AppSpec.reload:type_name -> pbapp.Reload
	3, // 3: pbapp.Reload.file_reload_spec:type_name -> pbapp.FileReloadSpec
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_app_proto_init() }
func file_app_proto_init() {
	if File_app_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*App); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_app_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_app_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reload); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_app_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileReloadSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_app_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_app_proto_goTypes,
		DependencyIndexes: file_app_proto_depIdxs,
		MessageInfos:      file_app_proto_msgTypes,
	}.Build()
	File_app_proto = out.File
	file_app_proto_rawDesc = nil
	file_app_proto_goTypes = nil
	file_app_proto_depIdxs = nil
}
