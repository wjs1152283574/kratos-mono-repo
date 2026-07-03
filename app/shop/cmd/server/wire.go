//go:build wireinject
// +build wireinject

package main

import (
	"casso/app/shop/internal/biz"
	"casso/app/shop/internal/conf"
	"casso/app/shop/internal/server"
	"casso/app/shop/internal/service"

	"casso/pkg/registry"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// initApp init kratos application.
func initApp(*conf.Server, log.Logger, *tracesdk.TracerProvider, registry.NacosConfig) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
