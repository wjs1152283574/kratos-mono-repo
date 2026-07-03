package biz

import (
	"casso/app/user/internal/model"
	"context"
)

// UserRepo 数据操作接口
type UserRepo interface {
	Create(ctx context.Context, c *model.User) (*model.User, error)
	Get(ctx context.Context, id int64) (*model.User, error)
	Update(ctx context.Context, u *model.User) (*model.User, error)
	Delete(ctx context.Context, id int64) (*model.User, error)
	List(ctx context.Context, pageNum, pageSize int64) ([]*model.User, error)
	GetUserByMobile(ctx context.Context, mobile string) (*model.User, error)
}
