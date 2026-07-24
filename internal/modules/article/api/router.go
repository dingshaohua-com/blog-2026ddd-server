package api

import (
	"github.com/danielgtaylor/huma/v2"
)

func RegisterRoutes(articleHandler *ArticleHandler, articleTypeHandler *ArticleTypeHandler, api huma.API) {
	//group := engine.Group("/article")
	//huma.Get(api, "", h.List))
	//group.GET("", h.List)
	//r.GET("/:id", h.Detail)
	//r.POST("", h.Create)
	//r.PUT("/:id", h.Update)
	//r.DELETE("/:id", h.Delete)

	articleGroup := huma.NewGroup(api, "/article")

	articleGroup.UseSimpleModifier(func(op *huma.Operation) {
		op.Tags = []string{"文章"}
	})
	huma.Get(articleGroup, "", articleHandler.List)

	// 2. 文章分类（ArticleType / Category）相关的路由组
	typeGroup := huma.NewGroup(api, "/article-types") // 或者 /categories
	typeGroup.UseSimpleModifier(func(op *huma.Operation) {
		op.Tags = []string{"文章"}
	})
	huma.Get(typeGroup, "", articleTypeHandler.List)

	//huma.Get(typeGroup, "", h.ListTypes)
	//huma.Get(typeGroup, "/{id}", h.GetType)
	//huma.Post(typeGroup, "", h.CreateType)
	//huma.Put(typeGroup, "/{id}", h.UpdateType)
	//huma.Delete(typeGroup, "/{id}", h.DeleteType)
}
