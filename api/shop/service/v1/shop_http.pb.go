// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.0.3

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type ShopHTTPServer interface {
	Login(context.Context, *LoginRequest) (*LoginReply, error)
}

func RegisterShopHTTPServer(s *http.Server, srv ShopHTTPServer) {
	r := s.Route("/")
	r.POST("/login", _Shop_Login0_HTTP_Handler(srv))
}

func _Shop_Login0_HTTP_Handler(srv ShopHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.shop.service.v1.Shop/Login")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginReply)
		return ctx.Result(200, reply)
	}
}

type ShopHTTPClient interface {
	Login(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *LoginReply, err error)
}

type ShopHTTPClientImpl struct {
	cc *http.Client
}

func NewShopHTTPClient(client *http.Client) ShopHTTPClient {
	return &ShopHTTPClientImpl{client}
}

func (c *ShopHTTPClientImpl) Login(ctx context.Context, in *LoginRequest, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.shop.service.v1.Shop/Login"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}