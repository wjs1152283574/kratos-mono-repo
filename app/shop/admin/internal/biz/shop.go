package biz

import (
	"github.com/go-kratos/kratos/v2/log"
)

type Shop struct {
}

type ShopRepo interface {
}

type ShopUseCase struct {
	repo ShopRepo
	log  *log.Helper
}

func NewShopUseCase(repo ShopRepo, logger log.Logger) *ShopUseCase {
	return &ShopUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/shop"))}
}
