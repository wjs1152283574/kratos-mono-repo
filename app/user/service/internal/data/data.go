/*
 * @Author: Casso
 * @Date: 2021-11-17 16:24:19
 * @LastEditors: Casso
 * @LastEditTime: 2021-11-26 12:12:33
 * @Description: file content
 * @FilePath: /kratos-mono-repo/app/user/service/internal/data/data.go
 */
package data

import (
	"casso/app/user/service/internal/conf"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewRd, NewUserRepo)

// Data .
type Data struct {
	rd  *redis.Client
	db  *gorm.DB
	log *log.Helper
}

func NewRd(conf *conf.Data, logger log.Logger) *redis.Client {
	log.NewHelper(log.With(logger, "module", "user-service/data/redis"))
	opts := redis.Options{
		// Addr:         conf.Addr,
		// WriteTimeout: conf.WriteTimeout.AsDuration(),
		// ReadTimeout:  conf.GetReadTimeout().AsDuration(),
	}
	return redis.NewClient(&opts)
}

func NewDB(conf *conf.Data, logger log.Logger) *gorm.DB {
	log := log.NewHelper(log.With(logger, "module", "user-service/data/gorm"))
	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}

	sqlDB, err := db.DB() // 维护链接池
	if err != nil {
		db.Statement.ReflectValue.Close()
		log.Fatalf("failed making connection-pool: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)           // 空闲最大数量
	sqlDB.SetMaxOpenConns(100)          // 最大链接数
	sqlDB.SetConnMaxLifetime(time.Hour) // 最大可复用时间

	if err := db.AutoMigrate(&User{}); err != nil {
		log.Fatal(err)
	}
	return db
}

// NewData .
func NewData(db *gorm.DB, rd *redis.Client, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "user-service/data"))

	d := &Data{
		rd:  rd,
		db:  db,
		log: log,
	}
	return d, func() {

	}, nil
}
