package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// Config 生成器配置
type Config struct {
	Name         string
	Type         string // bff 或 service
	Module       string
	RootDir      string
	TemplateName string
}

// FileTemplate 文件模板
type FileTemplate struct {
	Path     string
	Template string
}

func main() {
	name := flag.String("name", "", "Service name (lowercase, e.g. order)")
	svcType := flag.String("type", "service", "Service type: bff or service")
	module := flag.String("module", "casso", "Go module name")
	rootDir := flag.String("root", ".", "Project root directory")
	flag.Parse()

	if *name == "" {
		fmt.Println("Error: --name is required")
		fmt.Println("Usage: go run tools/kratos-gen/main.go --name=order --type=service")
		os.Exit(1)
	}
	if *svcType != "bff" && *svcType != "service" {
		fmt.Println("Error: --type must be 'bff' or 'service'")
		os.Exit(1)
	}

	config := Config{
		Name:         *name,
		Type:         *svcType,
		Module:       *module,
		RootDir:      *rootDir,
		TemplateName: toPascalCase(*name),
	}

	templates := getFileTemplates(config)

	fmt.Printf("==> Creating service: %s (type: %s)\n", config.Name, config.Type)
	for _, ft := range templates {
		path := renderTemplate(ft.Path, config)
		fullPath := filepath.Join(config.RootDir, path)
		dir := filepath.Dir(fullPath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Printf("Error creating directory %s: %v\n", dir, err)
			continue
		}
		if _, err := os.Stat(fullPath); err == nil {
			fmt.Printf("  Skipping (exists): %s\n", path)
			continue
		}
		content := renderTemplate(ft.Template, config)
		if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
			fmt.Printf("Error writing file %s: %v\n", fullPath, err)
			continue
		}
		fmt.Printf("  Created: %s\n", path)
	}

	fmt.Printf("\n==> Done! Next steps:\n")
	fmt.Printf("  1. Edit api/%s/v1/%s.proto to define your API\n", config.Name, config.Name)
	fmt.Printf("  2. Run: make proto\n")
	fmt.Printf("  3. Run: make wire\n")
	fmt.Printf("  4. Run: make -C app/%s run\n", config.Name)
}

func getFileTemplates(c Config) []FileTemplate {
	common := []FileTemplate{
		{Path: "api/{{.Name}}/v1/{{.Name}}.proto", Template: protoAPITemplate},
		{Path: "api/{{.Name}}/v1/{{.Name}}_error.proto", Template: protoErrorTemplate},
		{Path: "app/{{.Name}}/cmd/server/main.go", Template: mainTemplate},
		{Path: "app/{{.Name}}/cmd/server/wire.go", Template: wireTemplate},
		{Path: "app/{{.Name}}/generate.go", Template: generateTemplate},
		{Path: "app/{{.Name}}/Makefile", Template: makefileTemplate},
		{Path: "app/{{.Name}}/configs/config.yaml", Template: configTemplate},
		{Path: "app/{{.Name}}/internal/conf/conf.proto", Template: confProtoTemplate},
		{Path: "app/{{.Name}}/internal/biz/biz.go", Template: bizGoTemplate},
		{Path: "app/{{.Name}}/internal/biz/{{.Name}}.go", Template: bizLogicTemplate},
		{Path: "app/{{.Name}}/internal/server/server.go", Template: serverGoTemplate},
		{Path: "app/{{.Name}}/internal/service/service.go", Template: serviceGoTemplate},
		{Path: "app/{{.Name}}/internal/service/{{.Name}}.go", Template: serviceLogicTemplate},
	}

	if c.Type == "bff" {
		common = append(common,
			FileTemplate{Path: "app/{{.Name}}/internal/server/http.go", Template: bffHTTPTemplate},
			FileTemplate{Path: "app/{{.Name}}/internal/server/client.go", Template: bffClientTemplate},
		)
	} else {
		common = append(common,
			FileTemplate{Path: "app/{{.Name}}/internal/biz/interface.go", Template: interfaceTemplate},
			FileTemplate{Path: "app/{{.Name}}/internal/data/data.go", Template: dataGoTemplate},
			FileTemplate{Path: "app/{{.Name}}/internal/data/{{.Name}}.go", Template: dataRepoTemplate},
			FileTemplate{Path: "app/{{.Name}}/internal/server/grpc.go", Template: svcGRPCTemplate},
			FileTemplate{Path: "app/{{.Name}}/internal/model/model.go", Template: modelTemplate},
			FileTemplate{Path: "app/{{.Name}}/internal/data/mysql.go", Template: mysqlTemplate},
			FileTemplate{Path: "app/{{.Name}}/internal/data/redis.go", Template: redisTemplate},
		)
	}

	return common
}

