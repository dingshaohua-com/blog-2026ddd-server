package api

import (
	"blog-2026ddd-server/internal/article/application"
	"blog-2026ddd-server/internal/article/dto"
	"blog-2026ddd-server/internal/shared/api"
	"context"
)

type Handler struct {
	service *application.Service
}

func NewHandler(service *application.Service) *Handler {
	return &Handler{
		service: service,
	}
}

type ListRequest struct{}

func (h *Handler) List(ctx context.Context, req *ListRequest) (*api.BodyResponse[[]*dto.ArticleListItem], error) {
	articles, err := h.service.GetArticles(ctx)
	if err != nil {
		return nil, err
	}
	return api.NewBodyResponse(articles), nil
}
