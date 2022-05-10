/*
 * @Author: Casso
 * @Date: 2021-11-17 16:24:19
 * @LastEditors: Casso
 * @LastEditTime: 2021-11-22 10:28:12
 * @Description: file content
 * @FilePath: /kratos-mono-repo/app/shop/service/internal/server/grpc.go
 */
package server

import (
	v1 "casso/api/shop/service/v1"
	uv1 "casso/api/user/service/v1"
	"casso/app/shop/service/internal/conf"
	"casso/app/shop/service/internal/service"
	"context"

	nr "github.com/go-kratos/nacos/registry"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"

	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, logger log.Logger, tp *tracesdk.TracerProvider, s *service.ShopService) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			tracing.Server(
				tracing.WithTracerProvider(tp)),
			logging.Server(logger),
			logging.Client(logger),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterShopServer(srv, s)
	return srv
}

func NewDiscovery(conf *conf.Discovery) registry.Discovery {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(conf.Nacos.Address, uint64(conf.Nacos.Port)),
	}

	cc := &constant.ClientConfig{
		NamespaceId:         "public", //namespace id
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            "debug",
	}

	// a more graceful way to create naming client
	client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  cc,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		panic("err")
	}

	r := nr.New(client)

	return r
}

// NewUserServiceClient user service rpc client
func NewUserServiceClient(r registry.Discovery) uv1.UserClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///casso.user.service.grpc"), // 三个`/`省略掉/default/
		grpc.WithDiscovery(r),
	)
	if err != nil {
		panic(err)
	}
	return uv1.NewUserClient(conn)
}
