// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: article_service.proto

package proto

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Article_ArticleStatus int32

const (
	Article_DEFAULT Article_ArticleStatus = 0
	Article_ACTIVE  Article_ArticleStatus = 1
	Article_DISABLE Article_ArticleStatus = 2
)

// Enum value maps for Article_ArticleStatus.
var (
	Article_ArticleStatus_name = map[int32]string{
		0: "DEFAULT",
		1: "ACTIVE",
		2: "DISABLE",
	}
	Article_ArticleStatus_value = map[string]int32{
		"DEFAULT": 0,
		"ACTIVE":  1,
		"DISABLE": 2,
	}
)

func (x Article_ArticleStatus) Enum() *Article_ArticleStatus {
	p := new(Article_ArticleStatus)
	*p = x
	return p
}

func (x Article_ArticleStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Article_ArticleStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_article_service_proto_enumTypes[0].Descriptor()
}

func (Article_ArticleStatus) Type() protoreflect.EnumType {
	return &file_article_service_proto_enumTypes[0]
}

func (x Article_ArticleStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Article_ArticleStatus.Descriptor instead.
func (Article_ArticleStatus) EnumDescriptor() ([]byte, []int) {
	return file_article_service_proto_rawDescGZIP(), []int{0, 0}
}

type Article struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId          int64                  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Title           string                 `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	MetaTitle       string                 `protobuf:"bytes,4,opt,name=meta_title,json=metaTitle,proto3" json:"meta_title,omitempty"`
	MetaDescription string                 `protobuf:"bytes,5,opt,name=meta_description,json=metaDescription,proto3" json:"meta_description,omitempty"`
	PublishedTime   *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=published_time,json=publishedTime,proto3" json:"published_time,omitempty"`
	UpdatedTime     *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=updated_time,json=updatedTime,proto3" json:"updated_time,omitempty"`
	FromText        string                 `protobuf:"bytes,8,opt,name=from_text,json=fromText,proto3" json:"from_text,omitempty"`
	FromUrl         string                 `protobuf:"bytes,9,opt,name=from_url,json=fromUrl,proto3" json:"from_url,omitempty"`
	Summary         string                 `protobuf:"bytes,10,opt,name=summary,proto3" json:"summary,omitempty"`
	Content         string                 `protobuf:"bytes,11,opt,name=content,proto3" json:"content,omitempty"`
	Status          Article_ArticleStatus  `protobuf:"varint,12,opt,name=status,proto3,enum=article_service.proto.Article_ArticleStatus" json:"status,omitempty"`
	CreatedAt       *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt       *timestamppb.Timestamp `protobuf:"bytes,14,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Article) Reset() {
	*x = Article{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Article) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Article) ProtoMessage() {}

func (x *Article) ProtoReflect() protoreflect.Message {
	mi := &file_article_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Article.ProtoReflect.Descriptor instead.
func (*Article) Descriptor() ([]byte, []int) {
	return file_article_service_proto_rawDescGZIP(), []int{0}
}

func (x *Article) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Article) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Article) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Article) GetMetaTitle() string {
	if x != nil {
		return x.MetaTitle
	}
	return ""
}

func (x *Article) GetMetaDescription() string {
	if x != nil {
		return x.MetaDescription
	}
	return ""
}

func (x *Article) GetPublishedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.PublishedTime
	}
	return nil
}

func (x *Article) GetUpdatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedTime
	}
	return nil
}

func (x *Article) GetFromText() string {
	if x != nil {
		return x.FromText
	}
	return ""
}

func (x *Article) GetFromUrl() string {
	if x != nil {
		return x.FromUrl
	}
	return ""
}

func (x *Article) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *Article) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Article) GetStatus() Article_ArticleStatus {
	if x != nil {
		return x.Status
	}
	return Article_DEFAULT
}

func (x *Article) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Article) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type ArticleId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ArticleId) Reset() {
	*x = ArticleId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleId) ProtoMessage() {}

func (x *ArticleId) ProtoReflect() protoreflect.Message {
	mi := &file_article_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleId.ProtoReflect.Descriptor instead.
func (*ArticleId) Descriptor() ([]byte, []int) {
	return file_article_service_proto_rawDescGZIP(), []int{1}
}

func (x *ArticleId) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UserIdWithArticleId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Id     int64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserIdWithArticleId) Reset() {
	*x = UserIdWithArticleId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserIdWithArticleId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserIdWithArticleId) ProtoMessage() {}

func (x *UserIdWithArticleId) ProtoReflect() protoreflect.Message {
	mi := &file_article_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserIdWithArticleId.ProtoReflect.Descriptor instead.
func (*UserIdWithArticleId) Descriptor() ([]byte, []int) {
	return file_article_service_proto_rawDescGZIP(), []int{2}
}

func (x *UserIdWithArticleId) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserIdWithArticleId) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId        int64                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content       string                 `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Status        Article_ArticleStatus  `protobuf:"varint,4,opt,name=status,proto3,enum=article_service.proto.Article_ArticleStatus" json:"status,omitempty"`
	PublishedTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=published_time,json=publishedTime,proto3" json:"published_time,omitempty"`
	UpdatedTime   *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_time,json=updatedTime,proto3" json:"updated_time,omitempty"`
	Sort          int64                  `protobuf:"varint,7,opt,name=sort,proto3" json:"sort,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_article_service_proto_msgTypes[3]
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
	return file_article_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreateRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CreateRequest) GetStatus() Article_ArticleStatus {
	if x != nil {
		return x.Status
	}
	return Article_DEFAULT
}

func (x *CreateRequest) GetPublishedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.PublishedTime
	}
	return nil
}

func (x *CreateRequest) GetUpdatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedTime
	}
	return nil
}

func (x *CreateRequest) GetSort() int64 {
	if x != nil {
		return x.Sort
	}
	return 0
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        int64                  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Title         string                 `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Content       string                 `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Status        int64                  `protobuf:"varint,5,opt,name=status,proto3" json:"status,omitempty"`
	PublishedTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=published_time,json=publishedTime,proto3" json:"published_time,omitempty"`
	UpdatedTime   *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=updated_time,json=updatedTime,proto3" json:"updated_time,omitempty"`
	Sort          int64                  `protobuf:"varint,8,opt,name=sort,proto3" json:"sort,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_article_service_proto_msgTypes[4]
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
	return file_article_service_proto_rawDescGZIP(), []int{4}
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

func (x *UpdateRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *UpdateRequest) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *UpdateRequest) GetPublishedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.PublishedTime
	}
	return nil
}

func (x *UpdateRequest) GetUpdatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedTime
	}
	return nil
}

func (x *UpdateRequest) GetSort() int64 {
	if x != nil {
		return x.Sort
	}
	return 0
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page    int64  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PerPage int64  `protobuf:"varint,2,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
	Keyword string `protobuf:"bytes,3,opt,name=keyword,proto3" json:"keyword,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_article_service_proto_msgTypes[5]
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
	return file_article_service_proto_rawDescGZIP(), []int{5}
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

func (x *ListRequest) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
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
	Data        []*Article `protobuf:"bytes,7,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_article_service_proto_msgTypes[6]
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
	return file_article_service_proto_rawDescGZIP(), []int{6}
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

func (x *ListResponse) GetData() []*Article {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_article_service_proto protoreflect.FileDescriptor

var file_article_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xf3, 0x04, 0x0a, 0x07, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x74, 0x61, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x29, 0x0a,
	0x10, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6d, 0x65, 0x74, 0x61, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x41, 0x0a, 0x0e, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x72,
	0x6f, 0x6d, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x72, 0x6f, 0x6d, 0x54, 0x65, 0x78, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x72, 0x6f, 0x6d, 0x55,
	0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x44, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x35, 0x0a, 0x0d, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x00,
	0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07,
	0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x02, 0x22, 0x1b, 0x0a, 0x09, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3e, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x57, 0x69, 0x74, 0x68, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0xdd, 0x02, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1e, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x18, 0xff, 0x01, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x24, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0xff, 0x01, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x4e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x41, 0x0a, 0x0e, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x04, 0x73, 0x6f, 0x72,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x28, 0x00,
	0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x22, 0x96, 0x02, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x41, 0x0a, 0x0e, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x0c,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x6f, 0x72, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x22,
	0x56, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x22, 0xd7, 0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x19,
	0x0a, 0x08, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x6c, 0x61, 0x73, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x6c, 0x61, 0x73, 0x74, 0x50, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f,
	0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a,
	0x02, 0x74, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x32, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x32, 0x94, 0x04, 0x0a, 0x0e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x61, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1e,
	0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x1a, 0x1e,
	0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x22, 0x17,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x67, 0x0a, 0x08, 0x52, 0x65, 0x74, 0x72, 0x69,
	0x65, 0x76, 0x65, 0x12, 0x20, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x49, 0x64, 0x1a, 0x1e, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x12, 0x66, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x1a, 0x1e, 0x2e, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x16, 0x1a, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73,
	0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x67, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x2a, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x57, 0x69, 0x74, 0x68, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x2a, 0x11,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64,
	0x7d, 0x12, 0x65, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x22, 0x2e, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x69, 0x69, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x62,
	0x6c, 0x6f, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_article_service_proto_rawDescOnce sync.Once
	file_article_service_proto_rawDescData = file_article_service_proto_rawDesc
)

func file_article_service_proto_rawDescGZIP() []byte {
	file_article_service_proto_rawDescOnce.Do(func() {
		file_article_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_article_service_proto_rawDescData)
	})
	return file_article_service_proto_rawDescData
}

