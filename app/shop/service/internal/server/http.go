/*
 * @Author: Casso
 * @Date: 2021-11-17 16:24:19
 * @LastEditors: Casso
 * @LastEditTime: 2021-11-29 10:41:45
 * @Description: file content
 * @FilePath: /kratos-mono-repo/app/shop/service/internal/server/http.go
 */
/*
 * @Author: Casso
 * @Date: 2021-11-17 16:24:19
 * @LastEditors: Casso
 * @LastEditTime: 2021-11-20 16:32:10
 * @Description: file content
 * @FilePath: /kratos-mono-repo/app/shop/service/internal/server/http.go
 */
package server

import (
	v1 "casso/api/shop/service/v1"
	"casso/app/shop/service/internal/conf"
	"casso/app/shop/service/internal/service"
	"casso/pkg/errors/auth"
	"casso/pkg/util/contextkey"
	"casso/pkg/util/resencoder"
	"casso/pkg/util/token"
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"

	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, logger log.Logger, tp *tracesdk.TracerProvider, s *service.ShopService) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			selector.Server(
				recovery.Recovery(),
				tracing.Server(tracing.WithTracerProvider(tp)),
				logging.Server(logger),
				AuthMiddleware,
			).Path("/api.shop.service.v1.Shop/GetUser").Build(),
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

	// 自定义返回数据编码方式
	opts = append(opts, http.ResponseEncoder(resencoder.CustomResponeDeco())) // https://mp.weixin.qq.com/s/4ocdoAVXXKTvJ3U65YXltw

	srv := http.NewServer(opts...)
	v1.RegisterShopHTTPServer(srv, s)
	return srv
}

// AuthMiddleware BFF服务内鉴权中间件，这里只是暂时的。实际情况会在gate验证解析token，鉴权信息放入uri到达BFF
func AuthMiddleware(handler middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
		reply, err = handler(ctx, req)
		if tr, ok := transport.FromServerContext(ctx); ok {
			ht, ok := tr.(*http.Transport)
			if !ok && ht.Request().Header.Get("Authorization") == "" {
				return nil, auth.ErrAuthFail
			}

			uinfos, parserErr := token.NewJWT().ParseToken(ht.Request().Header.Get("Authorization"))
			if parserErr != nil {
				return nil, auth.ErrAuthFail
			}

			var key = contextkey.Key("userID")
			etxs := context.WithValue(ctx, key, uinfos.ID)
			reply, err = handler(etxs, req)
		}
		return
	}
}