func renderTemplate(tmplStr string, data Config) string {
	t, err := template.New("").Funcs(template.FuncMap{
		"upper": strings.ToUpper,
	}).Parse(tmplStr)
	if err != nil {
		fmt.Printf("Template parse error: %v\n", err)
		return tmplStr
	}
	var buf strings.Builder
	if err := t.Execute(&buf, data); err != nil {
		fmt.Printf("Template execute error: %v\n", err)
		return tmplStr
	}
	return buf.String()
}

func toPascalCase(s string) string {
	parts := strings.Split(s, "_")
	var result string
	for _, p := range parts {
		if len(p) > 0 {
			result += strings.ToUpper(p[:1]) + p[1:]
		}
	}
	return result
}

// ==================== Proto 模板 ====================

var protoAPITemplate = `syntax = "proto3";
package api.{{.Name}}.v1;
import "validate/validate.proto";
{{- if eq .Type "bff"}}
import "google/api/annotations.proto";
{{- end}}

option go_package = "{{.Module}}/gen/{{.Name}}/v1;v1";

service {{.TemplateName}} {
{{- if eq .Type "bff"}}
    rpc Create (Create{{.TemplateName}}Request) returns (Create{{.TemplateName}}Reply) {
        option (google.api.http) = { post: "/v1/{{.Name}}" body: "*" };
    }
    rpc Get (Get{{.TemplateName}}Request) returns (Get{{.TemplateName}}Reply) {
        option (google.api.http) = { get: "/v1/{{.Name}}/{id}" };
    }
    rpc List (List{{.TemplateName}}Request) returns (List{{.TemplateName}}Reply) {
        option (google.api.http) = { get: "/v1/{{.Name}}" };
    }
{{- else}}
    rpc Create{{.TemplateName}} (Create{{.TemplateName}}Request) returns (Create{{.TemplateName}}Reply);
    rpc Get{{.TemplateName}} (Get{{.TemplateName}}Request) returns (Get{{.TemplateName}}Reply);
    rpc Update{{.TemplateName}} (Update{{.TemplateName}}Request) returns (Update{{.TemplateName}}Reply);
    rpc Delete{{.TemplateName}} (Delete{{.TemplateName}}Request) returns (Delete{{.TemplateName}}Reply);
    rpc List{{.TemplateName}} (List{{.TemplateName}}Request) returns (List{{.TemplateName}}Reply);
{{- end}}
}

message Create{{.TemplateName}}Request { string name = 1; }
message Create{{.TemplateName}}Reply { int64 id = 1; string name = 2; }
message Get{{.TemplateName}}Request { int64 id = 1; }
message Get{{.TemplateName}}Reply { int64 id = 1; string name = 2; }
{{- if ne .Type "bff"}}
message Update{{.TemplateName}}Request { int64 id = 1; string name = 2; }
message Update{{.TemplateName}}Reply { int64 id = 1; string name = 2; }
message Delete{{.TemplateName}}Request { int64 id = 1; }
message Delete{{.TemplateName}}Reply { bool ok = 1; }
{{- end}}
message List{{.TemplateName}}Request { int64 page = 1; int64 limit = 2; }
message List{{.TemplateName}}Reply {
    message {{.TemplateName}}Item { int64 id = 1; string name = 2; }
    repeated {{.TemplateName}}Item items = 1;
}
`

var protoErrorTemplate = `syntax = "proto3";
package api.{{.Name}}.v1;
import "errors/errors.proto";
option go_package = "{{.Module}}/gen/{{.Name}}/v1;v1";

enum {{.TemplateName}}ServiceErrorReason {
    option (errors.default_code) = 500;
    {{.TemplateName | upper}}_RECORD_NOT_FOUND = 0 [(errors.code) = 404];
    {{.TemplateName | upper}}_CONTENT_MISSING = 1 [(errors.code) = 400];
    {{.TemplateName | upper}}_INVALID_PARAMS = 2 [(errors.code) = 400];
}
`