var file_article_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_article_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_article_service_proto_goTypes = []interface{}{
	(Article_ArticleStatus)(0),    // 0: article_service.proto.Article.ArticleStatus
	(*Article)(nil),               // 1: article_service.proto.Article
	(*ArticleId)(nil),             // 2: article_service.proto.ArticleId
	(*UserIdWithArticleId)(nil),   // 3: article_service.proto.UserIdWithArticleId
	(*CreateRequest)(nil),         // 4: article_service.proto.CreateRequest
	(*UpdateRequest)(nil),         // 5: article_service.proto.UpdateRequest
	(*ListRequest)(nil),           // 6: article_service.proto.ListRequest
	(*ListResponse)(nil),          // 7: article_service.proto.ListResponse
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 9: google.protobuf.Empty
}
var file_article_service_proto_depIdxs = []int32{
	8,  // 0: article_service.proto.Article.published_time:type_name -> google.protobuf.Timestamp
	8,  // 1: article_service.proto.Article.updated_time:type_name -> google.protobuf.Timestamp
	0,  // 2: article_service.proto.Article.status:type_name -> article_service.proto.Article.ArticleStatus
	8,  // 3: article_service.proto.Article.created_at:type_name -> google.protobuf.Timestamp
	8,  // 4: article_service.proto.Article.updated_at:type_name -> google.protobuf.Timestamp
	0,  // 5: article_service.proto.CreateRequest.status:type_name -> article_service.proto.Article.ArticleStatus
	8,  // 6: article_service.proto.CreateRequest.published_time:type_name -> google.protobuf.Timestamp
	8,  // 7: article_service.proto.CreateRequest.updated_time:type_name -> google.protobuf.Timestamp
	8,  // 8: article_service.proto.UpdateRequest.published_time:type_name -> google.protobuf.Timestamp
	8,  // 9: article_service.proto.UpdateRequest.updated_time:type_name -> google.protobuf.Timestamp
	1,  // 10: article_service.proto.ListResponse.data:type_name -> article_service.proto.Article
	1,  // 11: article_service.proto.ArticleService.Create:input_type -> article_service.proto.Article
	2,  // 12: article_service.proto.ArticleService.Retrieve:input_type -> article_service.proto.ArticleId
	1,  // 13: article_service.proto.ArticleService.Update:input_type -> article_service.proto.Article
	3,  // 14: article_service.proto.ArticleService.Delete:input_type -> article_service.proto.UserIdWithArticleId
	6,  // 15: article_service.proto.ArticleService.List:input_type -> article_service.proto.ListRequest
	1,  // 16: article_service.proto.ArticleService.Create:output_type -> article_service.proto.Article
	1,  // 17: article_service.proto.ArticleService.Retrieve:output_type -> article_service.proto.Article
	1,  // 18: article_service.proto.ArticleService.Update:output_type -> article_service.proto.Article
	9,  // 19: article_service.proto.ArticleService.Delete:output_type -> google.protobuf.Empty
	7,  // 20: article_service.proto.ArticleService.List:output_type -> article_service.proto.ListResponse
	16, // [16:21] is the sub-list for method output_type
	11, // [11:16] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_article_service_proto_init() }
func file_article_service_proto_init() {
	if File_article_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_article_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Article); i {
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
		file_article_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleId); i {
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
		file_article_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserIdWithArticleId); i {
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
		file_article_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_article_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_article_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_article_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_article_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_article_service_proto_goTypes,
		DependencyIndexes: file_article_service_proto_depIdxs,
		EnumInfos:         file_article_service_proto_enumTypes,
		MessageInfos:      file_article_service_proto_msgTypes,
	}.Build()
	File_article_service_proto = out.File
	file_article_service_proto_rawDesc = nil
	file_article_service_proto_goTypes = nil
	file_article_service_proto_depIdxs = nil
}
