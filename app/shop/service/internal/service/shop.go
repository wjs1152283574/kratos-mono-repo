package service

import (
	pb "casso/api/shop/service/v1"
	v1 "casso/api/user/service/v1"
	"casso/app/shop/service/internal/biz"
	"casso/pkg/util/contextkey"
	"casso/pkg/util/errreason"
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
		e := errors.FromError(err)
		return nil, errors.New(int(e.Code), e.Reason, e.Message)
	}
	return &pb.LoginReply{
		Code: 200,
		Data: &pb.LoginReply_Data{
			Token: token.Token,
		},
	}, nil
}

func (s *ShopService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	var key = contextkey.Key("userID")
	id := ctx.Value(key)
	if id == nil {
		return nil, pb.ErrorContentMissing(errreason.INVALID_PARAMS)
	}

	res, err := s.sc.GetUser(ctx, &v1.GetUserRequest{
		Id: int64(id.(int)),
	})
	if err != nil {
		e := errors.FromError(err)
		return nil, errors.New(int(e.Code), e.Reason, e.Message)
	}
	return &pb.GetUserReply{
		Code: 200,
		Data: &pb.GetUserReply_Data{
			Name: res.NickName,
		},
	}, nil
}
