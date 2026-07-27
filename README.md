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

[//]: # (https://blog.csdn.net/XingyeLuoyue/article/details/159254112)


对，基本正确。更严谨地标注如下：

- 领域实体（Entity）✅  
  `Post` 有 ID 和生命周期，是明确的实体；`Article` 也是实体雏形，但目前偏贫血。

- 值对象（Value Object）❌  
  暂时没有。`content` 目前仍是 `string`，以后可以抽象成 `PostContent`。

- 聚合与聚合根（Aggregate / Aggregate Root）🟡  
  `Post` 可以视为一个简单的单实体聚合根，但项目尚未显式体现包含多个实体、统一维护一致性的聚合结构。

- 业务规则和不变量✅  
  内容不能为空、不能超过 100 字是业务规则。如果它们要求任何合法 Post 始终满足，也是不变量。  
  不过 `RestorePost` 目前绕过校验，所以该不变量尚未被完全保护。

- 领域行为✅  
  `Post.ChangeContent()` 是明确的领域行为。

- 领域服务（Domain Service）❌  
  目前没有跨多个领域对象、且不适合放进单个实体的领域逻辑。

- 领域事件（Domain Event）❌  
  暂时没有类似 `PostCreated`、`PostContentChanged` 的事件。

- Repository 接口✅  
  `PostRepository`、`ArticleRepository` 和 `ArticleTypeRepository` 都属于领域层声明的持久化抽象。

- 工厂及重建方法✅  
  `NewPost()` 是工厂方法，`RestorePost()` 是重建方法。虽然没有单独的 `PostFactory` 类型，但并不需要为了形式专门创建。

- 领域错误或领域异常✅  
  `ErrPostNotFound`、`ErrPostContentEmpty`、`ErrPostContentTooLong` 都属于领域错误。

因此，最准确的项目现状是：

```text
✅ 领域实体
❌ 值对象
🟡 简单聚合根雏形
✅ 业务规则
🟡 不变量存在，但 RestorePost 尚未完全保护
✅ 领域行为
❌ 领域服务
❌ 领域事件
✅ Repository 接口
✅ 工厂及重建方法
✅ 领域错误
```

另外，DDD 不要求把所有元素都用一遍。当前 `Post` 业务简单，没有 Domain Service 或 Domain Event 很正常；只有出现实际业务需求时再引入。