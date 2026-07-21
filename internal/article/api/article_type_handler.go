package api

import (
	"blog-2026ddd-server/internal/article/api/dto"
	"blog-2026ddd-server/internal/article/application"
	"blog-2026ddd-server/internal/shared/api"
	"context"
	"log"
)

type ArticleTypeHandler struct {
	service *application.ArticleTypeService
}

func NewArticleTypeHandler(service *application.ArticleTypeService) *ArticleTypeHandler {
	return &ArticleTypeHandler{
		service: service,
	}
}

func (h *ArticleTypeHandler) List(ctx context.Context, _ *struct{}) (*api.BodyResponse[[]*dto.ArticleTypeListItem], error) {
	articleTypes, err := h.service.List(ctx)
	if err != nil {
		log.Printf("list article types: %v", err)
		return nil, api.InternalError(err.Error())
	}
	items := make([]*dto.ArticleTypeListItem, 0, len(articleTypes))
	for _, articleType := range articleTypes {
		items = append(items, &dto.ArticleTypeListItem{
			ID: articleType.ID, Name: articleType.Name, Slug: articleType.Slug,
		})
	}
	return api.NewSuccessResponse(items), nil
}
