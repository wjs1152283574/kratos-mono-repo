/*
 * @Author: Casso
 * @Date: 2021-11-17 16:24:19
 * @LastEditors: Casso
 * @LastEditTime: 2021-11-22 10:26:44
 * @Description: file content
 * @FilePath: /kratos-mono-repo/app/shop/service/internal/server/server.go
 */
package server

import (
	"github.com/google/wire"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer, NewUserServiceClient, NewDiscovery)
