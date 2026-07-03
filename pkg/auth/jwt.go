package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	ErrInvalidToken = errors.New("couldn't handle this token")
)

// JWT 签名配置，可通过 SetSignKey 在启动时覆盖
var (
	SignKey        = "cassoWong"
	ExpiredTime    = time.Minute * 60
	AddExpiredTime = time.Minute * 60
)

// SetSignKey 设置签名密钥，应在服务启动时调用
func SetSignKey(key string) {
	SignKey = key
}

// CustomClaims JWT 载荷
type CustomClaims struct {
	ID      int    `json:"id"`
	AppCode string `json:"app_code"`
	jwt.StandardClaims
}

// JWT 签名结构
type JWT struct {
	SigningKey []byte
}

// NewJWT 新建一个 JWT 实例
func NewJWT() *JWT {
	return &JWT{
		SigningKey: []byte(SignKey),
	}
}

// CreateToken 生成一个 token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	claims.StandardClaims.ExpiresAt = time.Now().Add(ExpiredTime).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken 解析 token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return nil, ErrInvalidToken
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrInvalidToken
}

// RefreshToken 更新 token
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
	return "", ErrInvalidToken
}
