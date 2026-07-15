## 目录结构
article/comment/mood/category 为业务模块，bootstrap 和 infrastructure 不是业务模块，而是支撑业务模块运行的“底座”。
bootstrap为应用启动/依赖组装，infrastructure为全局技术设施（业务模块和非业务支撑模块都可以放在 internal/ 下平级；平级不代表它们都是业务，只代表它们都是项目内部包。）。
```
internal/
├── article/          # 业务：文章
├──── 
├── comment/          # 业务：评论
├── mood/             # 业务：说说
├── category/         # 业务：分类
│
├── bootstrap/        # 把所有模块装起来
└── infrastructure/   # 创建工具：全局 DB、Redis、Logger、Config
```

模块也可以拥有自己的，比如这里的文章模块`article/infrastructure/mysql_repository.go`

```
Handler
    ↓
application.Service
    ↓
domain.Repository（接口）
    ↑
Repository（GORM实现）
    ↓
gorm.DB
```

每个业务模块下
domain（领域的意思）定义了该领域的实体和实体有哪些方法的抽象接口（至少最基础的CRUD），domain 应尽量保持“纯净”，不直接依赖 HTTP 框架、数据库 ORM、Redis、消息队列等技术细节
然后 infrastructure（基建的意思） 模块具体会实现这些方法，在 DDD 中，它负责实现具体的技术细节

application（应用的意思） 里主要是 Application Service（应用服务），组合业务流程，其中的
service是根据handler来组合业务逻辑，service依赖infrastructure提供的能力，但是参数确是domain的repository,这也是接口的体现


internal/
├── article/       # 文章业务模块
├── home/          # 首页业务模块
├── shared/        # 跨模块共享代码
├── bootstrap/     # 程序初始化、依赖组装
└── infrastructure/# 数据库连接、日志、配置