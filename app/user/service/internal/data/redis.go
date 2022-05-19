package data

import (
	"casso/app/user/service/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
)

func NewRd(conf *conf.Data, logger log.Logger) *redis.Client {
	log.NewHelper(log.With(logger, "module", "user-service/data/redis"))
	opts := redis.Options{
		Addr:         conf.Redis.Addr,
		Username:     conf.Redis.Auth,
		Password:     conf.Redis.Password,
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
		PoolSize:     int(conf.Redis.Pool),
		DB:           0,
	}
	return redis.NewClient(&opts)
}
