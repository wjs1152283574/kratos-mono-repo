package service

import (
	pb "casso/api/shop/service/v1"
	v1 "casso/api/user/service/v1"
	"casso/app/shop/service/internal/biz"
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/errors"
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

func (s *ShopService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	u, err := s.sc.Register(ctx, &v1.CreateUserRequest{Mobile: req.Mobile, Pass: req.Pass, Age: req.Age, NickName: req.NickName})
	if err != nil {
		return &pb.RegisterResponse{}, pb.ErrorDuplicateEntry(err.Error())
	}
	return &pb.RegisterResponse{
		Id:       u.Id,
		Mobile:   u.Mobile,
		NickName: u.NickName,
		Age:      u.Age,
	}, nil
}

func (s *ShopService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	token, err := s.sc.Login(ctx, &v1.GetTokenRequest{
		Mobile: req.Mobile,
		Pass:   req.Pass,
	})
	if err != nil {
		fmt.Printf("SHOP_SERVICE ==> err_type:=%#T\n,%#v\n", err, err)
		e := errors.FromError(err)
		return &pb.LoginReply{
			Token: "错误、",
		}, errors.New(int(e.Code), e.Reason, e.Message)
	}
	return &pb.LoginReply{
		Token: token.Token,
	}, nil
}
