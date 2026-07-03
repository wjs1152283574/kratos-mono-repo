package middleware

import (
	ht "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
)

// CORS 返回跨域中间件配置
// TODO: 引入网关后在网关处理跨域时，需要删除此中间件
func CORS() ht.FilterFunc {
	return handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{
			"DNT", "X-Mx-ReqToken", "Keep-Alive", "User-Agent",
			"X-Requested-With", "If-Modified-Since", "Cache-Control",
			"Content-Type", "Authorization", "udid", "appkey",
			"version", "authenticated", "cookie", "token",
		}),
		handlers.ExposedHeaders([]string{
			"DNT", "X-Mx-ReqToken", "Keep-Alive", "User-Agent",
			"X-Requested-With", "If-Modified-Since", "Cache-Control",
			"Content-Type", "Authorization", "udid", "appkey",
			"version", "authenticated", "cookie", "token",
		}),
		handlers.OptionStatusCode(204),
	)
}
