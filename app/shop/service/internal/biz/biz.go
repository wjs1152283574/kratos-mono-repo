package biz

import (
	v1 "casso/api/user/service/v1"
	"casso/app/shop/service/internal/conf"
	"context"

	nr "github.com/go-kratos/nacos/registry"

	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewShopUseCase, NewUserServiceClient, NewDiscovery)

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
func NewUserServiceClient(r registry.Discovery) v1.UserClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///casso.user.service.grpc"), // 三个`/`省略掉/default/
		grpc.WithDiscovery(r),
	)
	if err != nil {
		panic(err)
	}
	return v1.NewUserClient(conn)
}
