package biz

import (
	pb "casso/api/shop/service/v1"
	v1 "casso/api/user/service/v1"
	"context"
)

func (s *ShopUseCase) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterReply, error) {
	// 业务组装
	res, err := s.uc.CreateUser(ctx, &v1.CreateUserRequest{
		NickName: req.NickName,
	})
	if err != nil || !v1.IsRecordNotFound(err) {
		return &pb.RegisterReply{}, err
	}

	return &pb.RegisterReply{
		Id:       res.Id,
		NickName: res.NickName,
	}, nil
}

func (s *ShopUseCase) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	// 业务组装
	res, err := s.uc.GetToken(ctx, &v1.GetTokenRequest{
		Mobile: req.Mobile,
		Pass:   req.Pass,
	})
	if err != nil {
		return &pb.LoginReply{}, err
	}

	return &pb.LoginReply{
		Code: 20000,
		Data: &pb.LoginReply_Data{
			Token: res.Token,
		},
		Msg: "",
	}, nil
}

func (s *ShopUseCase) GetUser(ctx context.Context, id int64) (*pb.GetUserReply, error) {
	// 业务组装
	res, err := s.uc.GetUser(ctx, &v1.GetUserRequest{
		Id: id,
	})
	if err != nil {
		return &pb.GetUserReply{}, err
	}

	return &pb.GetUserReply{
		Code: 20000,
		Data: &pb.GetUserReply_Data{
			Name: res.NickName,
		},
	}, nil
}
