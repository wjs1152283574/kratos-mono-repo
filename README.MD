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
    

* `Makefile` : `generate` 源文件
    1.  `make help` 查看都有哪些命令
    2. `make all` 时需要存在至少一个源文件
    3. 自行查看`Makefile`，可以自己修改
    4. 如果不修改项目结构，一般 `make all`即可