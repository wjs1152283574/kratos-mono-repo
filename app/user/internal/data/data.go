package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewRd, NewUserRepo)

// Data 数据层
type Data struct {
	rd  *redis.Client
	db  *gorm.DB
	log *log.Helper
}

// NewData 创建 Data 实例
func NewData(db *gorm.DB, rd *redis.Client, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "user-service/data"))

	d := &Data{
		rd:  rd,
		db:  db,
		log: log,
	}

	return d, func() {
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
