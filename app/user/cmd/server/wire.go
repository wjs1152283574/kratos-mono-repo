//go:build wireinject
// +build wireinject

package main

import (
	"casso/app/user/internal/biz"
	"casso/app/user/internal/conf"
	"casso/app/user/internal/data"
	"casso/app/user/internal/server"
	"casso/app/user/internal/service"

	nacosreg "casso/pkg/registry"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, log.Logger, *tracesdk.TracerProvider, nacosreg.NacosConfig) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
