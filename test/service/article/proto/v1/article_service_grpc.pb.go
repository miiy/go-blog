// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: article_service.proto

package proto

import (
	__ "./"
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserPostServiceClient is the client API for UserPostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserPostServiceClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*__.Result, error)
	Get(ctx context.Context, in *__.Id, opts ...grpc.CallOption) (*__.Result, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*__.Result, error)
	Delete(ctx context.Context, in *UserIdWithArticleId, opts ...grpc.CallOption) (*__.Result, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
}

type userPostServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserPostServiceClient(cc grpc.ClientConnInterface) UserPostServiceClient {
	return &userPostServiceClient{cc}
}

func (c *userPostServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*__.Result, error) {
	out := new(__.Result)
	err := c.cc.Invoke(ctx, "/article_service.proto.UserPostService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userPostServiceClient) Get(ctx context.Context, in *__.Id, opts ...grpc.CallOption) (*__.Result, error) {
	out := new(__.Result)
	err := c.cc.Invoke(ctx, "/article_service.proto.UserPostService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userPostServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*__.Result, error) {
	out := new(__.Result)
	err := c.cc.Invoke(ctx, "/article_service.proto.UserPostService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userPostServiceClient) Delete(ctx context.Context, in *UserIdWithArticleId, opts ...grpc.CallOption) (*__.Result, error) {
	out := new(__.Result)
	err := c.cc.Invoke(ctx, "/article_service.proto.UserPostService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userPostServiceClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/article_service.proto.UserPostService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserPostServiceServer is the server API for UserPostService service.
// All implementations must embed UnimplementedUserPostServiceServer
// for forward compatibility
type UserPostServiceServer interface {
	Create(context.Context, *CreateRequest) (*__.Result, error)
	Get(context.Context, *__.Id) (*__.Result, error)
	Update(context.Context, *UpdateRequest) (*__.Result, error)
	Delete(context.Context, *UserIdWithArticleId) (*__.Result, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	mustEmbedUnimplementedUserPostServiceServer()
}

// UnimplementedUserPostServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserPostServiceServer struct {
}

func (UnimplementedUserPostServiceServer) Create(context.Context, *CreateRequest) (*__.Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedUserPostServiceServer) Get(context.Context, *__.Id) (*__.Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedUserPostServiceServer) Update(context.Context, *UpdateRequest) (*__.Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedUserPostServiceServer) Delete(context.Context, *UserIdWithArticleId) (*__.Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedUserPostServiceServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedUserPostServiceServer) mustEmbedUnimplementedUserPostServiceServer() {}

// UnsafeUserPostServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserPostServiceServer will
// result in compilation errors.
type UnsafeUserPostServiceServer interface {
	mustEmbedUnimplementedUserPostServiceServer()
}

func RegisterUserPostServiceServer(s grpc.ServiceRegistrar, srv UserPostServiceServer) {
	s.RegisterService(&UserPostService_ServiceDesc, srv)
}

func _UserPostService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserPostServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/article_service.proto.UserPostService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserPostServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserPostService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(__.Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserPostServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/article_service.proto.UserPostService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserPostServiceServer).Get(ctx, req.(*__.Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserPostService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserPostServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/article_service.proto.UserPostService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserPostServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserPostService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdWithArticleId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserPostServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/article_service.proto.UserPostService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserPostServiceServer).Delete(ctx, req.(*UserIdWithArticleId))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserPostService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserPostServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/article_service.proto.UserPostService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserPostServiceServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserPostService_ServiceDesc is the grpc.ServiceDesc for UserPostService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserPostService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "article_service.proto.UserPostService",
	HandlerType: (*UserPostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _UserPostService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _UserPostService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _UserPostService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _UserPostService_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _UserPostService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "article_service.proto",
}
