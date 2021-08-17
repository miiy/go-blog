// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FeedbackServiceClient is the client API for FeedbackService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FeedbackServiceClient interface {
	Create(ctx context.Context, in *CreateFeedback, opts ...grpc.CallOption) (*FeedbackId, error)
	Delete(ctx context.Context, in *FeedbackId, opts ...grpc.CallOption) (*RowsAffected, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
}

type feedbackServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFeedbackServiceClient(cc grpc.ClientConnInterface) FeedbackServiceClient {
	return &feedbackServiceClient{cc}
}

func (c *feedbackServiceClient) Create(ctx context.Context, in *CreateFeedback, opts ...grpc.CallOption) (*FeedbackId, error) {
	out := new(FeedbackId)
	err := c.cc.Invoke(ctx, "/feedback.FeedbackService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedbackServiceClient) Delete(ctx context.Context, in *FeedbackId, opts ...grpc.CallOption) (*RowsAffected, error) {
	out := new(RowsAffected)
	err := c.cc.Invoke(ctx, "/feedback.FeedbackService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedbackServiceClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/feedback.FeedbackService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeedbackServiceServer is the server API for FeedbackService service.
// All implementations must embed UnimplementedFeedbackServiceServer
// for forward compatibility
type FeedbackServiceServer interface {
	Create(context.Context, *CreateFeedback) (*FeedbackId, error)
	Delete(context.Context, *FeedbackId) (*RowsAffected, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	mustEmbedUnimplementedFeedbackServiceServer()
}

// UnimplementedFeedbackServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFeedbackServiceServer struct {
}

func (UnimplementedFeedbackServiceServer) Create(context.Context, *CreateFeedback) (*FeedbackId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedFeedbackServiceServer) Delete(context.Context, *FeedbackId) (*RowsAffected, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedFeedbackServiceServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedFeedbackServiceServer) mustEmbedUnimplementedFeedbackServiceServer() {}

// UnsafeFeedbackServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FeedbackServiceServer will
// result in compilation errors.
type UnsafeFeedbackServiceServer interface {
	mustEmbedUnimplementedFeedbackServiceServer()
}

func RegisterFeedbackServiceServer(s grpc.ServiceRegistrar, srv FeedbackServiceServer) {
	s.RegisterService(&FeedbackService_ServiceDesc, srv)
}

func _FeedbackService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFeedback)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedbackServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feedback.FeedbackService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedbackServiceServer).Create(ctx, req.(*CreateFeedback))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedbackService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedbackId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedbackServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feedback.FeedbackService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedbackServiceServer).Delete(ctx, req.(*FeedbackId))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedbackService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedbackServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feedback.FeedbackService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedbackServiceServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FeedbackService_ServiceDesc is the grpc.ServiceDesc for FeedbackService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FeedbackService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "feedback.FeedbackService",
	HandlerType: (*FeedbackServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _FeedbackService_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _FeedbackService_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _FeedbackService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "feedback.proto",
}
