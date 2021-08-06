package service

import (
	"context"
	"fmt"

	pb "casso/api/user/service/v1"
	"casso/app/user/service/internal/biz"
)

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	res, err := s.uc.Create(ctx, &biz.User{Mobile: req.Mobile, Pass: req.Pass, Name: req.NickName, Age: req.Age})
	return &pb.CreateUserReply{
		Id:       res.ID,
		Mobile:   res.Mobile,
		NickName: res.Name,
		Age:      res.Age,
	}, err
}
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	return &pb.UpdateUserReply{}, nil
}
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	return &pb.DeleteUserReply{}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	return &pb.GetUserReply{}, nil
}
func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	return &pb.ListUserReply{}, nil
}

func (s *UserService) GetToken(ctx context.Context, req *pb.GetTokenRequest) (*pb.GetTokenReply, error) {
	token, err := s.uc.Login(ctx, &biz.UserForToken{Mobile: req.Mobile, Pass: req.Pass, ID: 2})
	if err != nil {
		fmt.Printf("USER_SERVICE ==> err_type:=%#T\n", err)
		return &pb.GetTokenReply{}, err
	}
	return &pb.GetTokenReply{Token: token}, nil
}
