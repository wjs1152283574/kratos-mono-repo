// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: api/shop/service/v1/shop.proto

package v1

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

// ShopClient is the client API for Shop service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShopClient interface {
	// body:"*" 表示：请求数据全部从请求体映射，可以指定需要映射的字段
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserReply, error)
	Demo(ctx context.Context, in *DemoRequest, opts ...grpc.CallOption) (*DemoResponse, error)
}

type shopClient struct {
	cc grpc.ClientConnInterface
}

func NewShopClient(cc grpc.ClientConnInterface) ShopClient {
	return &shopClient{cc}
}

func (c *shopClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error) {
	out := new(RegisterReply)
	err := c.cc.Invoke(ctx, "/api.shop.service.v1.Shop/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, "/api.shop.service.v1.Shop/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserReply, error) {
	out := new(GetUserReply)
	err := c.cc.Invoke(ctx, "/api.shop.service.v1.Shop/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopClient) Demo(ctx context.Context, in *DemoRequest, opts ...grpc.CallOption) (*DemoResponse, error) {
	out := new(DemoResponse)
	err := c.cc.Invoke(ctx, "/api.shop.service.v1.Shop/Demo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShopServer is the server API for Shop service.
// All implementations must embed UnimplementedShopServer
// for forward compatibility
type ShopServer interface {
	// body:"*" 表示：请求数据全部从请求体映射，可以指定需要映射的字段
	Register(context.Context, *RegisterRequest) (*RegisterReply, error)
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserReply, error)
	Demo(context.Context, *DemoRequest) (*DemoResponse, error)
	mustEmbedUnimplementedShopServer()
}

// UnimplementedShopServer must be embedded to have forward compatible implementations.
type UnimplementedShopServer struct {
}

func (UnimplementedShopServer) Register(context.Context, *RegisterRequest) (*RegisterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedShopServer) Login(context.Context, *LoginRequest) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedShopServer) GetUser(context.Context, *GetUserRequest) (*GetUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedShopServer) Demo(context.Context, *DemoRequest) (*DemoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Demo not implemented")
}
func (UnimplementedShopServer) mustEmbedUnimplementedShopServer() {}

// UnsafeShopServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShopServer will
// result in compilation errors.
type UnsafeShopServer interface {
	mustEmbedUnimplementedShopServer()
}

func RegisterShopServer(s grpc.ServiceRegistrar, srv ShopServer) {
	s.RegisterService(&Shop_ServiceDesc, srv)
}

func _Shop_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.shop.service.v1.Shop/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shop_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.shop.service.v1.Shop/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shop_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.shop.service.v1.Shop/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shop_Demo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DemoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServer).Demo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.shop.service.v1.Shop/Demo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServer).Demo(ctx, req.(*DemoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Shop_ServiceDesc is the grpc.ServiceDesc for Shop service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Shop_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.shop.service.v1.Shop",
	HandlerType: (*ShopServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Shop_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Shop_Login_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _Shop_GetUser_Handler,
		},
		{
			MethodName: "Demo",
			Handler:    _Shop_Demo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/shop/service/v1/shop.proto",
}
