### Kratos-Mono-Repo

####  前置条件
* [`Kratos`](https://go-kratos.dev/en/docs/getting-started/start/) `docker  docker-compose` 正确安装

#### 启动
1. `make initdb`  // 初始化环境 
2. `make runserver SERVER=你的服务名称` // 运行服务，服务需要先搭建完毕才能启动
3. `docker`文件位于`/deply`目录下，路径不对的自行切换

#### 新增服务
* 新增`payment`服务:
`make newapp APPNAME=yourServerName`

```
以上脚本仅仅帮助我们创建目录，具体的go文件需要我们自己处理，官方也并不提供相关的工具
关于每个目录的作用，可以参考user || shop 服务,临时的脚本，后面再优化
```

#### 错误处理
1. 基础服务需要根据业务抛出合适的错误，在`pkg/errors/normal`中自定义`kratos`错误；基础服务应该返回的是`kratos`错误类型
2. BFF层处理(记录日志等操作)具体的错误并返回请求，解析`kratos`错误得到具体底层抛出的错误信息:`e := errors.FromError(err)`


#### 服务拆分 （按照业务拆分，服务间通过接口通讯）
1. 示例服务仅包含一个`user`服务
2. `shop`服务充当BFF聚合层，没有DB没有复杂业务操作
3. 以上是`Demo`功能列表（没有全部实现）

---

##### 设计指导思想：
```
对于外网的请求来说，我们通常在 API Gateway 
进行统一的认证拦截，一旦认证成功，我们会使
用 JWT 方式通过 RPC 元数据传递的方式带到
BFF 层，BFF 校验 Token 完整性后把身份信息
注入到应用的 Context 中，BFF 到其他下层的微
服务，建议是直接在 RPC Request 中带入用户
身份信息(UserID)请求服务。
```

* 依赖倒置 
    * 使用`wire`维护依赖注入 (`/service/cmd/server/wire.go`)
    * 可以观察`/service/cmd/server/wire_gen.go` 中的代码。 看到各个依赖之间的调用顺序

---

##### 模块权责
* `.vscode`
    * 个人配置喜好，已经添加到`.gitignore`，只上传首次 

* `/api`
    * 与`app`目录的服务一一对应，维护各个服务的`proto`文件与生成的源码
    * 内部服务一般只需定义`rpc`接口，BFF 服务则需要定义`option (google.api.http)`
    * 也可以去掉`/api`目录，就像`kratos`提供的单例服务`layout`一样，将`api/`放到`app`目录内独立维护

* `/app/${server_name}/service/internel/`
    * 需要在不同模块中注入不同的依赖(`ProviderSet`)
    * `BIZ`， 业务组装，定义数据操作方法，供`Service`层调用
    * `Data`，实现数据库/缓存等能力，实现`Biz`定义的接口，供`BIZ`层调用
    * `Conf`，解析`../../configs`内的`yaml`配置文件，供`Data  Server`使用（这两个目录会进行一些服务初始化需要用到配置信息，内含监听配置文件简单用法）
    * `Server`，服务注册，提供服务初始化依赖注入
    * `Service`，校验数据，调用`BIZ`组装好的用例，实现GRPC接口
    * 各个模块不能乱跨模块调用: `Service`实现RPC ==> `Biz`定义数据操作方法并组装业务 ==> `Data`数据操作

* `/deploy`
    * 存放一些自动化脚本
    * 本地测试使用起来挺方便

* `pkg`
    * 存放项目下所有服务均可使用的公共代码，这也是集中式管理的好处之一
    * 在这里添加公共代码需要按照规范分好目录层级（语意化），不能一坨放一起

* `third_party`
    * 存放三方依赖，目前是一些`proto`文件
    * 查看`third_party/google/app/http.proto`可以了解到接口的配置跟参数匹配机制

* `app_makefile`
    * `app`目录下每个服务根路径都会有一个`Makefile`文件，内部 `include ../../../app_makefile`
    * 最好不要私自改动`app_makefile`,需要添加功能可以在自己维护的服务内的`Makefile`中添加
    

* `Makefile`
    * 整个项目根目录的`Makefile`
    * 维护一些全局的`Makefile` 命令

* 其他目录自行了解

---

#### 其他
* 目前有配套后台管理的`consul nacos`
* 服务注册与发现使用`Nacos`，记录`Nacos`单例模式后台地址:`http://127.0.0.1:8848/nacos/` `login = nacos:nacos`
* 生产推荐 `gitlab`+`k8s`+`lens`

* `bloomrpc`[测试工具](https://github.com/bloomrpc/bloomrpc)
#### Commit 规范

###### 格式如下

* 例：`type(scope)`:本次提交概述 `git commit -m "docs(README.md):添加编码规范"`
* `type`: 本次 commit 的类型，诸如 bugfix docs style 等，参考如下:   

    * `feat`：添加新功能
    * `fix`：修补缺陷
    * `docs`：修改文档
    * `style`：修改格式
    * `refactor`：重构
    * `perf`：优化
    * `test`：增加测试
    * `chore`：构建过程或辅助工具的变动
    * `revert`：回滚到上一个版本
    * `merge`：合并

* `scope`: 本次 `commit` 波及的范围
* `subject`: 简明扼要的阐述下本次 `commit` 的主旨，在原文中特意强调了几点：

    1. 使用祈使句
    2. 首字母不要大写
    3. 结尾无需添加标点

---
##### 分支记录
* `dev`：开发版 `prod`：发布版 `stage`：稳定版
* 命令记录：
```
git checkout -b dev   // 新建分支
git push origin dev   // 推送新分支到线上
git branch -d dev     // 删除本地分支
git push origin :dev  // 删除远程分支,需要设置权限

git checkout pro // 切换到发布分支
git  merge dev   // 再合并开发分支
```
