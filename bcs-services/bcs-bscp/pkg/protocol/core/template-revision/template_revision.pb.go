// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: template_revision.proto

package pbtr

import (
	base "github.com/TencentBlueKing/bk-bcs/bcs-services/bcs-bscp/pkg/protocol/core/base"
	config_item "github.com/TencentBlueKing/bk-bcs/bcs-services/bcs-bscp/pkg/protocol/core/config-item"
	content "github.com/TencentBlueKing/bk-bcs/bcs-services/bcs-bscp/pkg/protocol/core/content"
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

// Template source resource reference: pkg/dal/table/template_revision.go
type TemplateRevision struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint32                      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Spec       *TemplateRevisionSpec       `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	Attachment *TemplateRevisionAttachment `protobuf:"bytes,3,opt,name=attachment,proto3" json:"attachment,omitempty"`
	Revision   *base.CreatedRevision       `protobuf:"bytes,4,opt,name=revision,proto3" json:"revision,omitempty"`
}

func (x *TemplateRevision) Reset() {
	*x = TemplateRevision{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_revision_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateRevision) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateRevision) ProtoMessage() {}

func (x *TemplateRevision) ProtoReflect() protoreflect.Message {
	mi := &file_template_revision_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateRevision.ProtoReflect.Descriptor instead.
func (*TemplateRevision) Descriptor() ([]byte, []int) {
	return file_template_revision_proto_rawDescGZIP(), []int{0}
}

func (x *TemplateRevision) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TemplateRevision) GetSpec() *TemplateRevisionSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *TemplateRevision) GetAttachment() *TemplateRevisionAttachment {
	if x != nil {
		return x.Attachment
	}
	return nil
}

func (x *TemplateRevision) GetRevision() *base.CreatedRevision {
	if x != nil {
		return x.Revision
	}
	return nil
}

// TemplateRevisionSpec source resource reference: pkg/dal/table/template_revision.go
type TemplateRevisionSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RevisionName string                      `protobuf:"bytes,1,opt,name=revision_name,json=revisionName,proto3" json:"revision_name,omitempty"`
	RevisionMemo string                      `protobuf:"bytes,2,opt,name=revision_memo,json=revisionMemo,proto3" json:"revision_memo,omitempty"`
	Name         string                      `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Path         string                      `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	FileType     string                      `protobuf:"bytes,5,opt,name=file_type,json=fileType,proto3" json:"file_type,omitempty"` // file_type is enum type
	FileMode     string                      `protobuf:"bytes,6,opt,name=file_mode,json=fileMode,proto3" json:"file_mode,omitempty"` // file_mode is enum type
	Permission   *config_item.FilePermission `protobuf:"bytes,7,opt,name=permission,proto3" json:"permission,omitempty"`
	ContentSpec  *content.ContentSpec        `protobuf:"bytes,8,opt,name=content_spec,json=contentSpec,proto3" json:"content_spec,omitempty"`
}

func (x *TemplateRevisionSpec) Reset() {
	*x = TemplateRevisionSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_revision_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateRevisionSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateRevisionSpec) ProtoMessage() {}

func (x *TemplateRevisionSpec) ProtoReflect() protoreflect.Message {
	mi := &file_template_revision_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateRevisionSpec.ProtoReflect.Descriptor instead.
func (*TemplateRevisionSpec) Descriptor() ([]byte, []int) {
	return file_template_revision_proto_rawDescGZIP(), []int{1}
}

func (x *TemplateRevisionSpec) GetRevisionName() string {
	if x != nil {
		return x.RevisionName
	}
	return ""
}

func (x *TemplateRevisionSpec) GetRevisionMemo() string {
	if x != nil {
		return x.RevisionMemo
	}
	return ""
}

func (x *TemplateRevisionSpec) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TemplateRevisionSpec) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *TemplateRevisionSpec) GetFileType() string {
	if x != nil {
		return x.FileType
	}
	return ""
}

func (x *TemplateRevisionSpec) GetFileMode() string {
	if x != nil {
		return x.FileMode
	}
	return ""
}

func (x *TemplateRevisionSpec) GetPermission() *config_item.FilePermission {
	if x != nil {
		return x.Permission
	}
	return nil
}

func (x *TemplateRevisionSpec) GetContentSpec() *content.ContentSpec {
	if x != nil {
		return x.ContentSpec
	}
	return nil
}

// TemplateRevisionAttachment source resource reference: pkg/dal/table/template_revision.go
type TemplateRevisionAttachment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BizId           uint32 `protobuf:"varint,1,opt,name=biz_id,json=bizId,proto3" json:"biz_id,omitempty"`
	TemplateSpaceId uint32 `protobuf:"varint,2,opt,name=template_space_id,json=templateSpaceId,proto3" json:"template_space_id,omitempty"`
	TemplateId      uint32 `protobuf:"varint,3,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
}

