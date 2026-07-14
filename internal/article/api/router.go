package api

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.Engine, h *Handler, api huma.API) {
	group := engine.Group("/article")
	//huma.Get(api, "", h.List))
	group.GET("", h.List)
	//r.GET("/:id", h.Detail)
	//r.POST("", h.Create)
	//r.PUT("/:id", h.Update)
	//r.DELETE("/:id", h.Delete)

}
