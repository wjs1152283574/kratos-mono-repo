package data

import (
	"casso/app/user/service/internal/biz"
	"context"

	"github.com/go-kratos/beer-shop/pkg/util/pagination"
	"gorm.io/gorm"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.UserRepo = (*UserRepo)(nil)

type UserRepo struct {
	data *Data
	log  *log.Helper
}

type User struct {
	gorm.Model
	Name string
	Age  int64
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &UserRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/user")),
	}
}

func (r *UserRepo) CreateUser(ctx context.Context, b *biz.User) (*biz.User, error) {
	o := User{Name: b.Name, Age: b.Age}
	result := r.data.db.WithContext(ctx).Create(o)
	return &biz.User{
		Name: o.Name,
	}, result.Error
}

func (r *UserRepo) GetUser(ctx context.Context, id int64) (*biz.User, error) {
	o := User{}
	result := r.data.db.WithContext(ctx).First(&o, id)
	return &biz.User{
		Name: o.Name,
	}, result.Error
}

func (r *UserRepo) UpdateUser(ctx context.Context, b *biz.User) (*biz.User, error) {
	o := User{}
	result := r.data.db.WithContext(ctx).First(&o, b.Name)
	if result.Error != nil {
		return nil, result.Error
	}
	o.Name = b.Name
	result = r.data.db.WithContext(ctx).Save(&o)
	if result.Error != nil {
		return nil, result.Error
	}
	return &biz.User{
		Name: o.Name,
	}, nil
}

func (r *UserRepo) DeleteUser(ctx context.Context, b *biz.User) (*biz.User, error) {
	o := User{Name: b.Name}
	result := r.data.db.WithContext(ctx).Delete(&o, b.Name)
	if result.Error != nil {
		return nil, result.Error
	}
	o.Name = b.Name
	result = r.data.db.WithContext(ctx).Save(&o)
	if result.Error != nil {
		return nil, result.Error
	}
	return &biz.User{
		Name: o.Name,
	}, nil
}

func (r *UserRepo) ListUser(ctx context.Context, pageNum, pageSize int64) ([]*biz.User, error) {
	var os []User
	result := r.data.db.WithContext(ctx).
		Limit(int(pageSize)).
		Offset(int(pagination.GetPageOffset(pageNum, pageSize))).
		Find(&os)
	if result.Error != nil {
		return nil, result.Error
	}
	rv := make([]*biz.User, 0)
	for _, o := range os {
		rv = append(rv, &biz.User{
			Name: o.Name,
		})
	}
	return rv, nil
}
