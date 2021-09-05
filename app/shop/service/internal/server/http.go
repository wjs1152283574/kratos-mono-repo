package server

import (
	v1 "casso/api/shop/service/v1"
	"casso/app/shop/service/internal/conf"
	"casso/app/shop/service/internal/service"
	"casso/pkg/util/token"
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"

	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, logger log.Logger, tp *tracesdk.TracerProvider, s *service.ShopService) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(tracing.WithTracerProvider(tp)),
			logging.Server(logger),
			AuthMiddleware,
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}

	srv := http.NewServer(opts...)
	v1.RegisterShopHTTPServer(srv, s)
	return srv
}

func AuthMiddleware(handler middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
		reply, err = handler(ctx, req)
		if tr, ok := transport.FromServerContext(ctx); ok {
			kind := tr.Kind().String()
			operation := tr.Operation()
			fmt.Println(kind, operation)
			// 断言成HTTP的Transport可以拿到特殊信息
			if ht, ok := tr.(*http.Transport); ok && ht.Request().Header.Get("Authorization") != "" {
				uinfos, errs := token.NewJWT().ParseToken(ht.Request().Header.Get("Authorization"))
				if errs != nil {
					fmt.Println(err)
				}
				etxs := context.WithValue(ctx, "uid", uinfos.ID)
				reply, err = handler(etxs, req)
			}
		}
		return
	}
}
