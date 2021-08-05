package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	ID   int
	Name string
	Age  int64
}

type UserForToken struct {
	Mobile string
	Pass   string
	ID     int
}

type UserRepo interface {
	CreateUser(ctx context.Context, c *User) (*User, error)
	GetUser(ctx context.Context, id int64) (*User, error)
	UpdateUser(ctx context.Context, c *User) (*User, error)
	DeleteUser(ctx context.Context, c *User) (*User, error)
	ListUser(ctx context.Context, pageNum, pageSize int64) ([]*User, error)
	GetToken(ctx context.Context, u *UserForToken) (token string, err error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/user"))}
}

func (uc *UserUseCase) Create(ctx context.Context, u *User) (*User, error) {
	return uc.repo.CreateUser(ctx, u)
}

func (uc *UserUseCase) Get(ctx context.Context, id int64) (*User, error) {
	return uc.repo.GetUser(ctx, id)
}

func (uc *UserUseCase) Delete(ctx context.Context, u *User) (*User, error) {
	return uc.repo.UpdateUser(ctx, u)
}

func (uc *UserUseCase) Update(ctx context.Context, u *User) (*User, error) {
	return uc.repo.UpdateUser(ctx, u)
}

func (uc *UserUseCase) List(ctx context.Context, pageNum, pageSize int64) ([]*User, error) {
	return uc.repo.ListUser(ctx, pageNum, pageSize)
}

func (uc *UserUseCase) Login(ctx context.Context, u *UserForToken) (token string, err error) {
	return uc.repo.GetToken(ctx, u)
}
