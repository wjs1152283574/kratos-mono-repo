package biz

import (
	v1 "casso/api/user/service/v1"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Shop struct {
}

type ShopRepo interface {
}

type ShopUseCase struct {
	repo ShopRepo
	log  *log.Helper

	uc v1.UserClient
}

func NewShopUseCase(repo ShopRepo, logger log.Logger, uc v1.UserClient) *ShopUseCase {
	return &ShopUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/shop")), uc: uc}
}

func (s *ShopUseCase) Register(ctx context.Context, req *v1.CreateUserRequest) (*v1.CreateUserReply, error) {
	return s.uc.CreateUser(ctx, req)
}

func (s *ShopUseCase) Login(ctx context.Context, req *v1.GetTokenRequest) (*v1.GetTokenReply, error) {
	return s.uc.GetToken(ctx, req)
}
