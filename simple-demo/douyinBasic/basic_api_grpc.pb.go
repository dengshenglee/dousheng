// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: basic_api.proto

package douyinBasic

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

// BasicApiServiceClient is the client API for BasicApiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BasicApiServiceClient interface {
	Feed(ctx context.Context, in *DouyinFeedRequest, opts ...grpc.CallOption) (*DouyinFeedResponse, error)
	UserRegister(ctx context.Context, in *DouyinUserRegisterRequest, opts ...grpc.CallOption) (*DouyinUserRegisterResponse, error)
	UserLogin(ctx context.Context, in *DouyinUserLoginRequest, opts ...grpc.CallOption) (*DouyinUserLoginResponse, error)
	User(ctx context.Context, in *DouyinUserRequest, opts ...grpc.CallOption) (*DouyinUserResponse, error)
	PublishAction(ctx context.Context, in *DouyinPublishActionRequest, opts ...grpc.CallOption) (*DouyinPublishActionResponse, error)
	PublishList(ctx context.Context, in *DouyinPublishListRequest, opts ...grpc.CallOption) (*DouyinPublishListResponse, error)
}

type basicApiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBasicApiServiceClient(cc grpc.ClientConnInterface) BasicApiServiceClient {
	return &basicApiServiceClient{cc}
}

func (c *basicApiServiceClient) Feed(ctx context.Context, in *DouyinFeedRequest, opts ...grpc.CallOption) (*DouyinFeedResponse, error) {
	out := new(DouyinFeedResponse)
	err := c.cc.Invoke(ctx, "/douyinBasic.BasicApiService/Feed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicApiServiceClient) UserRegister(ctx context.Context, in *DouyinUserRegisterRequest, opts ...grpc.CallOption) (*DouyinUserRegisterResponse, error) {
	out := new(DouyinUserRegisterResponse)
	err := c.cc.Invoke(ctx, "/douyinBasic.BasicApiService/UserRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicApiServiceClient) UserLogin(ctx context.Context, in *DouyinUserLoginRequest, opts ...grpc.CallOption) (*DouyinUserLoginResponse, error) {
	out := new(DouyinUserLoginResponse)
	err := c.cc.Invoke(ctx, "/douyinBasic.BasicApiService/UserLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicApiServiceClient) User(ctx context.Context, in *DouyinUserRequest, opts ...grpc.CallOption) (*DouyinUserResponse, error) {
	out := new(DouyinUserResponse)
	err := c.cc.Invoke(ctx, "/douyinBasic.BasicApiService/User", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicApiServiceClient) PublishAction(ctx context.Context, in *DouyinPublishActionRequest, opts ...grpc.CallOption) (*DouyinPublishActionResponse, error) {
	out := new(DouyinPublishActionResponse)
	err := c.cc.Invoke(ctx, "/douyinBasic.BasicApiService/PublishAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicApiServiceClient) PublishList(ctx context.Context, in *DouyinPublishListRequest, opts ...grpc.CallOption) (*DouyinPublishListResponse, error) {
	out := new(DouyinPublishListResponse)
	err := c.cc.Invoke(ctx, "/douyinBasic.BasicApiService/PublishList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BasicApiServiceServer is the server API for BasicApiService service.
// All implementations must embed UnimplementedBasicApiServiceServer
// for forward compatibility
type BasicApiServiceServer interface {
	Feed(context.Context, *DouyinFeedRequest) (*DouyinFeedResponse, error)
	UserRegister(context.Context, *DouyinUserRegisterRequest) (*DouyinUserRegisterResponse, error)
	UserLogin(context.Context, *DouyinUserLoginRequest) (*DouyinUserLoginResponse, error)
	User(context.Context, *DouyinUserRequest) (*DouyinUserResponse, error)
	PublishAction(context.Context, *DouyinPublishActionRequest) (*DouyinPublishActionResponse, error)
	PublishList(context.Context, *DouyinPublishListRequest) (*DouyinPublishListResponse, error)
	mustEmbedUnimplementedBasicApiServiceServer()
}

// UnimplementedBasicApiServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBasicApiServiceServer struct {
}

func (UnimplementedBasicApiServiceServer) Feed(context.Context, *DouyinFeedRequest) (*DouyinFeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Feed not implemented")
}
func (UnimplementedBasicApiServiceServer) UserRegister(context.Context, *DouyinUserRegisterRequest) (*DouyinUserRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRegister not implemented")
}
func (UnimplementedBasicApiServiceServer) UserLogin(context.Context, *DouyinUserLoginRequest) (*DouyinUserLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func (UnimplementedBasicApiServiceServer) User(context.Context, *DouyinUserRequest) (*DouyinUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method User not implemented")
}
func (UnimplementedBasicApiServiceServer) PublishAction(context.Context, *DouyinPublishActionRequest) (*DouyinPublishActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishAction not implemented")
}
func (UnimplementedBasicApiServiceServer) PublishList(context.Context, *DouyinPublishListRequest) (*DouyinPublishListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishList not implemented")
}
func (UnimplementedBasicApiServiceServer) mustEmbedUnimplementedBasicApiServiceServer() {}

// UnsafeBasicApiServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BasicApiServiceServer will
// result in compilation errors.
type UnsafeBasicApiServiceServer interface {
	mustEmbedUnimplementedBasicApiServiceServer()
}

func RegisterBasicApiServiceServer(s grpc.ServiceRegistrar, srv BasicApiServiceServer) {
	s.RegisterService(&BasicApiService_ServiceDesc, srv)
}

func _BasicApiService_Feed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinFeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasicApiServiceServer).Feed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/douyinBasic.BasicApiService/Feed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasicApiServiceServer).Feed(ctx, req.(*DouyinFeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasicApiService_UserRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinUserRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasicApiServiceServer).UserRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/douyinBasic.BasicApiService/UserRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasicApiServiceServer).UserRegister(ctx, req.(*DouyinUserRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasicApiService_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinUserLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasicApiServiceServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/douyinBasic.BasicApiService/UserLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasicApiServiceServer).UserLogin(ctx, req.(*DouyinUserLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasicApiService_User_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasicApiServiceServer).User(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/douyinBasic.BasicApiService/User",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasicApiServiceServer).User(ctx, req.(*DouyinUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasicApiService_PublishAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinPublishActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasicApiServiceServer).PublishAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/douyinBasic.BasicApiService/PublishAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasicApiServiceServer).PublishAction(ctx, req.(*DouyinPublishActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasicApiService_PublishList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinPublishListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasicApiServiceServer).PublishList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/douyinBasic.BasicApiService/PublishList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasicApiServiceServer).PublishList(ctx, req.(*DouyinPublishListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BasicApiService_ServiceDesc is the grpc.ServiceDesc for BasicApiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BasicApiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "douyinBasic.BasicApiService",
	HandlerType: (*BasicApiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Feed",
			Handler:    _BasicApiService_Feed_Handler,
		},
		{
			MethodName: "UserRegister",
			Handler:    _BasicApiService_UserRegister_Handler,
		},
		{
			MethodName: "UserLogin",
			Handler:    _BasicApiService_UserLogin_Handler,
		},
		{
			MethodName: "User",
			Handler:    _BasicApiService_User_Handler,
		},
		{
			MethodName: "PublishAction",
			Handler:    _BasicApiService_PublishAction_Handler,
		},
		{
			MethodName: "PublishList",
			Handler:    _BasicApiService_PublishList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "basic_api.proto",
}