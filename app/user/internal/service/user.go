package service

import (
	"context"

	"casso/app/user/internal/model"
	pb "casso/gen/user/v1"
)

// CreateUser 创建用户
func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	if req.NickName == "" || req.Mobile == "" {
		return &pb.CreateUserReply{}, pb.ErrorUserInvalidParams("invalid params")
	}
	return s.uc.CreateUser(ctx, &model.User{Mobile: req.Mobile, Pass: req.Pass, Name: req.NickName, Age: req.Age})
}

// GetUser 获取用户
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	if req.Id == 0 {
		return &pb.GetUserReply{}, pb.ErrorUserInvalidParams("invalid params")
	}
	return s.uc.GetUser(ctx, req.Id)
}

// UpdateUser 更新用户
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	if req.Id == 0 {
		return &pb.UpdateUserReply{}, pb.ErrorUserInvalidParams("invalid params")
	}
	return s.uc.UpdateUser(ctx, req)
}

// DeleteUser 删除用户
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	if req.Id == 0 {
		return &pb.DeleteUserReply{}, pb.ErrorUserInvalidParams("invalid params")
	}
	return s.uc.DeleteUser(ctx, req.Id)
}

// ListUser 用户列表
func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	if req.Page == 0 || req.Limit == 0 {
		return &pb.ListUserReply{}, pb.ErrorUserInvalidParams("invalid params")
	}
	return s.uc.UserList(ctx, req.Page, req.Limit)
}

// GetToken 登录获取 Token
func (s *UserService) GetToken(ctx context.Context, req *pb.GetTokenRequest) (*pb.GetTokenReply, error) {
	if req.Mobile == "" || req.Pass == "" {
		return &pb.GetTokenReply{}, pb.ErrorUserInvalidParams("invalid params")
	}
	return s.uc.Login(ctx, req)
}

// CreateTestUser 创建测试用户
func (s *UserService) CreateTestUser(ctx context.Context, req *pb.CreateTestUserRequest) (*pb.CreateTestUserReply, error) {
	return &pb.CreateTestUserReply{Result: true}, nil
}
