package middleware

import (
	"context"

	"casso/pkg/contextx"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	ht "github.com/go-kratos/kratos/v2/transport/http"
)

// AuthFromQuery 从 URL 查询参数提取 UserID 注入 Context
// 适用于 BFF 层：API Gateway 认证后将 userid 添加到查询参数转发到 BFF
func AuthFromQuery(key string) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				if httpTr, ok := tr.(*ht.Transport); ok {
					vals := httpTr.Request().URL.Query()[key]
					if len(vals) > 0 {
						ctx = context.WithValue(ctx, contextx.NewKey(), vals[0])
					}
				}
			}
			return handler(ctx, req)
		}
	}
}
