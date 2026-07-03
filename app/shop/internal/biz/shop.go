package biz

import (
	"context"

	pb "casso/gen/shop/v1"
	v1 "casso/gen/user/v1"
)

// Register 注册
func (s *ShopUseCase) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterReply, error) {
	res, err := s.uc.CreateUser(ctx, &v1.CreateUserRequest{
		NickName: req.NickName,
	})
	if err != nil || !v1.IsUserRecordNotFound(err) {
		return &pb.RegisterReply{}, err
	}

	return &pb.RegisterReply{
		Id:       res.Id,
		NickName: res.NickName,
	}, nil
}

// Login 登录
func (s *ShopUseCase) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	res, err := s.uc.GetToken(ctx, &v1.GetTokenRequest{
		Mobile: req.Mobile,
		Pass:   req.Pass,
	})
	if err != nil {
		return &pb.LoginReply{}, err
	}

	return &pb.LoginReply{
		Token: res.Token,
	}, nil
}

// GetUser 获取用户信息
func (s *ShopUseCase) GetUser(ctx context.Context, id int64) (*pb.GetUserReply, error) {
	res, err := s.uc.GetUser(ctx, &v1.GetUserRequest{
		Id: id,
	})
	if err != nil {
		return &pb.GetUserReply{}, err
	}

	return &pb.GetUserReply{
		Name: res.NickName,
	}, nil
}
