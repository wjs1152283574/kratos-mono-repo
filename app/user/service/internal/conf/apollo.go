package conf

import (
	"github.com/go-kratos/kratos/contrib/config/apollo/v2"
	"github.com/go-kratos/kratos/v2/config"
)

// LoadApollo apollo 加载远程配置
func LoadApollo(apo *Data_Apollo) config.Config {
	c := config.New(
		config.WithSource(
			apollo.NewSource(
				apollo.WithAppID(apo.AppId),
				apollo.WithCluster(apo.Cluster),
				apollo.WithEndpoint(apo.Addr),
				apollo.WithNamespace(apo.Namespace),
				// apollo.WithEnableBackup(),
				apollo.WithSecret(apo.Screct),
			),
		),
	)
	// defer c.Close()

	// 加载远程配置
	if err := c.Load(); err != nil {
		panic(err)
	}

	return c
}
