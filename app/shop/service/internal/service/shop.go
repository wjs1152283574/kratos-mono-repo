package service

import (
	pb "casso/api/shop/service/v1"
	"casso/app/shop/service/internal/biz"
	"casso/pkg/errors/normal"
	"casso/pkg/util/contextkey"
	"context"

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
	res, err := s.sc.Register(ctx, req)
	if err != nil {
		return &pb.RegisterResponse{}, pb.ErrorDuplicateEntry(err.Error())
	}
	return res, nil
}

func (s *ShopService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	res, err := s.sc.Login(ctx, req)
	if err != nil {
		e := errors.FromError(err)

		res.Code = e.Code
		res.Msg = e.Message
		return res, nil //errors.New(int(e.Code), e.Reason, e.Message)
	}
	return res, nil
}

func (s *ShopService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	var key = contextkey.Key("userID")
	id := ctx.Value(key)
	if id == nil {
		return nil, normal.InvalidParams
	}

	res, err := s.sc.GetUser(ctx, int64(id.(int)))
	if err != nil {
		e := errors.FromError(err)
		return nil, errors.New(int(e.Code), e.Reason, e.Message)
	}
	return res, nil
}
