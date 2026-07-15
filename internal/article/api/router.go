package api

import (
	"github.com/danielgtaylor/huma/v2"
)

func RegisterRoutes(h *Handler, api huma.API) {
	//group := engine.Group("/article")
	//huma.Get(api, "", h.List))
	//group.GET("", h.List)
	//r.GET("/:id", h.Detail)
	//r.POST("", h.Create)
	//r.PUT("/:id", h.Update)
	//r.DELETE("/:id", h.Delete)

	group := huma.NewGroup(api, "/article")
	huma.Get(group, "", h.List)
}
