package biz

import (
	user_proto "casso/api/user/service/v1"
	"casso/app/user/service/internal/pkg/utill/passmd5"
	"casso/pkg/errors/normal"
	"casso/pkg/util/token"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// 在此实现对data层的数据操作
type UserRepo interface {
	// 新建用户
	Create(ctx context.Context, c *User) (*User, error)
	// 获取用户信息
	Get(ctx context.Context, id int64) (*User, error)
	// 编辑用户信息
	Update(ctx context.Context, u *User) (*User, error)
	// 删除
	Delete(ctx context.Context, id int64) (*User, error)
	// 列表
	List(ctx context.Context, pageNum, pageSize int64) ([]*User, error)
	// 通过电话获取用户
	GetUserByMobile(ctx context.Context, mobile string) (user *User, err error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/user"))}
}

// ********* 以下实现业务组装，实现service需求 ***********
func (uc *UserUseCase) CreateUser(ctx context.Context, u *User) (*user_proto.CreateUserReply, error) {
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

func (uc *UserUseCase) DeleteUser(ctx context.Context, id int64) (*user_proto.DeleteUserReply, error) {
	_, err := uc.repo.Delete(ctx, id)
	if err != nil {
		return &user_proto.DeleteUserReply{
			Ok: false,
		}, nil
	}
	return &user_proto.DeleteUserReply{
		Ok: true,
	}, nil
}

func (uc *UserUseCase) UpdateUser(ctx context.Context, req *user_proto.UpdateUserRequest) (*user_proto.UpdateUserReply, error) {
	var user User
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

func (uc *UserUseCase) UserList(ctx context.Context, pageNum, pageSize int64) (*user_proto.ListUserReply, error) {
	var res []*user_proto.ListUserReply_User
	list, err := uc.repo.List(ctx, pageNum, pageSize)
	if err != nil {
		return &user_proto.ListUserReply{}, err
	}

	// 拼接数据
	for _, v := range list {
		res = append(res, &user_proto.ListUserReply_User{
			Id:       int64(v.ID),
			Mobile:   v.Mobile,
			NickName: v.Name,
		})
	}
	return &user_proto.ListUserReply{Users: res}, nil
}

func (uc *UserUseCase) Login(ctx context.Context, u *user_proto.GetTokenRequest) (res *user_proto.GetTokenReply, err error) {
	user, err := uc.repo.GetUserByMobile(ctx, u.Mobile)
	if err != nil {
		return res, err
	}

	if user.Pass != passmd5.Base64Md5(u.Pass) {
		return res, normal.InvalidParams
	}

	t, err := token.NewJWT().CreateToken(token.CustomClaims{
		ID: int(user.ID),
	})

	if err != nil {
		return res, normal.MakeTokenFaild
	}
	return &user_proto.GetTokenReply{Token: t}, nil
}
