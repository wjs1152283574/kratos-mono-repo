package data

import (
	"casso/app/shop/service/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewShopRepo)

// Data .
type Data struct {
	log *log.Helper
}

// NewData .
func NewData(conf *conf.Data, logger log.Logger) (*Data, error) {
	l := log.NewHelper(log.With(logger, "module", "shop-admin/data"))
	return &Data{log: l}, nil
}
