package registry

import (
	"log"

	kr "github.com/go-kratos/kratos/v2/registry"
	nr "github.com/go-kratos/nacos/registry"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

// NacosConfig Nacos 配置
type NacosConfig struct {
	Address string
	Port    uint64
}

// NewDiscovery 创建 Nacos 服务发现
func NewDiscovery(conf NacosConfig) kr.Discovery {
	client, _ := newNamingClient(conf)
	return nr.New(client)
}

// NewRegistrar 创建 Nacos 服务注册
func NewRegistrar(conf NacosConfig) kr.Registrar {
	client, _ := newNamingClient(conf)
	return nr.New(client)
}

func newNamingClient(conf NacosConfig) (naming_client.INamingClient, error) {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(conf.Address, conf.Port),
	}

	cc := &constant.ClientConfig{
		NamespaceId:         "public",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            "debug",
	}

	client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		log.Panicf("failed to create nacos naming client: %v", err)
	}
	return client, nil
}
