package data

import (
	"casso/app/user/service/internal/conf"
	"casso/app/user/service/internal/model"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDB(conf *conf.Data, loggers log.Logger) *gorm.DB {
	log := log.NewHelper(log.With(loggers, "module", "user-service/data/gorm"))
	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 以Infos 形式输出SQL语句
	})

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
