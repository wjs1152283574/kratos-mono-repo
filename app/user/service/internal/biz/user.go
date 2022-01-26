package biz

import (
	user_proto "casso/api/user/service/v1"
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Mobile string `gorm:"unique"`
	Pass   string
	Name   string
	Age    int64
}

type UserReply struct {
	Name, Mobile string
	Age, ID      int64
}

type UserForToken struct {
	Mobile string
	Pass   string
	ID     int
}

// 在此实现对data层的数据操作
type UserRepo interface {
	// 新建用户
	Create(ctx context.Context, c *User) (*User, error)
	// 获取用户信息
	Get(ctx context.Context, id int64) (*User, error)
	// 编辑用户信息
	Update(ctx context.Context, c *User) (*User, error)
	// 删除
	Delete(ctx context.Context, id int64) (*User, error)

	ListUser(ctx context.Context, pageNum, pageSize int64) ([]*User, error)
	GetToken(ctx context.Context, u *UserForToken) (token string, err error)
	GetUserByName(ctx context.Context, name string) (*User, error)
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

func (uc *UserUseCase) List(ctx context.Context, pageNum, pageSize int64) ([]*User, error) {
	return uc.repo.ListUser(ctx, pageNum, pageSize)
}

func (uc *UserUseCase) Login(ctx context.Context, u *UserForToken) (token string, err error) {
	return uc.repo.GetToken(ctx, u)
}
