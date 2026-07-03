# Kratos Mono-Repo

基于 [go-kratos/kratos](https://github.com/go-kratos/kratos) v2 的微服务单体仓库，采用 DDD 分层架构 + BFF 聚合模式。

## 架构设计

```
                          ┌────────────────────────────────────────────┐
                          │              Nacos (注册/发现)              │
                          └──────────┬──────────────────┬──────────────┘
                                     │ 注册              │ 发现
                                     ▼                  ▼
┌──────────┐    HTTP    ┌──────────────────┐   gRPC    ┌──────────────────┐    ┌──────────┐
│  客户端   │ ────────▶  │  Shop BFF        │ ────────▶ │  User Service    │───▶│  MySQL   │
│ (Web/App)│            │  (HTTP Only)     │           │  (gRPC Only)     │───▶│  Redis   │
└──────────┘            └──────────────────┘           └──────────────────┘    └──────────┘
```

| 服务类型 | 服务名 | 暴露协议 | 有DB | Nacos 注册名 | 说明 |
|---------|--------|---------|------|-------------|------|
| BFF | shop | HTTP | ❌ | `casso.shop.service.http` | 聚合层，发现下游服务 |
| Service | user | gRPC | ✅ | `casso.user.service.grpc` | 基础服务，提供数据操作 |

## 项目结构

```
kratos-mono-repo/
├── api/                    # Proto API 定义
│   ├── shop/v1/            #   Shop BFF API（含 HTTP 注解）
│   └── user/v1/            #   User 服务 API
├── gen/                    # Proto 生成代码（make proto 自动生成）
│   ├── shop/v1/
│   └── user/v1/
├── app/                    # 服务实现
│   ├── shop/               #   Shop BFF 聚合层
│   │   ├── cmd/server/     #     入口 + Wire
│   │   ├── configs/        #     配置文件
│   │   └── internal/       #     内部代码
│   │       ├── biz/        #       业务逻辑（调用下游 gRPC）
│   │       ├── conf/       #       配置结构（Proto）
│   │       ├── server/     #       HTTP Server + gRPC Client
│   │       └── service/    #       服务层（实现 HTTP Handler）
│   └── user/               #   User 基础服务
│       ├── cmd/server/     #     入口 + Wire
│       ├── configs/        #     配置文件
│       └── internal/       #     内部代码
│           ├── biz/        #       业务逻辑（定义 Repo 接口）
│           ├── conf/       #       配置结构（Proto）
│           ├── data/       #       数据层（MySQL + Redis，实现 Repo）
│           ├── model/      #       数据模型
│           ├── server/     #       gRPC Server
│           └── service/    #       服务层（实现 gRPC Handler）
├── pkg/                    # 公共代码
│   ├── auth/               #   JWT + 密码加密
│   ├── bigint/             #   大数运算
│   ├── contextx/           #   Context 辅助
│   ├── crypto/             #   MD5, Snowflake
│   ├── datacopy/           #   数据复制
│   ├── encoding/           #   自定义 HTTP 响应编码
│   ├── errors/             #   统一错误码
│   ├── lottery/            #   加权随机抽奖
│   ├── middleware/         #   通用中间件（Auth, CORS）
│   ├── pagination/         #   分页
│   ├── registry/           #   Nacos 服务注册/发现封装
│   ├── strx/               #   字符串工具
│   └── transaction/        #   GORM 事务封装
├── deploy/                 # 部署配置
│   ├── docker/             #   Dockerfile
│   ├── docker-compose/     #   本地开发环境
│   └── kubernetes/         #   K8s 部署
├── third_party/            # 第三方 Proto 依赖
├── tools/                  # 开发工具
│   └── kratos-gen/         #   脚手架代码生成器
├── Makefile                # 全局构建入口
├── makefile.svc            # 服务共享 Makefile
├── go.mod                  # Go 模块（module: casso）
└── go.sum
```

## 快速开始

### 前置要求

| 工具 | 版本要求 | 说明 |
|------|---------|------|
| Go | ≥ 1.22 | 项目使用 go 1.25.0 |
| Docker | 最新 | 用于本地基础设施 |
| protoc | 最新 | Protocol Buffers 编译器 |

### 1. 克隆项目

```bash
git clone <your-repo-url> kratos-mono-repo
cd kratos-mono-repo
```

### 2. 安装开发工具

安装 protoc 相关 Go 插件和 Wire：

```bash
make init
```

这会安装以下工具：
- `protoc-gen-go` — 生成 Go 结构体
- `protoc-gen-go-grpc` — 生成 gRPC 客户端/服务端
- `protoc-gen-go-http` — 生成 Kratos HTTP 服务端
- `protoc-gen-go-errors` — 生成 Kratos 错误码
- `protoc-gen-openapiv2` — 生成 Swagger 文档
- `protoc-gen-validate` — 生成校验代码
- `wire` — 依赖注入代码生成

### 3. 启动基础设施

使用 Docker Compose 启动 MySQL、Redis、Nacos：

```bash
make initdb
```

服务启动后：
- **MySQL**: `127.0.0.1:3307`，用户名 `root`，密码 `password`，数据库 `goweb`
- **Redis**: `127.0.0.1:6379`
- **Nacos 控制台**: http://127.0.0.1:8848/nacos/ （账号密码: `nacos`/`nacos`）

### 4. 生成 Proto 代码

```bash
make proto
```

这会将 `api/` 下的 `.proto` 文件编译到 `gen/` 目录。

### 5. 生成内部配置 Proto

```bash
make config
```

这会编译 `app/*/internal/conf/conf.proto` 生成 Go 配置结构体。

> 也可以一步完成: `make all`（等于 `make proto && make config`）

### 6. 生成 Wire 依赖注入

```bash
make wire
```

### 7. 运行服务

```bash
# 先启动基础服务（User Service），它需要连接 MySQL/Redis
make -C app/user run

# 再启动 BFF 层（Shop），它通过 Nacos 发现 User 服务
make -C app/shop run
```

**启动顺序很重要**：基础服务需要先启动并注册到 Nacos，BFF 层才能发现并调用它们。

### 8. 验证服务

```bash
# 测试 Shop BFF HTTP 接口
curl -X POST http://localhost:8000/v1/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"123456"}'
```

在 Nacos 控制台中可以看到：
- `casso.shop.service.http` — Shop BFF 的 HTTP 服务实例
- `casso.user.service.grpc` — User 服务的 gRPC 服务实例

## 创建新服务

使用脚手架生成器创建新服务：

```bash
# 创建新的 BFF 聚合层（HTTP Only，无数据库）
make new name=admin type=bff

# 创建新的基础服务（gRPC Only，有数据库）
make new name=order type=service
```

或直接使用生成器：

```bash
go run tools/kratos-gen/main.go --name=payment --type=service
```

生成器会根据服务类型创建不同的结构：

**BFF (type=bff)**:
- HTTP Server + gRPC Client（调用下游服务）
- 无 Data 层，无数据库
- Biz 层直接依赖下游 gRPC 客户端

**Service (type=service)**:
- gRPC Server（对外提供服务）
- Data 层（MySQL + Redis）
- Biz 层定义 Repo 接口，Data 层实现

创建后需要：
1. 编辑 `api/{name}/v1/{name}.proto` 定义 API
2. 运行 `make proto` 生成代码
3. 运行 `make wire` 生成依赖注入
4. 运行 `make -C app/{name} run` 启动服务

## Makefile 命令

### 全局命令（项目根目录）

| 命令 | 说明 |
|------|------|
| `make init` | 安装 protoc 相关工具 |
| `make proto` | 生成所有 API proto 代码到 gen/ |
| `make config` | 生成内部配置 proto 结构体 |
| `make all` | proto + config |
| `make build` | 编译所有服务 |
| `make wire` | 生成所有服务的 Wire 依赖注入 |
| `make test` | 运行所有测试 |
| `make new name=X type=Y` | 创建新服务 |
| `make initdb` | 启动基础设施（Docker Compose） |
| `make help` | 显示帮助信息 |

### 服务级命令（在 `app/{name}/` 下执行）

| 命令 | 说明 |
|------|------|
| `make build` | 编译服务 |
| `make wire` | 生成 Wire 依赖注入 |
| `make run` | 编译并运行服务 |
| `make proto` | 生成内部配置 proto |
| `make docker` | 构建 Docker 镜像 |
| `make test` | 运行测试 |
| `make clean` | 清理编译产物 |

## 设计原则

### DDD 分层架构

```
Service 层 → 校验数据，实现 RPC/HTTP 接口
    ↓
Biz 层    → 业务逻辑，定义数据操作接口（Repo interface）
    ↓
Data 层   → 实现 Biz 的 Repo 接口，操作 DB/Redis
```

**严格规则**：不能跨模块调用，依赖倒置通过 Wire 注入。

### BFF 与 Service 的区别

| | BFF | Service |
|--|-----|---------|
| 暴露协议 | HTTP | gRPC |
| 数据层 | 无 | MySQL + Redis |
| Biz 依赖 | 下游 gRPC Client | Repo 接口 |
| Nacos 角色 | 发现者 | 注册者 + 发现者 |
| 适用场景 | 面向前端，聚合多个基础服务 | 提供领域数据操作 |

### 认证流程

外网请求在 API Gateway 进行认证拦截，认证成功后通过 JWT 在 RPC 元数据中传递到 BFF 层。BFF 校验 Token 完整性后将身份信息注入 Context。BFF 到下层微服务，直接在 RPC Request 中传入 UserID。

> 接口请求的业务参数显式由 Request 传递，扩展性更优（同时支持 HTTP 与 RPC 调用）。

### 错误处理

1. 基础服务根据业务抛出合适的错误（在 `pkg/errors/` 中定义）
2. BFF 层解析 kratos 错误: `e := errors.FromError(err)`
3. 错误只在第一次产生的地方输出日志，上层传递无需重复处理

### 依赖注入

使用 `wire` 维护依赖注入（`app/{name}/cmd/server/wire.go`），可通过 `wire_gen.go` 了解依赖调用顺序。

## 技术栈

| 类别 | 技术 |
|------|------|
| 框架 | go-kratos/kratos v2 |
| ORM | GORM + MySQL |
| 缓存 | Redis |
| 服务注册/发现 | Nacos |
| 链路追踪 | OpenTelemetry + Jaeger |
| 依赖注入 | Wire |
| 认证 | JWT |
| 配置 | YAML + Protobuf |
| 部署 | Docker + Kubernetes |

## 注意事项

- **macOS 用户**：必须使用 `CGO_ENABLED=0` 编译运行，否则会出现 dyld crash
- **Go 1.25 兼容性**：使用 `go build + 执行` 代替 `go run`，使用 `go generate` 代替直接调用 `wire` 命令
- **Nacos**：单例模式后台 http://127.0.0.1:8848/nacos/ （nacos:nacos）
- **gRPC 调试**：推荐使用 [`grpcui`](https://github.com/fullstorydev/grpcui)

## Commit 规范

格式: `type(scope): 描述`

| type | 说明 |
|------|------|
| `feat` | 添加新功能 |
| `fix` | 修补缺陷 |
| `docs` | 修改文档 |
| `style` | 格式调整（不影响逻辑） |
| `refactor` | 重构 |
| `test` | 测试相关 |
| `chore` | 构建/工具变更 |

## License

See [LICENSE](LICENSE)
