package biz

import (
	v1 "casso/api/user/service/v1"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewShopUseCase)

type ShopUseCase struct {
	repo ShopRepo
	log  *log.Helper

	uc v1.UserClient
}

func NewShopUseCase(repo ShopRepo, logger log.Logger, uc v1.UserClient) *ShopUseCase {
	return &ShopUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/shop")), uc: uc}
}
