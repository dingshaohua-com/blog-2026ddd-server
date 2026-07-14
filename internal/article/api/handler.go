package api

import (
	"blog-2026ddd-server/internal/article/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *application.Service
}

func NewHandler(service *application.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) List(c *gin.Context) {
	articles, _ := h.service.GetArticles(c)
	c.JSON(http.StatusOK, articles)
}
