package bootstrap

import (
	articleApi "blog-2026ddd-server/internal/article/api"
	articleApp "blog-2026ddd-server/internal/article/application"
	articleInfra "blog-2026ddd-server/internal/article/infrastructure"
	"net/http"

	"blog-2026ddd-server/internal/infrastructure"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type App struct {
	Router   *gin.Engine
	Config   *infrastructure.Config
	Database *gorm.DB
	Server   *http.Server
	Redis    *redis.Client
}

func (app *App) Run() {
	err := app.Router.Run(":" + app.Config.HTTPPort)
	if err != nil {
		panic(err)
	}
}

// NewApp 这里负责组装整个系统
func NewApp() *App {
	cfg := infrastructure.LoadConfig()

	db, err := infrastructure.NewPostgres(cfg.Database)

	if err != nil {
		panic(err)
	}

	router := gin.Default()
	_redis, _ := infrastructure.NewRedis(cfg.Redis)

	// 初始化模块
	articleRepo := articleInfra.NewRepository(db)
	articleSvc := articleApp.NewService(articleRepo)
	articleHandler := articleApi.NewHandler(articleSvc)
	articleApi.RegisterRoutes(router, articleHandler)

	return &App{
		Router:   router,
		Config:   cfg,
		Database: db,
		Redis:    _redis,
	}

}
