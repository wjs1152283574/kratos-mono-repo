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
	user := &User{Name: b.Name, Age: b.Age, Mobile: b.Mobile, Pass: passmd5.Base64Md5(b.Pass)}
	result := r.data.db.WithContext(ctx).Create(user).First(user)
	return &biz.UserReply{
		Name:   user.Name,
		Mobile: user.Mobile,
		Age:    user.Age,
		ID:     int64(user.ID),
	}, result.Error
}

func (r *UserRepo) GetUser(ctx context.Context, id int64) (*biz.User, error) {
	user := User{}
	result := r.data.db.WithContext(ctx).First(&user, id)
	return &biz.User{
		Name: user.Name,
	}, result.Error
}

func (r *UserRepo) UpdateUser(ctx context.Context, b *biz.User) (*biz.User, error) {
	user := User{}
	result := r.data.db.WithContext(ctx).First(&user, b.Name)
	if result.Error != nil {
		return nil, result.Error
	}
	user.Name = b.Name
	result = r.data.db.WithContext(ctx).Save(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &biz.User{
		Name: user.Name,
	}, nil
}

func (r *UserRepo) DeleteUser(ctx context.Context, b *biz.User) (*biz.User, error) {
	user := User{Name: b.Name}
	result := r.data.db.WithContext(ctx).Delete(&user, b.Name)
	if result.Error != nil {
		return nil, result.Error
	}
	user.Name = b.Name
	result = r.data.db.WithContext(ctx).Save(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &biz.User{
		Name: user.Name,
	}, nil
}

func (r *UserRepo) ListUser(ctx context.Context, pageNum, pageSize int64) ([]*biz.User, error) {
	var userList []User
	result := r.data.db.WithContext(ctx).
		Limit(int(pageSize)).
		Offset(int(pagination.GetPageOffset(pageNum, pageSize))).
		Find(&userList)
	if result.Error != nil {
		return nil, result.Error
	}
	rv := make([]*biz.User, 0)
	for _, user := range userList {
		rv = append(rv, &biz.User{
			Name: user.Name,
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
