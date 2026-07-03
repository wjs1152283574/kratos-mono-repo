package server

import (
	"context"

	uv1 "casso/gen/user/v1"
	"casso/pkg/registry"

	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewUserServiceClient 创建 user 服务 gRPC 客户端（BFF 通过 Nacos 发现 user 服务）
func NewUserServiceClient(nc registry.NacosConfig) uv1.UserClient {
	r := registry.NewDiscovery(nc)
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///casso.user.service.grpc"),
		grpc.WithDiscovery(r),
	)
	if err != nil {
		panic(err)
	}
	return uv1.NewUserClient(conn)
}
