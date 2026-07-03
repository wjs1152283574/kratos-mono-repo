package data

import (
	"time"

	"casso/app/user/internal/conf"
	"casso/app/user/internal/model"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewDB 创建 MySQL 连接
func NewDB(conf *conf.Data, loggers log.Logger) *gorm.DB {
	log := log.NewHelper(log.With(loggers, "module", "user-service/data/gorm"))
	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		db.Statement.ReflectValue.Close()
		log.Fatalf("failed making connection-pool: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatal(err)
	}
	return db
}
