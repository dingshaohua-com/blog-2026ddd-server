package post

import (
	"blog-2026ddd-server/internal/modules/post/api"
	"blog-2026ddd-server/internal/modules/post/application"
	"blog-2026ddd-server/internal/modules/post/infrastructure"

	"github.com/danielgtaylor/huma/v2"
	"gorm.io/gorm"
)

func RegisterModule(db *gorm.DB, serverApi huma.API) {
	repo := infrastructure.NewPostRepository(db)
	serv := application.NewPostService(repo)
	hand := api.NewPostHandler(serv)
	api.RegisterRoutes(hand, serverApi)
}
