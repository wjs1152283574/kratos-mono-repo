package data

import (
	"casso/app/user/service/internal/biz"
	"casso/app/user/service/internal/pkg/utill/passmd5"
	"casso/pkg/errors/normal"
	"casso/pkg/util/pagination"
	"casso/pkg/util/token"
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

var _ biz.UserRepo = (*UserRepo)(nil)

type UserRepo struct {
	data *Data
	log  *log.Helper
}

type User struct {
	gorm.Model
	Mobile string `gorm:"unique"`
	Pass   string
	Name   string
	Age    int64
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &UserRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/user")),
	}
}

func (r *UserRepo) CreateUser(ctx context.Context, b *biz.User) (*biz.UserReply, error) {
	o := &User{Name: b.Name, Age: b.Age, Mobile: b.Mobile, Pass: passmd5.Base64Md5(b.Pass)}
	result := r.data.db.WithContext(ctx).Create(o).First(o)
	return &biz.UserReply{
		Name:   o.Name,
		Mobile: o.Mobile,
		Age:    o.Age,
		ID:     int64(o.ID),
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

// GetToken check user exit and return token
func (r *UserRepo) GetToken(ctx context.Context, u *biz.UserForToken) (string, error) {
	var user User
	result := r.data.db.WithContext(ctx).Where("mobile = ?", u.Mobile).First(&user)
	if result.Error != nil {
		return "", normal.RecordNotFound
	}
	if user.Pass != passmd5.Base64Md5(u.Pass) {
		return "", normal.InvalidParams
	}
	t, err := token.NewJWT().CreateToken(token.CustomClaims{
		Mobile: u.Mobile,
		ID:     int(user.ID),
	})
	if err != nil {
		return "", normal.MakeTokenFaild
	}
	return t, nil
}
