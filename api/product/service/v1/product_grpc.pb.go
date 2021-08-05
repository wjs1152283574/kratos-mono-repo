// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// ProductClient is the client API for Product service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductClient interface {
	CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*CreateProductReply, error)
	UpdateProduct(ctx context.Context, in *UpdateProductRequest, opts ...grpc.CallOption) (*UpdateProductReply, error)
	DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*DeleteProductReply, error)
	GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*GetProductReply, error)
	ListProduct(ctx context.Context, in *ListProductRequest, opts ...grpc.CallOption) (*ListProductReply, error)
	CreateCatalog(ctx context.Context, in *CreateCatalogRequest, opts ...grpc.CallOption) (*CreateCatalogReply, error)
	UpdateCatalog(ctx context.Context, in *UpdateCatalogRequest, opts ...grpc.CallOption) (*UpdateCatalogReply, error)
	DeleteCatalog(ctx context.Context, in *DeleteCatalogRequest, opts ...grpc.CallOption) (*DeleteCatalogReply, error)
	ListCatalog(ctx context.Context, in *ListCatalogRequest, opts ...grpc.CallOption) (*ListCatalogReply, error)
}

type productClient struct {
	cc grpc.ClientConnInterface
}

func NewProductClient(cc grpc.ClientConnInterface) ProductClient {
	return &productClient{cc}
}

