## kratos monorepo
---

#### 目录

* `/api` : 存放`proto`文件及 `generate` 源码，每个服务目录区分
    1. 使用命令：`kratos proto add api/user/v1/user.proto` 添加模版文件
    2. 需要支持`http`需要引入:`import "google/api/annotations.proto";`;
    3. 只有指定`http option` 才会生成对应的源码
    4. 生成源码：`kratos proto client api/user/v1/user.proto`
    5. `go mod tidy` 下载依赖 


* `/app` : 存放各个服务实现代码
    1. 生成模版 : `kratos proto server api/user/v1/user.proto -t app/user/service/internal/service`

* `deploy` : 自动化配置脚本


* `pkg`: 公用部分，多个服务多会用到的


* `third_party`: 第三方依赖

* `Makefile` : `generate` 源文件
    1.  `make help` 查看都有哪些命令
    2. `make all` 时需要存在至少一个源文件
    3. 自行查看`Makefile`，可以自己修改
    4. 如果不修改项目结构，一般 `make all`即可

* 运行:
    1. `go generate ./...` 生成所有proto源码、wire等等
    2. kratos run
    3. 运行之前需要注意配置好各个数据库以及注册中心的远程链接信息

* [示例代码beer-shop](https://go-kratos.github.io/beer-shop/#/design)

#### 设计思路：
* 一般服务并不对外暴露http 接口，需要特定的BFF服务来提供http接口。内部服务间通过RPC通讯
* BFF 维护二级缓存，具体服务维护一级缓存
* 服务降级一般在BFF完成
