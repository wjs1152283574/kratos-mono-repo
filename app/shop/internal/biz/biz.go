package biz

import (
	v1 "casso/gen/user/v1"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewShopUseCase)

// ShopUseCase BFF 业务用例
type ShopUseCase struct {
	log *log.Helper
	uc  v1.UserClient
}

// NewShopUseCase 创建 ShopUseCase
func NewShopUseCase(logger log.Logger, uc v1.UserClient) *ShopUseCase {
	return &ShopUseCase{
		log: log.NewHelper(log.With(logger, "module", "usecase/shop")),
		uc:  uc,
	}
}
