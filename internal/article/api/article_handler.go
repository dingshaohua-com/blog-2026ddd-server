package api

import (
	"blog-2026ddd-server/internal/article/application"
	"blog-2026ddd-server/internal/article/dto"
	"blog-2026ddd-server/internal/shared/api"
	"context"
	"fmt"
)

type ArticleHandler struct {
	service *application.Service
}

func NewArticleHandler(service *application.Service) *ArticleHandler {
	return &ArticleHandler{
		service: service,
	}
}

type ListRequest struct {
	api.Page
}

// func (h *Handler) List(ctx context.Context, req *ListRequest) (*api.BodyResponse[[]*dto.ArticleListItem], error) {
func (h *ArticleHandler) List(ctx context.Context, req *ListRequest) (*api.BodyResponse[api.PageResult[*dto.ArticleListItem]], error) {
	fmt.Print("req.param: ", req.Page)
	articles, err := h.service.GetArticles(ctx, req.Page)
	if err != nil {
		return nil, err
	}
	return api.NewBodyResponse(articles), nil
}