// ==================== Main 模板 ====================

var mainTemplate = `package main

import (
	"flag"
	"os"
	"time"

	"{{.Module}}/app/{{.Name}}/internal/conf"
	nacosreg "{{.Module}}/pkg/registry"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
{{- if eq .Type "bff"}}
	"github.com/go-kratos/kratos/v2/transport/http"
{{- else}}
	"github.com/go-kratos/kratos/v2/transport/grpc"
{{- end}}
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"google.golang.org/protobuf/types/known/durationpb"
	"gopkg.in/yaml.v2"
)

var (
	Name     = "{{.Module}}.{{.Name}}.service"
	Version  = "v1"
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

// AppConfig 普通配置结构体（用于 YAML 反序列化）
type AppConfig struct {
	Trace struct {
		Endpoint string ` + "`" + `yaml:"endpoint"` + "`" + `
	} ` + "`" + `yaml:"trace"` + "`" + `
	Server struct {
{{- if eq .Type "bff"}}
		HTTP struct {
			Addr    string        ` + "`" + `yaml:"addr"` + "`" + `
			Timeout time.Duration ` + "`" + `yaml:"timeout"` + "`" + `
		} ` + "`" + `yaml:"http"` + "`" + `
{{- else}}
		GRPC struct {
			Addr    string        ` + "`" + `yaml:"addr"` + "`" + `
			Timeout time.Duration ` + "`" + `yaml:"timeout"` + "`" + `
		} ` + "`" + `yaml:"grpc"` + "`" + `
{{- end}}
	} ` + "`" + `yaml:"server"` + "`" + `
{{- if ne .Type "bff"}}
	Data struct {
		Database struct {
			Driver string ` + "`" + `yaml:"driver"` + "`" + `
			Source string ` + "`" + `yaml:"source"` + "`" + `
		} ` + "`" + `yaml:"database"` + "`" + `
		Redis struct {
			Addr         string        ` + "`" + `yaml:"addr"` + "`" + `
			ReadTimeout  time.Duration ` + "`" + `yaml:"read_timeout"` + "`" + `
			WriteTimeout time.Duration ` + "`" + `yaml:"write_timeout"` + "`" + `
			Password     string        ` + "`" + `yaml:"password"` + "`" + `
		} ` + "`" + `yaml:"redis"` + "`" + `
	} ` + "`" + `yaml:"data"` + "`" + `
{{- end}}
{{- if eq .Type "bff"}}
	Discovery struct {
		Nacos struct {
			Address string ` + "`" + `yaml:"address"` + "`" + `
			Port    int    ` + "`" + `yaml:"port"` + "`" + `
		} ` + "`" + `yaml:"nacos"` + "`" + `
	} ` + "`" + `yaml:"discovery"` + "`" + `
{{- else}}
	Registry struct {
		Nacos struct {
			Address string ` + "`" + `yaml:"address"` + "`" + `
			Port    int    ` + "`" + `yaml:"port"` + "`" + `
		} ` + "`" + `yaml:"nacos"` + "`" + `
	} ` + "`" + `yaml:"registry"` + "`" + `
{{- end}}
}

{{- if eq .Type "bff"}}

func newApp(logger log.Logger, hs *http.Server, rr registry.Registrar) *kratos.App {
	return kratos.New(
		kratos.Name(Name), kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger), kratos.Server(hs), kratos.Registrar(rr),
	)
}
{{- else}}

func newApp(logger log.Logger, gs *grpc.Server, rr registry.Registrar) *kratos.App {
	return kratos.New(
		kratos.Name(Name), kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger), kratos.Server(gs), kratos.Registrar(rr),
	)
}
{{- end}}

func main() {
	flag.Parse()
	logger := log.With(log.NewStdLogger(os.Stdout),
		"service.name", Name, "service.version", Version,
		"ts", log.DefaultTimestamp, "caller", log.DefaultCaller)

	data, err := os.ReadFile(flagconf + "/config.yaml")
	if err != nil { panic(err) }
	var cfg AppConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil { panic(err) }

	nacosConf := nacosreg.NacosConfig{
{{- if eq .Type "bff"}}
		Address: cfg.Discovery.Nacos.Address,
		Port:    uint64(cfg.Discovery.Nacos.Port),
{{- else}}
		Address: cfg.Registry.Nacos.Address,
		Port:    uint64(cfg.Registry.Nacos.Port),
{{- end}}
	}

	serverConf := &conf.Server{
{{- if eq .Type "bff"}}
		Http: &conf.Server_HTTP{
			Addr: cfg.Server.HTTP.Addr,
			Timeout: durationpb.New(cfg.Server.HTTP.Timeout),
		},
{{- else}}
		Grpc: &conf.Server_GRPC{
			Addr: cfg.Server.GRPC.Addr,
			Timeout: durationpb.New(cfg.Server.GRPC.Timeout),
		},
{{- end}}
	}
{{- if ne .Type "bff"}}
	dataConf := &conf.Data{
		Database: &conf.Data_Database{Driver: cfg.Data.Database.Driver, Source: cfg.Data.Database.Source},
		Redis: &conf.Data_Redis{
			Addr: cfg.Data.Redis.Addr, ReadTimeout: durationpb.New(cfg.Data.Redis.ReadTimeout),
			WriteTimeout: durationpb.New(cfg.Data.Redis.WriteTimeout), Password: cfg.Data.Redis.Password,
		},
	}
{{- end}}

	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(cfg.Trace.Endpoint)))
	if err != nil { panic(err) }
	tp := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exp),
		tracesdk.WithResource(resource.NewSchemaless(
			semconv.ServiceNameKey.String(Name),
			attribute.String("service", "{{.Name}}"),
		)),
	)

{{- if eq .Type "bff"}}
	app, cleanup, err := initApp(serverConf, logger, tp, nacosConf)
{{- else}}
	app, cleanup, err := initApp(serverConf, dataConf, logger, tp, nacosConf)
{{- end}}
	if err != nil { panic(err) }
	defer cleanup()
	if err := app.Run(); err != nil { panic(err) }
}
`

