// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: usertag.proto

package proto

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type UserTag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId    int64                `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name      string               `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamp.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *UserTag) Reset() {
	*x = UserTag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usertag_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserTag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserTag) ProtoMessage() {}

func (x *UserTag) ProtoReflect() protoreflect.Message {
	mi := &file_usertag_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserTag.ProtoReflect.Descriptor instead.
func (*UserTag) Descriptor() ([]byte, []int) {
	return file_usertag_proto_rawDescGZIP(), []int{0}
}

func (x *UserTag) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserTag) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserTag) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserTag) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *UserTag) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type UserTagId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserTagId) Reset() {
	*x = UserTagId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usertag_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserTagId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserTagId) ProtoMessage() {}

func (x *UserTagId) ProtoReflect() protoreflect.Message {
	mi := &file_usertag_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserTagId.ProtoReflect.Descriptor instead.
func (*UserTagId) Descriptor() ([]byte, []int) {
	return file_usertag_proto_rawDescGZIP(), []int{1}
}

func (x *UserTagId) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UserIdWithUserTagId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Id     int64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserIdWithUserTagId) Reset() {
	*x = UserIdWithUserTagId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usertag_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserIdWithUserTagId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserIdWithUserTagId) ProtoMessage() {}

func (x *UserIdWithUserTagId) ProtoReflect() protoreflect.Message {
	mi := &file_usertag_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserIdWithUserTagId.ProtoReflect.Descriptor instead.
func (*UserIdWithUserTagId) Descriptor() ([]byte, []int) {
	return file_usertag_proto_rawDescGZIP(), []int{2}
}

func (x *UserIdWithUserTagId) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserIdWithUserTagId) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type RowsAffected struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RowsAffected int64 `protobuf:"varint,1,opt,name=rows_affected,json=rowsAffected,proto3" json:"rows_affected,omitempty"`
}

func (x *RowsAffected) Reset() {
	*x = RowsAffected{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usertag_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RowsAffected) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RowsAffected) ProtoMessage() {}

func (x *RowsAffected) ProtoReflect() protoreflect.Message {
	mi := &file_usertag_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RowsAffected.ProtoReflect.Descriptor instead.
func (*RowsAffected) Descriptor() ([]byte, []int) {
	return file_usertag_proto_rawDescGZIP(), []int{3}
}

func (x *RowsAffected) GetRowsAffected() int64 {
	if x != nil {
		return x.RowsAffected
	}
	return 0
}

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usertag_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_usertag_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_usertag_proto_rawDescGZIP(), []int{4}
}

func (x *CreateRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId int64  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name   string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usertag_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_usertag_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_usertag_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Page    int64 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	PerPage int64 `protobuf:"varint,3,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usertag_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_usertag_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_usertag_proto_rawDescGZIP(), []int{6}
}

func (x *ListRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ListRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListRequest) GetPerPage() int64 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total       int64      `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	PerPage     int64      `protobuf:"varint,2,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
	CurrentPage int64      `protobuf:"varint,3,opt,name=current_page,json=currentPage,proto3" json:"current_page,omitempty"`
	LastPage    int64      `protobuf:"varint,4,opt,name=last_page,json=lastPage,proto3" json:"last_page,omitempty"`
	From        int64      `protobuf:"varint,5,opt,name=from,proto3" json:"from,omitempty"`
	To          int64      `protobuf:"varint,6,opt,name=to,proto3" json:"to,omitempty"`
	Data        []*UserTag `protobuf:"bytes,7,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usertag_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_usertag_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_usertag_proto_rawDescGZIP(), []int{7}
}

func (x *ListResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListResponse) GetPerPage() int64 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

func (x *ListResponse) GetCurrentPage() int64 {
	if x != nil {
		return x.CurrentPage
	}
	return 0
}

func (x *ListResponse) GetLastPage() int64 {
	if x != nil {
		return x.LastPage
	}
	return 0
}

func (x *ListResponse) GetFrom() int64 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *ListResponse) GetTo() int64 {
	if x != nil {
		return x.To
	}
	return 0
}

