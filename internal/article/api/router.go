package api

import (
	"github.com/danielgtaylor/huma/v2"
)

func RegisterRoutes(articleHandler *ArticleHandler, api huma.API) {
	//group := engine.Group("/article")
	//huma.Get(api, "", h.List))
	//group.GET("", h.List)
	//r.GET("/:id", h.Detail)
	//r.POST("", h.Create)
	//r.PUT("/:id", h.Update)
	//r.DELETE("/:id", h.Delete)

	articleGroup := huma.NewGroup(api, "/article")
	huma.Get(articleGroup, "", articleHandler.List)

	// 2. 文章分类（ArticleType / Category）相关的路由组
	typeGroup := huma.NewGroup(api, "/article-types") // 或者 /categories
	huma.Get(typeGroup, "", h.List)

	//huma.Get(typeGroup, "", h.ListTypes)
	//huma.Get(typeGroup, "/{id}", h.GetType)
	//huma.Post(typeGroup, "", h.CreateType)
	//huma.Put(typeGroup, "/{id}", h.UpdateType)
	//huma.Delete(typeGroup, "/{id}", h.DeleteType)
}
