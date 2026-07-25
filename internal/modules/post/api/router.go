package api

import (
	"github.com/danielgtaylor/huma/v2"
)

func RegisterRoutes(handler *PostHandler, api huma.API) {
	postGroup := huma.NewGroup(api, "/post")
	postGroup.UseSimpleModifier(func(op *huma.Operation) {
		op.Tags = []string{"说说"}
	})
	huma.Get(postGroup, "", handler.List)
	huma.Get(postGroup, "/{id}", handler.Get)
	huma.Post(postGroup, "", handler.Create)
	huma.Put(postGroup, "/{id}", handler.Update)
	huma.Delete(postGroup, "/{id}", handler.Delete)
}
