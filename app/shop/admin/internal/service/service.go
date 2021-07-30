package service

import (
	v1 "casso/api/shop/v1/admin"
	"casso/app/shop/admin/internal/biz"
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewShopAdmin)

type ShopAdmin struct {
	v1.UnimplementedShopAdminServer

	log *log.Helper
}

func NewShopAdmin(uc *biz.ShopUseCase, logger log.Logger) *ShopAdmin {
	return &ShopAdmin{
		log: log.NewHelper(log.With(logger, "module", "service/shopadmin")),
	}
}

func (s *ShopAdmin) ListShopAdmin(ctx context.Context, in *v1.ListShopAdminRequest) (*v1.ListShopAdminReply, error) {
	return &v1.ListShopAdminReply{
		Code: 20000,
		Msg:  "操作成功！",
		Data: &v1.ListShopAdminReply_Data{Id: 1, Name: "casso", Age: 23},
	}, nil
}
