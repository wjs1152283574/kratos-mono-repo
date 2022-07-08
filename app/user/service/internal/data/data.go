/*
 * @Author: Casso
 * @Date: 2021-11-17 16:24:19
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2022-05-19 10:33:09
 * @Description: file content
 * @FilePath: /kratos-mono-repo/app/user/service/internal/data/data.go
 */
package data

import (
	"casso/app/user/service/internal/conf"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
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
			fmt.Println("初始化自定义配置文件：", v.CassoConf[0].ID, v.CassoConf[0].Addr)
		}
	}()

	// 启动定时任务
	go InitTimer(*d)

	return d, func() {
		// 程序退出，释放资源
		sqldb, err := d.db.DB()
		if err != nil {
			log.Errorf("sqldb resource got fail: %v", err)
		}
		if err := sqldb.Close(); err != nil {
			log.Errorf("sqldb closing resource got fail: %v", err)
		}

		if err := d.rd.Close(); err != nil {
			log.Errorf("redis closing resource got fail: %v", err)
		}

		log.Info("resource close successed !")
	}, nil
}
