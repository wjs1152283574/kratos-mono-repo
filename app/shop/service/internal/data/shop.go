package data

import (
	"casso/app/shop/service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.ShopRepo = (*ShopRepo)(nil)

type ShopRepo struct {
	data *Data
	log  *log.Helper
}

func NewShopRepo(data *Data, logger log.Logger) biz.ShopRepo {
	return &ShopRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/shop")),
	}
}
