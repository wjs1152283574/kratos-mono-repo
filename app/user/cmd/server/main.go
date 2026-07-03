package main

import (
	"flag"
	"os"
	"time"

	"casso/app/user/internal/conf"
	nacosreg "casso/pkg/registry"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"google.golang.org/protobuf/types/known/durationpb"
	"gopkg.in/yaml.v2"
)

var (
	Name     = "casso.user.service"
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
		GRPC struct {
			Addr    string        `yaml:"addr"`
			Timeout time.Duration `yaml:"timeout"`
		} `yaml:"grpc"`
	} `yaml:"server"`
	Data struct {
		Database struct {
			Driver string `yaml:"driver"`
			Source string `yaml:"source"`
		} `yaml:"database"`
		Redis struct {
			Addr         string        `yaml:"addr"`
			ReadTimeout  time.Duration `yaml:"read_timeout"`
			WriteTimeout time.Duration `yaml:"write_timeout"`
			Password     string        `yaml:"password"`
		} `yaml:"redis"`
	} `yaml:"data"`
	Registry struct {
		Nacos struct {
			Address string `yaml:"address"`
			Port    int    `yaml:"port"`
		} `yaml:"nacos"`
	} `yaml:"registry"`
}

func newApp(logger log.Logger, gs *grpc.Server, rr registry.Registrar) *kratos.App {
	return kratos.New(
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(gs),
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
		Address: cfg.Registry.Nacos.Address,
		Port:    uint64(cfg.Registry.Nacos.Port),
	}

	// 构造 protobuf conf 对象（供 wire 注入）
	serverConf := &conf.Server{
		Grpc: &conf.Server_GRPC{
			Addr:    cfg.Server.GRPC.Addr,
			Timeout: durationpb.New(cfg.Server.GRPC.Timeout),
		},
	}
	dataConf := &conf.Data{
		Database: &conf.Data_Database{
			Driver: cfg.Data.Database.Driver,
			Source: cfg.Data.Database.Source,
		},
		Redis: &conf.Data_Redis{
			Addr:         cfg.Data.Redis.Addr,
			ReadTimeout:  durationpb.New(cfg.Data.Redis.ReadTimeout),
			WriteTimeout: durationpb.New(cfg.Data.Redis.WriteTimeout),
			Password:     cfg.Data.Redis.Password,
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
			attribute.String("service", "user"),
		)),
	)

	app, cleanup, err := initApp(serverConf, dataConf, logger, tp, nacosConf)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	if err := app.Run(); err != nil {
		panic(err)
	}
}
