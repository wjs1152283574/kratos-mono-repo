// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"casso/app/shop/service/internal/biz"
	"casso/app/shop/service/internal/conf"
	"casso/app/shop/service/internal/data"
	"casso/app/shop/service/internal/server"
	"casso/app/shop/service/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Discovery, *conf.Data, log.Logger, *tracesdk.TracerProvider) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
