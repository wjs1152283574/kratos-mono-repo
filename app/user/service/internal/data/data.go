/*
 * @Author: Casso
 * @Date: 2021-11-17 16:24:19
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2022-05-19 10:32:08
 * @Description: file content
 * @FilePath: /kratos-mono-repo/app/user/service/internal/data/data.go
 */
package data

import (
	"casso/app/user/service/internal/conf"
	"casso/app/user/service/internal/model"
	"fmt"
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

	if err := db.AutoMigrate(&model.User{}); err != nil {
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

	// 监听配置文件并处理
	go func() {
		for v := range conf.ConfCh {
			fmt.Println("初始化自定义配置文件：", v.CassoConf)
		}
	}()

	// 启动定时任务
	go InitTimer(*d)

	return d, func() {

	}, nil
}
