package api

import "github.com/gin-gonic/gin"

func RegisterRoutes(engine *gin.Engine, h *Handler) {
	group := engine.Group("/article")
	group.GET("", h.List)
	//r.GET("/:id", h.Detail)
	//r.POST("", h.Create)
	//r.PUT("/:id", h.Update)
	//r.DELETE("/:id", h.Delete)

}
