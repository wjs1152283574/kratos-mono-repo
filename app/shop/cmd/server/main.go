package main

import (
	"flag"
	"os"
	"time"

	"casso/app/shop/internal/conf"
	nacosreg "casso/pkg/registry"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/http"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"google.golang.org/protobuf/types/known/durationpb"
	"gopkg.in/yaml.v2"
)

var (
	Name     = "casso.shop.service"
	Version  = "v1"
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

// 普通配置结构体（用于 YAML 反序列化）
type AppConfig struct {
	Trace struct {
		Endpoint string `yaml:"endpoint"`
	} `yaml:"trace"`
	Server struct {
		HTTP struct {
			Addr    string        `yaml:"addr"`
			Timeout time.Duration `yaml:"timeout"`
		} `yaml:"http"`
	} `yaml:"server"`
	Discovery struct {
		Nacos struct {
			Address string `yaml:"address"`
			Port    int    `yaml:"port"`
		} `yaml:"nacos"`
	} `yaml:"discovery"`
}

func newApp(logger log.Logger, hs *http.Server, rr registry.Registrar) *kratos.App {
	return kratos.New(
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(hs),
		kratos.Registrar(rr),
	)
}

func main() {
	flag.Parse()
	logger := log.With(log.NewStdLogger(os.Stdout),
		"service.name", Name,
		"service.version", Version,
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
	)

	// 读取 YAML 配置
	data, err := os.ReadFile(flagconf + "/config.yaml")
	if err != nil {
		panic(err)
	}
	var cfg AppConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		panic(err)
	}

	// 构造 Nacos 配置
	nacosConf := nacosreg.NacosConfig{
		Address: cfg.Discovery.Nacos.Address,
		Port:    uint64(cfg.Discovery.Nacos.Port),
	}

	// 构造 protobuf conf 对象（供 wire 注入）
	serverConf := &conf.Server{
		Http: &conf.Server_HTTP{
			Addr:    cfg.Server.HTTP.Addr,
			Timeout: durationpb.New(cfg.Server.HTTP.Timeout),
		},
	}

	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(cfg.Trace.Endpoint)))
	if err != nil {
		panic(err)
	}
	tp := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exp),
		tracesdk.WithResource(resource.NewSchemaless(
			semconv.ServiceNameKey.String(Name),
			attribute.String("service", "shop"),
		)),
	)

	app, cleanup, err := initApp(serverConf, logger, tp, nacosConf)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	if err := app.Run(); err != nil {
		panic(err)
	}
}
