/*
 * @Author: Casso
 * @Date: 2021-11-17 16:24:19
 * @LastEditors: Casso
 * @LastEditTime: 2021-11-18 19:45:44
 * @Description: file content
 * @FilePath: /kratos-mono-repo/pkg/util/token/token.go
 */
package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

// Parse token
var (
	InvalidErr     error         = errors.New("Couldn't handle this token:")
	SignKey        string        = "cassoWong"
	ExpiredTime    time.Duration = time.Minute * 60
	AddExpiredTime time.Duration = time.Minute * 60
)

// JWT 签名结构
type JWT struct {
	SigningKey []byte
}

// CustomClaims 载荷，可以加一些自己需要的信息
type CustomClaims struct {
	ID      int    `json:"id"`
	AppCode string `json:"app_code"` // 区分app用户
	jwt.StandardClaims
}

// NewJWT 新建一个jwt实例
func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSignKey()),
	}
}

// GetSignKey 获取signKey
func GetSignKey() string {
	return SignKey
}

// SetSignKey 这是SignKey
func SetSignKey(key string) string {
	SignKey = key
	return SignKey
}

// CreateToken 生成一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	// 设置自定义token过期时间
	claims.StandardClaims.ExpiresAt = time.Now().Add(ExpiredTime).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken 解析Tokne
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return nil, InvalidErr
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, InvalidErr
}

// RefreshToken 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(AddExpiredTime).Unix()
		return j.CreateToken(*claims)
	}
	return "", InvalidErr
}