// ==================== Wire 模板 ====================

var wireTemplate = `//go:build wireinject
// +build wireinject
package main

import (
	"{{.Module}}/app/{{.Name}}/internal/biz"
	"{{.Module}}/app/{{.Name}}/internal/conf"
	"{{.Module}}/app/{{.Name}}/internal/server"
	"{{.Module}}/app/{{.Name}}/internal/service"
{{- if ne .Type "bff"}}
	"{{.Module}}/app/{{.Name}}/internal/data"
{{- end}}
	nacosreg "{{.Module}}/pkg/registry"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

{{- if eq .Type "bff"}}
func initApp(*conf.Server, log.Logger, *tracesdk.TracerProvider, nacosreg.NacosConfig) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
{{- else}}
func initApp(*conf.Server, *conf.Data, log.Logger, *tracesdk.TracerProvider, nacosreg.NacosConfig) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
{{- end}}
`

var generateTemplate = `package generate
//go:generate kratos proto client .
`

var makefileTemplate = `APP_NAME := {{.Name}}
include ../../makefile.svc
`

// ==================== Config 模板 ====================

var configTemplate = `trace:
  endpoint: http://127.0.0.1:9080/v1/trace
server:
{{- if eq .Type "bff"}}
  http:
    addr: 0.0.0.0:8000
    timeout: 1s
{{- else}}
  grpc:
    addr: 0.0.0.0:9000
    timeout: 1s
{{- end}}
{{- if ne .Type "bff"}}
data:
  database:
    driver: mysql
    source: root:password@tcp(127.0.0.1:3307)/goweb?parseTime=true
  redis:
    addr: 127.0.0.1:6379
    read_timeout: 0.2s
    write_timeout: 0.2s
{{- end}}
{{- if eq .Type "bff"}}
discovery:
{{- else}}
registry:
{{- end}}
  nacos:
    address: 127.0.0.1
    port: 8848
`

var confProtoTemplate = `syntax = "proto3";
package {{.Name}}.api;
option go_package = "{{.Module}}/app/{{.Name}}/internal/conf;conf";
import "google/protobuf/duration.proto";

message Bootstrap { Trace trace = 1; Server server = 2; {{- if ne .Type "bff"}} Data data = 3; {{- end}} }
message Trace { string endpoint = 1; }
message Server {
{{- if eq .Type "bff"}}
  message HTTP { string network = 1; string addr = 2; google.protobuf.Duration timeout = 3; }
  HTTP http = 1;
{{- else}}
  message GRPC { string network = 1; string addr = 2; google.protobuf.Duration timeout = 3; }
  GRPC grpc = 1;
{{- end}}
}
{{- if ne .Type "bff"}}
message Data {
  message Database { string driver = 1; string source = 2; }
  message Redis { string network = 1; string addr = 2; google.protobuf.Duration read_timeout = 3; google.protobuf.Duration write_timeout = 4; string password = 5; }
  Database database = 1; Redis redis = 2;
}
{{- end}}
`

