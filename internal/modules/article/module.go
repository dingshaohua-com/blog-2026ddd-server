package article

import (
	"blog-2026ddd-server/internal/modules/article/api"
	"blog-2026ddd-server/internal/modules/article/application"
	"blog-2026ddd-server/internal/modules/article/infrastructure"
	"blog-2026ddd-server/internal/modules/article/query"

	"github.com/danielgtaylor/huma/v2"
	"gorm.io/gorm"
)

func RegisterModule(db *gorm.DB, serverApi huma.API) {

	// 1. 写路径依赖组装 (Command Side)
	articleRepo := infrastructure.NewArticleRepository(db)
	articleTypeRepo := infrastructure.NewArticleTypeRepository(db)

	articleCmdSvc := application.NewArticleService(articleRepo) // 可以叫 CmdSvc，强调写路径
	articleTypeCmdSvc := application.NewArticleTypeService(articleTypeRepo)

	// 2. 读路径依赖组装 (Query Side - 直连 db，没有任何其他依赖！)
	articleQuerySvc := query.NewArticleService(db) // 推荐把结构体/构造命名为 Service 或 QueryService

	// 3. API 层同时注入【写服务】与【读服务】
	articleHandler := api.NewArticleHandler(articleCmdSvc, articleQuerySvc)
	articleTypeHandler := api.NewArticleTypeHandler(articleTypeCmdSvc)

	api.RegisterRoutes(articleHandler, articleTypeHandler, serverApi)
}