func (x *ListResponse) GetData() []*UserTag {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_usertag_proto protoreflect.FileDescriptor

var file_usertag_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x12, 0x75, 0x73, 0x65, 0x72, 0x74, 0x61, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbc, 0x01, 0x0a, 0x07,
	0x55, 0x73, 0x65, 0x72, 0x54, 0x61, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x1b, 0x0a, 0x09, 0x55, 0x73,
	0x65, 0x72, 0x54, 0x61, 0x67, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3e, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x57, 0x69, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x54, 0x61, 0x67, 0x49, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x33, 0x0a, 0x0c, 0x52, 0x6f, 0x77, 0x73, 0x41,
	0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x6f, 0x77, 0x73, 0x5f,
	0x61, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c,
	0x72, 0x6f, 0x77, 0x73, 0x41, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x22, 0x51, 0x0a, 0x0d,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1e, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa,
	0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0xff, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x6a, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22,
	0x02, 0x20, 0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05,
	0x10, 0x01, 0x18, 0xff, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x55, 0x0a, 0x0b, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x65, 0x72, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61,
	0x67, 0x65, 0x22, 0xd4, 0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x65, 0x72,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x70, 0x65, 0x72,
	0x50, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x50, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74,
	0x50, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x2f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x74, 0x61, 0x67,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x54, 0x61, 0x67, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xd7, 0x04, 0x0a, 0x0e, 0x55, 0x73,
	0x65, 0x72, 0x54, 0x61, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6f, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x21, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x74, 0x61, 0x67,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x74, 0x61, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x54, 0x61, 0x67, 0x49, 0x64, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d,
	0x22, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x74, 0x61, 0x67, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x72, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x27, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x74, 0x61, 0x67, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x57, 0x69, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x54, 0x61, 0x67, 0x49, 0x64, 0x1a, 0x1b, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x74, 0x61, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x54, 0x61, 0x67, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1f, 0x12, 0x1d, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x74, 0x61, 0x67, 0x73, 0x2f, 0x7b, 0x69, 0x64,
	0x7d, 0x12, 0x77, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x21, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x74, 0x61, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x74, 0x61, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x77, 0x73, 0x41, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x1a, 0x1d, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x74, 0x61,
	0x67, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x7a, 0x0a, 0x06, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x27, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x74, 0x61, 0x67, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x57, 0x69, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x54, 0x61, 0x67, 0x49, 0x64, 0x1a, 0x20, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x74, 0x61, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x6f, 0x77, 0x73, 0x41, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x22,
	0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x2a, 0x1d, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x74, 0x61, 0x67,
	0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x6b, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x74, 0x61, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x20, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x74, 0x61, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x74,
	0x61, 0x67, 0x73, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6d, 0x69, 0x69, 0x79, 0x2f, 0x75, 0x70, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x74, 0x61, 0x67, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_usertag_proto_rawDescOnce sync.Once
	file_usertag_proto_rawDescData = file_usertag_proto_rawDesc
)

func file_usertag_proto_rawDescGZIP() []byte {
	file_usertag_proto_rawDescOnce.Do(func() {
		file_usertag_proto_rawDescData = protoimpl.X.CompressGZIP(file_usertag_proto_rawDescData)
	})
	return file_usertag_proto_rawDescData
}

var file_usertag_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_usertag_proto_goTypes = []interface{}{
	(*UserTag)(nil),             // 0: usertag_service.v1.UserTag
	(*UserTagId)(nil),           // 1: usertag_service.v1.UserTagId
	(*UserIdWithUserTagId)(nil), // 2: usertag_service.v1.UserIdWithUserTagId
	(*RowsAffected)(nil),        // 3: usertag_service.v1.RowsAffected
	(*CreateRequest)(nil),       // 4: usertag_service.v1.CreateRequest
	(*UpdateRequest)(nil),       // 5: usertag_service.v1.UpdateRequest
	(*ListRequest)(nil),         // 6: usertag_service.v1.ListRequest
	(*ListResponse)(nil),        // 7: usertag_service.v1.ListResponse
	(*timestamp.Timestamp)(nil), // 8: google.protobuf.Timestamp
}
var file_usertag_proto_depIdxs = []int32{
	8, // 0: usertag_service.v1.UserTag.created_at:type_name -> google.protobuf.Timestamp
	8, // 1: usertag_service.v1.UserTag.updated_at:type_name -> google.protobuf.Timestamp
	0, // 2: usertag_service.v1.ListResponse.data:type_name -> usertag_service.v1.UserTag
	4, // 3: usertag_service.v1.UserTagService.Create:input_type -> usertag_service.v1.CreateRequest
	2, // 4: usertag_service.v1.UserTagService.Get:input_type -> usertag_service.v1.UserIdWithUserTagId
	5, // 5: usertag_service.v1.UserTagService.Update:input_type -> usertag_service.v1.UpdateRequest
	2, // 6: usertag_service.v1.UserTagService.Delete:input_type -> usertag_service.v1.UserIdWithUserTagId
	6, // 7: usertag_service.v1.UserTagService.List:input_type -> usertag_service.v1.ListRequest
	1, // 8: usertag_service.v1.UserTagService.Create:output_type -> usertag_service.v1.UserTagId
	0, // 9: usertag_service.v1.UserTagService.Get:output_type -> usertag_service.v1.UserTag
	3, // 10: usertag_service.v1.UserTagService.Update:output_type -> usertag_service.v1.RowsAffected
	3, // 11: usertag_service.v1.UserTagService.Delete:output_type -> usertag_service.v1.RowsAffected
	7, // 12: usertag_service.v1.UserTagService.List:output_type -> usertag_service.v1.ListResponse
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_usertag_proto_init() }
func file_usertag_proto_init() {
	if File_usertag_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_usertag_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserTag); i {
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
		file_usertag_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserTagId); i {
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
		file_usertag_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserIdWithUserTagId); i {
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
		file_usertag_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RowsAffected); i {
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
		file_usertag_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_usertag_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_usertag_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_usertag_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
			RawDescriptor: file_usertag_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_usertag_proto_goTypes,
		DependencyIndexes: file_usertag_proto_depIdxs,
		MessageInfos:      file_usertag_proto_msgTypes,
	}.Build()
	File_usertag_proto = out.File
	file_usertag_proto_rawDesc = nil
	file_usertag_proto_goTypes = nil
	file_usertag_proto_depIdxs = nil
}