// ==================== Biz 模板 ====================

var bizGoTemplate = `package biz

import (
{{- if eq .Type "bff"}}
	v1 "{{.Module}}/gen/user/v1"
{{- end}}
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(New{{.TemplateName}}UseCase)

type {{.TemplateName}}UseCase struct {
{{- if eq .Type "bff"}}
	log *log.Helper
	uc  v1.UserClient
{{- else}}
	repo {{.TemplateName}}Repo
	log  *log.Helper
{{- end}}
}

func New{{.TemplateName}}UseCase({{ if eq .Type "bff" }}logger log.Logger, uc v1.UserClient{{ else }}repo {{.TemplateName}}Repo, logger log.Logger{{ end }}) *{{.TemplateName}}UseCase {
	return &{{.TemplateName}}UseCase{
{{- if eq .Type "bff"}}
		log: log.NewHelper(log.With(logger, "module", "usecase/{{.Name}}")),
		uc:  uc,
{{- else}}
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "usecase/{{.Name}}")),
{{- end}}
	}
}
`

var bizLogicTemplate = `package biz

import "context"

{{- if eq .Type "bff"}}

// Get{{.TemplateName}} 获取信息（BFF 通过 gRPC 调用下游服务）
func (uc *{{.TemplateName}}UseCase) Get{{.TemplateName}}(ctx context.Context, id int64) error {
	// TODO: 调用下游基础服务
	return nil
}
{{- else}}

func (uc *{{.TemplateName}}UseCase) Create{{.TemplateName}}(ctx context.Context, name string) (int64, error) {
	return uc.repo.Create(ctx, name)
}
func (uc *{{.TemplateName}}UseCase) Get{{.TemplateName}}(ctx context.Context, id int64) (*{{.TemplateName}}Item, error) {
	return uc.repo.Get(ctx, id)
}
func (uc *{{.TemplateName}}UseCase) Update{{.TemplateName}}(ctx context.Context, id int64, name string) error {
	return uc.repo.Update(ctx, id, name)
}
func (uc *{{.TemplateName}}UseCase) Delete{{.TemplateName}}(ctx context.Context, id int64) error {
	return uc.repo.Delete(ctx, id)
}
func (uc *{{.TemplateName}}UseCase) List{{.TemplateName}}(ctx context.Context, pageNum, pageSize int64) ([]*{{.TemplateName}}Item, error) {
	return uc.repo.List(ctx, pageNum, pageSize)
}
{{- end}}
`

var interfaceTemplate = `package biz

import "context"

type {{.TemplateName}}Repo interface {
	Create(ctx context.Context, name string) (int64, error)
	Get(ctx context.Context, id int64) (*{{.TemplateName}}Item, error)
	Update(ctx context.Context, id int64, name string) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, pageNum, pageSize int64) ([]*{{.TemplateName}}Item, error)
}

type {{.TemplateName}}Item struct {
	ID   int64
	Name string
}
`

// ==================== Data 模板 ====================

var dataGoTemplate = `package data

import (
	"{{.Module}}/app/{{.Name}}/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/gorm"
	_ "github.com/go-sql-driver/mysql"
)

var ProviderSet = wire.NewSet(NewData, NewDB, NewRd, New{{.TemplateName}}Repo)

type Data struct { rd *redis.Client; db *gorm.DB; log *log.Helper }

func NewData(db *gorm.DB, rd *redis.Client, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "{{.Name}}-service/data"))
	d := &Data{rd: rd, db: db, log: log}
	return d, func() {
		sqldb, _ := d.db.DB()
		if sqldb != nil { sqldb.Close() }
		d.rd.Close()
		log.Info("resource close successed !")
	}, nil
}
`

