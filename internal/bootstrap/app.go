package bootstrap

import (
	articleApi "blog-2026ddd-server/internal/article/api"
	articleApp "blog-2026ddd-server/internal/article/application"
	articleInfra "blog-2026ddd-server/internal/article/infrastructure"
	sharedApi "blog-2026ddd-server/internal/shared/api"
	"log"
	"net/http"

	"blog-2026ddd-server/internal/infrastructure"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type App struct {
	Router   *http.ServeMux
	Config   *infrastructure.Config
	Database *gorm.DB
	Server   *http.Server
	Redis    *redis.Client
}

func (app *App) Run() {
	addr := ":" + app.Config.HTTPPort
	server := &http.Server{
		Addr:    addr,
		Handler: app.Router,
	}
	log.Printf(
		"HTTP已服务启动: http://localhost%s",
		addr,
	)
	log.Printf(
		"OpenAPI文档: http://localhost%s/docs",
		addr,
	)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

// NewApp 这里负责组装整个系统
func NewApp() *App {
	cfg := infrastructure.LoadConfig()
	sharedApi.ConfigureHumaErrors()

	db, err := infrastructure.NewPostgres(cfg.Database)

	if err != nil {
		panic(err)
	}

	router := http.NewServeMux()
	config := huma.DefaultConfig("My API", "1.0.0")
	config.CreateHooks = nil
	api := humago.New(router, config)
	_redis, _ := infrastructure.NewRedis(cfg.Redis)

	// 初始化模块
	articleRepo := articleInfra.NewArticleRepository(db)
	articleTypeRepo := articleInfra.NewArticleTypeRepository(db)
	articleSvc := articleApp.NewArticleService(articleRepo, articleRepo)
	articleTypeSvc := articleApp.NewArticleTypeService(articleTypeRepo)
	articleHandler := articleApi.NewArticleHandler(articleSvc)
	articleTypeHandler := articleApi.NewArticleTypeHandler(articleTypeSvc)
	articleApi.RegisterRoutes(articleHandler, articleTypeHandler, api)

	return &App{
		Router:   router,
		Config:   cfg,
		Database: db,
		Redis:    _redis,
	}

}
