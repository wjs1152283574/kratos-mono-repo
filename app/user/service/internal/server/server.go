/*
 * @PackageName: server
 * @Description: server register
 * @Author: Casso
 * @Date: 2021-08-05 11:22:07
 * @LastModifiedBy: Casso
 * @LastEditTime: 2022-02-05 16:33:52
 */
package server

import (
	"casso/app/user/service/internal/conf"
	"log"

	"github.com/google/wire"

	kr "github.com/go-kratos/kratos/v2/registry"
	nr "github.com/go-kratos/nacos/registry"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewRegistrar, NewGRPCServer)

func NewRegistrar(conf *conf.Registry) kr.Registrar {
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
		log.Panic(err)
	}
	r := nr.New(client)
	return r
}