var dataRepoTemplate = `package data

import (
	"context"
	"{{.Module}}/app/{{.Name}}/internal/biz"
	"{{.Module}}/app/{{.Name}}/internal/model"
	"{{.Module}}/pkg/pagination"
	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.{{.TemplateName}}Repo = (*{{.TemplateName}}Repo)(nil)

type {{.TemplateName}}Repo struct { data *Data; log *log.Helper }

func New{{.TemplateName}}Repo(data *Data, logger log.Logger) biz.{{.TemplateName}}Repo {
	return &{{.TemplateName}}Repo{data: data, log: log.NewHelper(log.With(logger, "module", "data/{{.Name}}"))}
}

func (r *{{.TemplateName}}Repo) Create(ctx context.Context, name string) (int64, error) {
	m := &model.{{.TemplateName}}{Name: name}
	if err := r.data.db.WithContext(ctx).Create(m).Error; err != nil { return 0, err }
	return int64(m.ID), nil
}
func (r *{{.TemplateName}}Repo) Get(ctx context.Context, id int64) (*biz.{{.TemplateName}}Item, error) {
	var m model.{{.TemplateName}}
	if err := r.data.db.WithContext(ctx).First(&m, id).Error; err != nil { return nil, err }
	return &biz.{{.TemplateName}}Item{ID: int64(m.ID), Name: m.Name}, nil
}
func (r *{{.TemplateName}}Repo) Update(ctx context.Context, id int64, name string) error {
	return r.data.db.WithContext(ctx).Model(&model.{{.TemplateName}}{}).Where("id = ?", id).Update("name", name).Error
}
func (r *{{.TemplateName}}Repo) Delete(ctx context.Context, id int64) error {
	return r.data.db.WithContext(ctx).Delete(&model.{{.TemplateName}}{}, id).Error
}
func (r *{{.TemplateName}}Repo) List(ctx context.Context, pageNum, pageSize int64) ([]*biz.{{.TemplateName}}Item, error) {
	var list []*model.{{.TemplateName}}
	if err := r.data.db.WithContext(ctx).Limit(int(pageSize)).Offset(int(pagination.GetPageOffset(pageNum, pageSize))).Find(&list).Error; err != nil { return nil, err }
	items := make([]*biz.{{.TemplateName}}Item, 0, len(list))
	for _, m := range list { items = append(items, &biz.{{.TemplateName}}Item{ID: int64(m.ID), Name: m.Name}) }
	return items, nil
}
`

// ==================== Server 模板 ====================

var serverGoTemplate = `package server

import (
	"github.com/google/wire"
	"{{.Module}}/pkg/registry"
{{- if ne .Type "bff"}}
	kr "github.com/go-kratos/kratos/v2/registry"
{{- end}}
)

var ProviderSet = wire.NewSet(
{{- if eq .Type "bff"}}
	NewHTTPServer, NewRegistrar, NewUserServiceClient,
{{- else}}
	NewGRPCServer, NewRegistrar,
{{- end}}
)

{{- if eq .Type "bff"}}
func NewRegistrar(nc registry.NacosConfig) registry.NacosConfig { return nc }
{{- else}}
func NewRegistrar(nc registry.NacosConfig) kr.Registrar { return registry.NewRegistrar(nc) }
{{- end}}
`

var bffHTTPTemplate = `package server

import (
	v1 "{{.Module}}/gen/{{.Name}}/v1"
	"{{.Module}}/app/{{.Name}}/internal/conf"
	"{{.Module}}/app/{{.Name}}/internal/service"
	"{{.Module}}/pkg/encoding"
	"{{.Module}}/pkg/middleware"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/ratelimit"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

func NewHTTPServer(c *conf.Server, logger log.Logger, tp *tracesdk.TracerProvider, s *service.{{.TemplateName}}Service) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(tracing.WithTracerProvider(tp)),
			logging.Server(logger),
			ratelimit.Server(),
		),
	}
	if c.Http.Network != "" { opts = append(opts, http.Network(c.Http.Network)) }
	if c.Http.Addr != "" { opts = append(opts, http.Address(c.Http.Addr)) }
	if c.Http.Timeout != nil { opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration())) }
	opts = append(opts, http.Filter(middleware.CORS()))
	opts = append(opts, http.ResponseEncoder(encoding.ResponseEncoder()))
	srv := http.NewServer(opts...)
	v1.Register{{.TemplateName}}HTTPServer(srv, s)
	return srv
}
`

