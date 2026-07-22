package article

import (
	articleApi "blog-2026ddd-server/internal/modules/article/api"
	"blog-2026ddd-server/internal/modules/article/application"
	articleInfra "blog-2026ddd-server/internal/modules/article/infrastructure"

	"github.com/danielgtaylor/huma/v2"
	"gorm.io/gorm"
)

func RegisterModule(db *gorm.DB, api huma.API) {
	articleTypeRepo := articleInfra.NewArticleTypeRepository(db)
	articleRepo := articleInfra.NewArticleRepository(db)
	articleQuery := articleInfra.NewArticleQuery(db)
	articleSvc := application.NewArticleService(articleRepo, articleQuery)
	articleTypeSvc := application.NewArticleTypeService(articleTypeRepo)
	articleHandler := articleApi.NewArticleHandler(articleSvc)
	articleTypeHandler := articleApi.NewArticleTypeHandler(articleTypeSvc)
	articleApi.RegisterRoutes(articleHandler, articleTypeHandler, api)
}
