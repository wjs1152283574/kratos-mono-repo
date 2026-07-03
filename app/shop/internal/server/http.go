package server

import (
	"casso/app/shop/internal/conf"
	"casso/app/shop/internal/service"
	v1 "casso/gen/shop/v1"
	"casso/pkg/contextx"
	"casso/pkg/encoding"
	"casso/pkg/middleware"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/ratelimit"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, logger log.Logger, tp *tracesdk.TracerProvider, s *service.ShopService) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			selector.Server(
				recovery.Recovery(),
				tracing.Server(tracing.WithTracerProvider(tp)),
				logging.Server(logger),
				middleware.AuthFromQuery(string(contextx.UserID)),
			).Path("/api.shop.v1.Shop/GetUser").Build(),
			logging.Client(logger),
			ratelimit.Server(),
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

	// 跨域处理
	opts = append(opts, http.Filter(middleware.CORS()))

	// 自定义 JSON 响应编码
	opts = append(opts, http.ResponseEncoder(encoding.ResponseEncoder()))

	srv := http.NewServer(opts...)
	v1.RegisterShopHTTPServer(srv, s)

	return srv
}