var bffClientTemplate = `package server

import (
	"context"
	uv1 "{{.Module}}/gen/user/v1"
	"{{.Module}}/pkg/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewUserServiceClient 创建下游服务 gRPC 客户端（通过 Nacos 发现）
func NewUserServiceClient(nc registry.NacosConfig) uv1.UserClient {
	r := registry.NewDiscovery(nc)
	conn, err := grpc.DialInsecure(context.Background(),
		grpc.WithEndpoint("discovery:///{{.Module}}.user.service.grpc"),
		grpc.WithDiscovery(r),
	)
	if err != nil { panic(err) }
	return uv1.NewUserClient(conn)
}
`

var svcGRPCTemplate = `package server

import (
	v1 "{{.Module}}/gen/{{.Name}}/v1"
	"{{.Module}}/app/{{.Name}}/internal/conf"
	"{{.Module}}/app/{{.Name}}/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

func NewGRPCServer(c *conf.Server, logger log.Logger, tp *tracesdk.TracerProvider, s *service.{{.TemplateName}}Service) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(recovery.Recovery(), tracing.Server(tracing.WithTracerProvider(tp)), logging.Server(logger)),
	}
	if c.Grpc.Network != "" { opts = append(opts, grpc.Network(c.Grpc.Network)) }
	if c.Grpc.Addr != "" { opts = append(opts, grpc.Address(c.Grpc.Addr)) }
	if c.Grpc.Timeout != nil { opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration())) }
	srv := grpc.NewServer(opts...)
	v1.Register{{.TemplateName}}Server(srv, s)
	return srv
}
`

// ==================== Service 模板 ====================

var serviceGoTemplate = `package service

import (
	v1 "{{.Module}}/gen/{{.Name}}/v1"
	"{{.Module}}/app/{{.Name}}/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(New{{.TemplateName}}Service)

type {{.TemplateName}}Service struct {
	v1.Unimplemented{{.TemplateName}}Server
	uc  *biz.{{.TemplateName}}UseCase
	log *log.Helper
}

func New{{.TemplateName}}Service(uc *biz.{{.TemplateName}}UseCase, logger log.Logger) *{{.TemplateName}}Service {
	return &{{.TemplateName}}Service{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "service/{{.Name}}")),
	}
}
`

