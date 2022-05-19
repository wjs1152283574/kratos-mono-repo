package biz

import (
	"casso/app/user/service/internal/model"
	"context"
)

// 在此实现对data层的数据操作
type UserRepo interface {
	// 新建用户
	Create(ctx context.Context, c *model.User) (*model.User, error)
	// 获取用户信息
	Get(ctx context.Context, id int64) (*model.User, error)
	// 编辑用户信息
	Update(ctx context.Context, u *model.User) (*model.User, error)
	// 删除
	Delete(ctx context.Context, id int64) (*model.User, error)
	// 列表
	List(ctx context.Context, pageNum, pageSize int64) ([]*model.User, error)
	// 通过电话获取用户
	GetUserByMobile(ctx context.Context, mobile string) (user *model.User, err error)
}
