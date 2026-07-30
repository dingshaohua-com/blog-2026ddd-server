package api

import (
	"github.com/danielgtaylor/huma/v2"
)

func RegisterRoutes(articleHandler *ArticleHandler, articleTypeHandler *ArticleTypeHandler, api huma.API) {

	// 1. 文章相关的路由组
	articleGroup := huma.NewGroup(api, "/article")

	articleGroup.UseSimpleModifier(func(op *huma.Operation) {
		op.Tags = []string{"文章"}
	})
	huma.Get(articleGroup, "", articleHandler.List)

	// 2. 文章分类 相关的路由组
	articleTypeGroup := huma.NewGroup(api, "/article-types") // 或者 /categories
	articleTypeGroup.UseSimpleModifier(func(op *huma.Operation) {
		op.Tags = []string{"文章"}
	})
	huma.Get(articleTypeGroup, "", articleTypeHandler.List)
}