func (x *TemplateRevisionAttachment) Reset() {
	*x = TemplateRevisionAttachment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_revision_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateRevisionAttachment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateRevisionAttachment) ProtoMessage() {}

func (x *TemplateRevisionAttachment) ProtoReflect() protoreflect.Message {
	mi := &file_template_revision_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateRevisionAttachment.ProtoReflect.Descriptor instead.
func (*TemplateRevisionAttachment) Descriptor() ([]byte, []int) {
	return file_template_revision_proto_rawDescGZIP(), []int{2}
}

func (x *TemplateRevisionAttachment) GetBizId() uint32 {
	if x != nil {
		return x.BizId
	}
	return 0
}

func (x *TemplateRevisionAttachment) GetTemplateSpaceId() uint32 {
	if x != nil {
		return x.TemplateSpaceId
	}
	return 0
}

func (x *TemplateRevisionAttachment) GetTemplateId() uint32 {
	if x != nil {
		return x.TemplateId
	}
	return 0
}

// TemplateRevisionNamesDetail is template revision names Detail.
type TemplateRevisionNamesDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateId               uint32                                      `protobuf:"varint,1,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	TemplateName             string                                      `protobuf:"bytes,2,opt,name=template_name,json=templateName,proto3" json:"template_name,omitempty"`
	LatestTemplateRevisionId uint32                                      `protobuf:"varint,3,opt,name=latest_template_revision_id,json=latestTemplateRevisionId,proto3" json:"latest_template_revision_id,omitempty"`
	TemplateRevisions        []*TemplateRevisionNamesDetailRevisionNames `protobuf:"bytes,4,rep,name=template_revisions,json=templateRevisions,proto3" json:"template_revisions,omitempty"`
}

func (x *TemplateRevisionNamesDetail) Reset() {
	*x = TemplateRevisionNamesDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_revision_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateRevisionNamesDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateRevisionNamesDetail) ProtoMessage() {}

func (x *TemplateRevisionNamesDetail) ProtoReflect() protoreflect.Message {
	mi := &file_template_revision_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateRevisionNamesDetail.ProtoReflect.Descriptor instead.
func (*TemplateRevisionNamesDetail) Descriptor() ([]byte, []int) {
	return file_template_revision_proto_rawDescGZIP(), []int{3}
}

func (x *TemplateRevisionNamesDetail) GetTemplateId() uint32 {
	if x != nil {
		return x.TemplateId
	}
	return 0
}

func (x *TemplateRevisionNamesDetail) GetTemplateName() string {
	if x != nil {
		return x.TemplateName
	}
	return ""
}

func (x *TemplateRevisionNamesDetail) GetLatestTemplateRevisionId() uint32 {
	if x != nil {
		return x.LatestTemplateRevisionId
	}
	return 0
}

func (x *TemplateRevisionNamesDetail) GetTemplateRevisions() []*TemplateRevisionNamesDetailRevisionNames {
	if x != nil {
		return x.TemplateRevisions
	}
	return nil
}

type TemplateRevisionNamesDetailRevisionNames struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateRevisionId   uint32 `protobuf:"varint,1,opt,name=template_revision_id,json=templateRevisionId,proto3" json:"template_revision_id,omitempty"`
	TemplateRevisionName string `protobuf:"bytes,2,opt,name=template_revision_name,json=templateRevisionName,proto3" json:"template_revision_name,omitempty"`
	TemplateRevisionMemo string `protobuf:"bytes,3,opt,name=template_revision_memo,json=templateRevisionMemo,proto3" json:"template_revision_memo,omitempty"`
}

func (x *TemplateRevisionNamesDetailRevisionNames) Reset() {
	*x = TemplateRevisionNamesDetailRevisionNames{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_revision_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateRevisionNamesDetailRevisionNames) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateRevisionNamesDetailRevisionNames) ProtoMessage() {}

func (x *TemplateRevisionNamesDetailRevisionNames) ProtoReflect() protoreflect.Message {
	mi := &file_template_revision_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateRevisionNamesDetailRevisionNames.ProtoReflect.Descriptor instead.
func (*TemplateRevisionNamesDetailRevisionNames) Descriptor() ([]byte, []int) {
	return file_template_revision_proto_rawDescGZIP(), []int{3, 0}
}

func (x *TemplateRevisionNamesDetailRevisionNames) GetTemplateRevisionId() uint32 {
	if x != nil {
		return x.TemplateRevisionId
	}
	return 0
}

func (x *TemplateRevisionNamesDetailRevisionNames) GetTemplateRevisionName() string {
	if x != nil {
		return x.TemplateRevisionName
	}
	return ""
}

func (x *TemplateRevisionNamesDetailRevisionNames) GetTemplateRevisionMemo() string {
	if x != nil {
		return x.TemplateRevisionMemo
	}
	return ""
}

var File_template_revision_proto protoreflect.FileDescriptor

var file_template_revision_proto_rawDesc = []byte{
	0x0a, 0x17, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x70, 0x62, 0x74, 0x72, 0x1a,
	0x21, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2d, 0x69, 0x74, 0x65,
	0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc9, 0x01, 0x0a,
	0x10, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x2e, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x70, 0x62, 0x74, 0x72, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65,
	0x63, 0x12, 0x40, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x62, 0x74, 0x72, 0x2e, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x33, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x62, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x08,
	0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xb3, 0x02, 0x0a, 0x14, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65,
	0x63, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72,
	0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x6d, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x34, 0x0a,
	0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x70, 0x62, 0x63, 0x69, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x73,
	0x70, 0x65, 0x63, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x62, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x53, 0x70, 0x65,
	0x63, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x53, 0x70, 0x65, 0x63, 0x22, 0x80,
	0x01, 0x0a, 0x1a, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x15, 0x0a,
	0x06, 0x62, 0x69, 0x7a, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x62,
	0x69, 0x7a, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x5f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49,
	0x64, 0x22, 0xb4, 0x03, 0x0a, 0x1b, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x1b, 0x6c, 0x61, 0x74, 0x65, 0x73,
	0x74, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x18, 0x6c, 0x61,
	0x74, 0x65, 0x73, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x5f, 0x0a, 0x12, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x30, 0x2e, 0x70, 0x62, 0x74, 0x72, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x52, 0x11, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0xae, 0x01, 0x0a, 0x0e, 0x72, 0x65, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x14, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x16,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x34, 0x0a, 0x16, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x72,
	0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x14, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x6d, 0x6f, 0x42, 0x62, 0x5a, 0x60, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x42, 0x6c,
	0x75, 0x65, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x62, 0x6b, 0x2d, 0x62, 0x63, 0x73, 0x2f, 0x62, 0x63,
	0x73, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x63, 0x73, 0x2d, 0x62,
	0x73, 0x63, 0x70, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2d, 0x72,
	0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x3b, 0x70, 0x62, 0x74, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_template_revision_proto_rawDescOnce sync.Once
	file_template_revision_proto_rawDescData = file_template_revision_proto_rawDesc
)

func file_template_revision_proto_rawDescGZIP() []byte {
	file_template_revision_proto_rawDescOnce.Do(func() {
		file_template_revision_proto_rawDescData = protoimpl.X.CompressGZIP(file_template_revision_proto_rawDescData)
	})
	return file_template_revision_proto_rawDescData
}

var file_template_revision_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_template_revision_proto_goTypes = []interface{}{
	(*TemplateRevision)(nil),                         // 0: pbtr.TemplateRevision
	(*TemplateRevisionSpec)(nil),                     // 1: pbtr.TemplateRevisionSpec
	(*TemplateRevisionAttachment)(nil),               // 2: pbtr.TemplateRevisionAttachment
	(*TemplateRevisionNamesDetail)(nil),              // 3: pbtr.TemplateRevisionNamesDetail
	(*TemplateRevisionNamesDetailRevisionNames)(nil), // 4: pbtr.TemplateRevisionNamesDetail.revision_names
	(*base.CreatedRevision)(nil),                     // 5: pbbase.CreatedRevision
	(*config_item.FilePermission)(nil),               // 6: pbci.FilePermission
	(*content.ContentSpec)(nil),                      // 7: pbcontent.ContentSpec
}
var file_template_revision_proto_depIdxs = []int32{
	1, // 0: pbtr.TemplateRevision.spec:type_name -> pbtr.TemplateRevisionSpec
	2, // 1: pbtr.TemplateRevision.attachment:type_name -> pbtr.TemplateRevisionAttachment
	5, // 2: pbtr.TemplateRevision.revision:type_name -> pbbase.CreatedRevision
	6, // 3: pbtr.TemplateRevisionSpec.permission:type_name -> pbci.FilePermission
	7, // 4: pbtr.TemplateRevisionSpec.content_spec:type_name -> pbcontent.ContentSpec
	4, // 5: pbtr.TemplateRevisionNamesDetail.template_revisions:type_name -> pbtr.TemplateRevisionNamesDetail.revision_names
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_template_revision_proto_init() }
func file_template_revision_proto_init() {
	if File_template_revision_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_template_revision_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateRevision); i {
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
		file_template_revision_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateRevisionSpec); i {
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
		file_template_revision_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateRevisionAttachment); i {
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
		file_template_revision_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateRevisionNamesDetail); i {
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
		file_template_revision_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateRevisionNamesDetailRevisionNames); i {
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
			RawDescriptor: file_template_revision_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_template_revision_proto_goTypes,
		DependencyIndexes: file_template_revision_proto_depIdxs,
		MessageInfos:      file_template_revision_proto_msgTypes,
	}.Build()
	File_template_revision_proto = out.File
	file_template_revision_proto_rawDesc = nil
	file_template_revision_proto_goTypes = nil
	file_template_revision_proto_depIdxs = nil
}
