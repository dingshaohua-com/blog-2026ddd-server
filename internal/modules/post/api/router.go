package api

import (
	"github.com/danielgtaylor/huma/v2"
)

func RegisterRoutes(handler *PostHandler, api huma.API) {
	postGroup := huma.NewGroup(api, "/post")
	huma.Get(postGroup, "", handler.List)
}