var serviceLogicTemplate = `package service

import (
	"context"
	pb "{{.Module}}/gen/{{.Name}}/v1"
)

{{- if eq .Type "bff"}}

func (s *{{.TemplateName}}Service) Create(ctx context.Context, req *pb.Create{{.TemplateName}}Request) (*pb.Create{{.TemplateName}}Reply, error) {
	return &pb.Create{{.TemplateName}}Reply{}, nil
}
func (s *{{.TemplateName}}Service) Get(ctx context.Context, req *pb.Get{{.TemplateName}}Request) (*pb.Get{{.TemplateName}}Reply, error) {
	return &pb.Get{{.TemplateName}}Reply{}, nil
}
func (s *{{.TemplateName}}Service) List(ctx context.Context, req *pb.List{{.TemplateName}}Request) (*pb.List{{.TemplateName}}Reply, error) {
	return &pb.List{{.TemplateName}}Reply{}, nil
}
{{- else}}

func (s *{{.TemplateName}}Service) Create{{.TemplateName}}(ctx context.Context, req *pb.Create{{.TemplateName}}Request) (*pb.Create{{.TemplateName}}Reply, error) {
	if req.Name == "" { return &pb.Create{{.TemplateName}}Reply{}, pb.Error{{.TemplateName}}InvalidParams("name is required") }
	id, err := s.uc.Create{{.TemplateName}}(ctx, req.Name)
	if err != nil { return &pb.Create{{.TemplateName}}Reply{}, err }
	return &pb.Create{{.TemplateName}}Reply{Id: id, Name: req.Name}, nil
}
func (s *{{.TemplateName}}Service) Get{{.TemplateName}}(ctx context.Context, req *pb.Get{{.TemplateName}}Request) (*pb.Get{{.TemplateName}}Reply, error) {
	if req.Id == 0 { return &pb.Get{{.TemplateName}}Reply{}, pb.Error{{.TemplateName}}InvalidParams("id is required") }
	item, err := s.uc.Get{{.TemplateName}}(ctx, req.Id)
	if err != nil { return &pb.Get{{.TemplateName}}Reply{}, err }
	return &pb.Get{{.TemplateName}}Reply{Id: item.ID, Name: item.Name}, nil
}
func (s *{{.TemplateName}}Service) Update{{.TemplateName}}(ctx context.Context, req *pb.Update{{.TemplateName}}Request) (*pb.Update{{.TemplateName}}Reply, error) {
	if req.Id == 0 { return &pb.Update{{.TemplateName}}Reply{}, pb.Error{{.TemplateName}}InvalidParams("id is required") }
	if err := s.uc.Update{{.TemplateName}}(ctx, req.Id, req.Name); err != nil { return &pb.Update{{.TemplateName}}Reply{}, err }
	return &pb.Update{{.TemplateName}}Reply{Id: req.Id, Name: req.Name}, nil
}
func (s *{{.TemplateName}}Service) Delete{{.TemplateName}}(ctx context.Context, req *pb.Delete{{.TemplateName}}Request) (*pb.Delete{{.TemplateName}}Reply, error) {
	if req.Id == 0 { return &pb.Delete{{.TemplateName}}Reply{}, pb.Error{{.TemplateName}}InvalidParams("id is required") }
	if err := s.uc.Delete{{.TemplateName}}(ctx, req.Id); err != nil { return &pb.Delete{{.TemplateName}}Reply{Ok: false}, nil }
	return &pb.Delete{{.TemplateName}}Reply{Ok: true}, nil
}
func (s *{{.TemplateName}}Service) List{{.TemplateName}}(ctx context.Context, req *pb.List{{.TemplateName}}Request) (*pb.List{{.TemplateName}}Reply, error) {
	if req.Page == 0 || req.Limit == 0 { return &pb.List{{.TemplateName}}Reply{}, pb.Error{{.TemplateName}}InvalidParams("page and limit are required") }
	items, err := s.uc.List{{.TemplateName}}(ctx, req.Page, req.Limit)
	if err != nil { return &pb.List{{.TemplateName}}Reply{}, err }
	reply := &pb.List{{.TemplateName}}Reply{}
	for _, item := range items {
		reply.Items = append(reply.Items, &pb.List{{.TemplateName}}Reply_{{.TemplateName}}Item{Id: item.ID, Name: item.Name})
	}
	return reply, nil
}
{{- end}}
`

// ==================== Model 模板 ====================

var modelTemplate = `package model

var {{.TemplateName}}TableName = "{{.Name}}"

type {{.TemplateName}} struct {
	ID          uint   ` + "`" + `gorm:"primarykey"` + "`" + `
	Name        string ` + "`" + `gorm:"COMMENT:名称"` + "`" + `
	UpdatedTime int64  ` + "`" + `gorm:"type:bigint(20);COMMENT:最后修改时间"` + "`" + `
	CreatedTime int64  ` + "`" + `gorm:"type:bigint(20);COMMENT:创建时间"` + "`" + `
	DeleteTime  int64  ` + "`" + `gorm:"type:bigint(20);COMMENT:删除时间"` + "`" + `
}

func (m *{{.TemplateName}}) TableName() string { return {{.TemplateName}}TableName }
`

// ==================== MySQL 模板 ====================

var mysqlTemplate = `package data

import (
	"time"
	"{{.Module}}/app/{{.Name}}/internal/conf"
	"{{.Module}}/app/{{.Name}}/internal/model"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDB(conf *conf.Data, loggers log.Logger) *gorm.DB {
	log := log.NewHelper(log.With(loggers, "module", "{{.Name}}-service/data/gorm"))
	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil { log.Fatalf("failed opening connection to mysql: %v", err) }
	sqlDB, err := db.DB()
	if err != nil { log.Fatalf("failed making connection-pool: %v", err) }
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	if err := db.AutoMigrate(&model.{{.TemplateName}}{}); err != nil { log.Fatal(err) }
	return db
}
`

// ==================== Redis 模板 ====================

var redisTemplate = `package data

import (
	"{{.Module}}/app/{{.Name}}/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
)

func NewRd(conf *conf.Data, logger log.Logger) *redis.Client {
	log.NewHelper(log.With(logger, "module", "{{.Name}}-service/data/redis"))
	return redis.NewClient(&redis.Options{
		Addr:         conf.Redis.Addr,
		Password:     conf.Redis.Password,
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
		DB:           0,
	})
}
`
