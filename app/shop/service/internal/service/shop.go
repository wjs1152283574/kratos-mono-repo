package service

import (
	pb "casso/api/shop/service/v1"
	v1 "casso/api/user/service/v1"
	"casso/app/shop/service/internal/biz"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type ShopService struct {
	pb.UnimplementedShopServer

	sc  *biz.ShopUseCase
	log *log.Helper
}

func NewShopService(sc *biz.ShopUseCase, logger log.Logger) *ShopService {
	return &ShopService{
		sc:  sc,
		log: log.NewHelper(log.With(logger, "module", "service/shop"))}
}

func (s *ShopService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	token, err := s.sc.Login(ctx, &v1.GetTokenRequest{
		Mobile: "17620439807",
		Pass:   "casso",
	})
	if err != nil {
		return &pb.LoginReply{
			Token: "错误、",
		}, v1.ErrorRecordNotFound("用户不存在！")
	}
	return &pb.LoginReply{
		Token: token.Token,
	}, nil
}
