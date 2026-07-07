## 目录结构
article/comment/mood/category 为业务模块，bootstrap 和 infrastructure 不是业务模块，而是支撑业务模块运行的“底座”。
bootstrap为应用启动/依赖组装，infrastructure为全局技术设施（业务模块和非业务支撑模块都可以放在 internal/ 下平级；平级不代表它们都是业务，只代表它们都是项目内部包。）。
```
internal/
├── article/          # 业务：文章
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
Application Service
    ↓
domain.Repository（接口）
    ↑
Repository（GORM实现）
    ↓
gorm.DB
```