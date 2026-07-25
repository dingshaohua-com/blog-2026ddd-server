package article

import (
	"blog-2026ddd-server/internal/modules/article/api"
	"blog-2026ddd-server/internal/modules/article/application"
	"blog-2026ddd-server/internal/modules/article/infrastructure"

	"github.com/danielgtaylor/huma/v2"
	"gorm.io/gorm"
)

func RegisterModule(db *gorm.DB, serverApi huma.API) {

	articleRepo := infrastructure.NewArticleRepository(db)
	articleTypeRepo := infrastructure.NewArticleTypeRepository(db)

	articleQuery := infrastructure.NewArticleQuery(db)
	articleSvc := application.NewArticleService(articleRepo, articleQuery)
	articleTypeSvc := application.NewArticleTypeService(articleTypeRepo)

	articleHandler := api.NewArticleHandler(articleSvc)
	articleTypeHandler := api.NewArticleTypeHandler(articleTypeSvc)

	api.RegisterRoutes(articleHandler, articleTypeHandler, serverApi)
}
