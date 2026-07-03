package server

import (
	"casso/pkg/registry"

	kr "github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewRegistrar, NewHTTPServer, NewUserServiceClient)

// NewRegistrar 创建 Nacos 注册器
func NewRegistrar(nc registry.NacosConfig) kr.Registrar {
	return registry.NewRegistrar(nc)
}