func (c *productClient) CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*CreateProductReply, error) {
	out := new(CreateProductReply)
	err := c.cc.Invoke(ctx, "/api.product.service.v1.Product/CreateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) UpdateProduct(ctx context.Context, in *UpdateProductRequest, opts ...grpc.CallOption) (*UpdateProductReply, error) {
	out := new(UpdateProductReply)
	err := c.cc.Invoke(ctx, "/api.product.service.v1.Product/UpdateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*DeleteProductReply, error) {
	out := new(DeleteProductReply)
	err := c.cc.Invoke(ctx, "/api.product.service.v1.Product/DeleteProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*GetProductReply, error) {
	out := new(GetProductReply)
	err := c.cc.Invoke(ctx, "/api.product.service.v1.Product/GetProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) ListProduct(ctx context.Context, in *ListProductRequest, opts ...grpc.CallOption) (*ListProductReply, error) {
	out := new(ListProductReply)
	err := c.cc.Invoke(ctx, "/api.product.service.v1.Product/ListProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) CreateCatalog(ctx context.Context, in *CreateCatalogRequest, opts ...grpc.CallOption) (*CreateCatalogReply, error) {
	out := new(CreateCatalogReply)
	err := c.cc.Invoke(ctx, "/api.product.service.v1.Product/CreateCatalog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) UpdateCatalog(ctx context.Context, in *UpdateCatalogRequest, opts ...grpc.CallOption) (*UpdateCatalogReply, error) {
	out := new(UpdateCatalogReply)
	err := c.cc.Invoke(ctx, "/api.product.service.v1.Product/UpdateCatalog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) DeleteCatalog(ctx context.Context, in *DeleteCatalogRequest, opts ...grpc.CallOption) (*DeleteCatalogReply, error) {
	out := new(DeleteCatalogReply)
	err := c.cc.Invoke(ctx, "/api.product.service.v1.Product/DeleteCatalog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) ListCatalog(ctx context.Context, in *ListCatalogRequest, opts ...grpc.CallOption) (*ListCatalogReply, error) {
	out := new(ListCatalogReply)
	err := c.cc.Invoke(ctx, "/api.product.service.v1.Product/ListCatalog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServer is the server API for Product service.
// All implementations must embed UnimplementedProductServer
// for forward compatibility
type ProductServer interface {
	CreateProduct(context.Context, *CreateProductRequest) (*CreateProductReply, error)
	UpdateProduct(context.Context, *UpdateProductRequest) (*UpdateProductReply, error)
	DeleteProduct(context.Context, *DeleteProductRequest) (*DeleteProductReply, error)
	GetProduct(context.Context, *GetProductRequest) (*GetProductReply, error)
	ListProduct(context.Context, *ListProductRequest) (*ListProductReply, error)
	CreateCatalog(context.Context, *CreateCatalogRequest) (*CreateCatalogReply, error)
	UpdateCatalog(context.Context, *UpdateCatalogRequest) (*UpdateCatalogReply, error)
	DeleteCatalog(context.Context, *DeleteCatalogRequest) (*DeleteCatalogReply, error)
	ListCatalog(context.Context, *ListCatalogRequest) (*ListCatalogReply, error)
	mustEmbedUnimplementedProductServer()
}

// UnimplementedProductServer must be embedded to have forward compatible implementations.
type UnimplementedProductServer struct {
}

func (UnimplementedProductServer) CreateProduct(context.Context, *CreateProductRequest) (*CreateProductReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (UnimplementedProductServer) UpdateProduct(context.Context, *UpdateProductRequest) (*UpdateProductReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProduct not implemented")
}
func (UnimplementedProductServer) DeleteProduct(context.Context, *DeleteProductRequest) (*DeleteProductReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProduct not implemented")
}
func (UnimplementedProductServer) GetProduct(context.Context, *GetProductRequest) (*GetProductReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedProductServer) ListProduct(context.Context, *ListProductRequest) (*ListProductReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProduct not implemented")
}
func (UnimplementedProductServer) CreateCatalog(context.Context, *CreateCatalogRequest) (*CreateCatalogReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCatalog not implemented")
}
func (UnimplementedProductServer) UpdateCatalog(context.Context, *UpdateCatalogRequest) (*UpdateCatalogReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCatalog not implemented")
}
func (UnimplementedProductServer) DeleteCatalog(context.Context, *DeleteCatalogRequest) (*DeleteCatalogReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCatalog not implemented")
}
func (UnimplementedProductServer) ListCatalog(context.Context, *ListCatalogRequest) (*ListCatalogReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCatalog not implemented")
}
func (UnimplementedProductServer) mustEmbedUnimplementedProductServer() {}

// UnsafeProductServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServer will
// result in compilation errors.
type UnsafeProductServer interface {
	mustEmbedUnimplementedProductServer()
}

func RegisterProductServer(s grpc.ServiceRegistrar, srv ProductServer) {
	s.RegisterService(&Product_ServiceDesc, srv)
}

func _Product_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.product.service.v1.Product/CreateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).CreateProduct(ctx, req.(*CreateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_UpdateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).UpdateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.product.service.v1.Product/UpdateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).UpdateProduct(ctx, req.(*UpdateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.product.service.v1.Product/DeleteProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).DeleteProduct(ctx, req.(*DeleteProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.product.service.v1.Product/GetProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).GetProduct(ctx, req.(*GetProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_ListProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).ListProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.product.service.v1.Product/ListProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).ListProduct(ctx, req.(*ListProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_CreateCatalog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCatalogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).CreateCatalog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.product.service.v1.Product/CreateCatalog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).CreateCatalog(ctx, req.(*CreateCatalogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_UpdateCatalog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCatalogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).UpdateCatalog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.product.service.v1.Product/UpdateCatalog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).UpdateCatalog(ctx, req.(*UpdateCatalogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_DeleteCatalog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCatalogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).DeleteCatalog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.product.service.v1.Product/DeleteCatalog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).DeleteCatalog(ctx, req.(*DeleteCatalogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_ListCatalog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCatalogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).ListCatalog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.product.service.v1.Product/ListCatalog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).ListCatalog(ctx, req.(*ListCatalogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Product_ServiceDesc is the grpc.ServiceDesc for Product service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Product_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.product.service.v1.Product",
	HandlerType: (*ProductServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProduct",
			Handler:    _Product_CreateProduct_Handler,
		},
		{
			MethodName: "UpdateProduct",
			Handler:    _Product_UpdateProduct_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _Product_DeleteProduct_Handler,
		},
		{
			MethodName: "GetProduct",
			Handler:    _Product_GetProduct_Handler,
		},
		{
			MethodName: "ListProduct",
			Handler:    _Product_ListProduct_Handler,
		},
		{
			MethodName: "CreateCatalog",
			Handler:    _Product_CreateCatalog_Handler,
		},
		{
			MethodName: "UpdateCatalog",
			Handler:    _Product_UpdateCatalog_Handler,
		},
		{
			MethodName: "DeleteCatalog",
			Handler:    _Product_DeleteCatalog_Handler,
		},
		{
			MethodName: "ListCatalog",
			Handler:    _Product_ListCatalog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/product/service/v1/product.proto",
}
