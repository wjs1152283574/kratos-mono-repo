package biz

import (
	"context"

	"casso/app/user/internal/model"
	user_proto "casso/gen/user/v1"
	"casso/pkg/auth"
	"casso/pkg/errors"
)

// CreateUser 创建用户
func (uc *UserUseCase) CreateUser(ctx context.Context, u *model.User) (*user_proto.CreateUserReply, error) {
	res, err := uc.repo.Create(ctx, u)
	if err != nil {
		return &user_proto.CreateUserReply{}, err
	}
	return &user_proto.CreateUserReply{
		Id:       int64(res.ID),
		Mobile:   res.Mobile,
		NickName: res.Name,
		Age:      res.Age,
	}, nil
}

// GetUser 获取用户
func (uc *UserUseCase) GetUser(ctx context.Context, id int64) (*user_proto.GetUserReply, error) {
	res, err := uc.repo.Get(ctx, id)
	if err != nil {
		return &user_proto.GetUserReply{}, err
	}
	return &user_proto.GetUserReply{
		Id:       int64(res.ID),
		Mobile:   res.Mobile,
		NickName: res.Mobile,
		Age:      res.Age,
	}, nil
}

// DeleteUser 删除用户
func (uc *UserUseCase) DeleteUser(ctx context.Context, id int64) (*user_proto.DeleteUserReply, error) {
	_, err := uc.repo.Delete(ctx, id)
	if err != nil {
		return &user_proto.DeleteUserReply{Ok: false}, nil
	}
	return &user_proto.DeleteUserReply{Ok: true}, nil
}

// UpdateUser 更新用户
func (uc *UserUseCase) UpdateUser(ctx context.Context, req *user_proto.UpdateUserRequest) (*user_proto.UpdateUserReply, error) {
	var user model.User
	user.ID = uint(req.Id)
	user.Age = req.Age
	user.Name = req.NickName
	res, err := uc.repo.Update(ctx, &user)
	if err != nil {
		return &user_proto.UpdateUserReply{}, err
	}
	return &user_proto.UpdateUserReply{
		Mobile:   res.Mobile,
		NickName: res.Name,
		Age:      res.Age,
	}, nil
}

// UserList 用户列表
func (uc *UserUseCase) UserList(ctx context.Context, pageNum, pageSize int64) (*user_proto.ListUserReply, error) {
	var res []*user_proto.ListUserReply_User
	list, err := uc.repo.List(ctx, pageNum, pageSize)
	if err != nil {
		return &user_proto.ListUserReply{}, err
	}

	for _, v := range list {
		res = append(res, &user_proto.ListUserReply_User{
			Id:       int64(v.ID),
			Mobile:   v.Mobile,
			NickName: v.Name,
		})
	}
	return &user_proto.ListUserReply{Users: res}, nil
}

// Login 登录
func (uc *UserUseCase) Login(ctx context.Context, u *user_proto.GetTokenRequest) (*user_proto.GetTokenReply, error) {
	user, err := uc.repo.GetUserByMobile(ctx, u.Mobile)
	if err != nil {
		return nil, err
	}

	if user.Pass != auth.Base64Md5(u.Pass) {
		return nil, errors.InvalidParams
	}

	t, err := auth.NewJWT().CreateToken(auth.CustomClaims{
		ID: int(user.ID),
	})
	if err != nil {
		return nil, errors.MakeTokenFaild
	}
	return &user_proto.GetTokenReply{Token: t}, nil
}
