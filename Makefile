GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
export PATH:=$(GOPATH)/bin:$(PATH)

# 服务列表（自动发现 app/ 下的目录）
SERVICES=$(shell ls -d app/*/ | xargs -n1 basename)

# Proto 文件搜索路径
API_PROTO_FILES=$(shell find api -name "*.proto")
INTERNAL_PROTO_FILES=$(shell find app -name *.proto)

.PHONY: init
# init 安装 protoc 相关工具
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	go install github.com/envoyproxy/protoc-gen-validate@latest
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest

.PHONY: proto
# proto 生成所有 API proto 代码到 gen/
proto:
	@echo "==> Generating API proto files..."
	@protoc --proto_path=. \
	       --proto_path=./third_party \
	       --go_out=module=casso:. \
	       --go-http_out=module=casso:. \
	       --go-grpc_out=module=casso:. \
	       --go-errors_out=module=casso:. \
	       --validate_out=module=casso,lang=go:. \
	       $(API_PROTO_FILES)
	@echo "==> Generating swagger files..."
	@for proto in $(API_PROTO_FILES); do \
		dir=$$(dirname $$proto); \
		svc=$$(basename $$(dirname $$dir)); \
		outdir=gen/$$svc/v1; \
		mkdir -p $$outdir; \
		protoc --proto_path=. \
		       --proto_path=./third_party \
		       --openapiv2_out=./$$outdir \
		       --openapiv2_opt=logtostderr=true \
		       --openapiv2_opt=json_names_for_fields=false \
		       $$proto 2>/dev/null; \
		for nested in $$(find ./$$outdir/api -name "*.swagger.json" 2>/dev/null); do \
			fname=$$(basename $$nested); \
			mv "$$nested" "./$$outdir/$$fname"; \
		done; \
		rm -rf "./$$outdir/api" ; \
		echo "  Generated: $$outdir/$$svc.swagger.json"; \
	done

.PHONY: errors
# errors 生成错误代码
errors:
	@echo "==> Generating error proto files..."
	@protoc --proto_path=. \
	       --proto_path=./third_party \
	       --go_out=module=casso:. \
	       --go-errors_out=module=casso:. \
	       $(API_PROTO_FILES)

.PHONY: config
# config 生成内部 proto 配置结构
config:
	@echo "==> Generating internal config proto files..."
	@protoc --proto_path=. \
	       --proto_path=./third_party \
	       --go_out=paths=source_relative:. \
	       $(INTERNAL_PROTO_FILES)

.PHONY: all
# all 生成所有代码
all: proto config

.PHONY: build
# build 编译所有服务
build:
	@for svc in $(SERVICES); do \
		echo "==> Building $$svc..."; \
		make -C app/$$svc build; \
	done

.PHONY: wire
# wire 生成所有服务的 Wire 依赖注入
wire:
	@for svc in $(SERVICES); do \
		echo "==> Wiring $$svc..."; \
		cd app/$$svc/cmd/server && go generate && cd $(PWD); \
	done

.PHONY: test
# test 运行所有测试
test:
	go test -v ./... -cover

.PHONY: new
# new 创建新服务，用法: make new name=order type=service
new:
	@go run tools/kratos-gen/main.go --name=$(name) --type=$(type)

.PHONY: initdb
# initdb 启动基础设施（MySQL, Redis, Nacos 等）
initdb:
	cd deploy/docker-compose && docker-compose up -d

.PHONY: help
# help 显示帮助信息
help:
	@echo ''
	@echo 'Usage:'
	@echo '  make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
		helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
