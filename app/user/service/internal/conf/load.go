package conf

import (
	"fmt"

	"github.com/go-kratos/kratos/v2/config"
)

// 监听配置更新&本地缓存｜｜ 持久化

// 此处可以使用proto定义的结构体，这里简易示范
type Items struct {
	ID   int64  `json:"id"`
	Addr string `json:"addr"`
}

type Users struct {
	CassoConf []Items `json:"casso_conf"`
}

// keys 这个key需要对应json文件里面的最外层的键，即监听这个键，而不是这个json文件名
var keys = []string{
	"casso_conf",
}

var UConf Users
var ConfCh = make(chan *Users, 2)

// LoadConf 项目启动&配置文件改变重新加载配置并将数据送入通道，在data层进行持久化处理；不过只是需要本地缓存可以忽略通道
func LoadConf(conf config.Config) (err error) {
	for _, key := range keys {
		err = conf.Watch(key, func(key string, value config.Value) {
			switch key {
			case "temp":
				err = value.Scan(&UConf) // conf.Scan(&UConf)
				ConfCh <- &UConf         // 每次改定会继续送往通道
			default:
				fmt.Println("load empty configs")
			}
		})
		if err != nil {
			panic(err)
		}
	}

	ConfCh <- &UConf // 项目启动送往通道

	return conf.Scan(&UConf)
}
