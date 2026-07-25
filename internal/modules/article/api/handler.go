package api

import (
	"blog-2026ddd-server/internal/modules/article/api/dto"
	"blog-2026ddd-server/internal/modules/article/application"
	"blog-2026ddd-server/internal/shared/api"
	"context"
	"log"
)

// ArticleHandler ---start--
type ArticleHandler struct {
	service *application.ArticleService
}

func NewArticleHandler(service *application.ArticleService) *ArticleHandler {
	return &ArticleHandler{
		service: service,
	}
}

type ListRequest struct {
	api.Page
}

func (h *ArticleHandler) List(ctx context.Context, req *ListRequest) (*api.PageBodyResponse[*dto.ArticleListItemDTO], error) {
	result, err := h.service.List(ctx, application.ListQuery{
		Page: req.Page.Page, PageSize: req.Page.PageSize,
	})
	if err != nil {
		log.Printf("list articles: %v", err)
		return nil, api.InternalError("文章列表加载失败")
	}
	items := dto.ToArticleDTOList(result.Items)
	page := api.Page{Page: result.Page, PageSize: result.PageSize}
	return api.NewSuccessResponse(api.NewPageResult(items, result.Total, &page)), nil
}

// ArticleTypeHandler ---start---
type ArticleTypeHandler struct {
	service *application.ArticleTypeService
}

func NewArticleTypeHandler(service *application.ArticleTypeService) *ArticleTypeHandler {
	return &ArticleTypeHandler{
		service: service,
	}
}

func (h *ArticleTypeHandler) List(ctx context.Context, _ *struct{}) (*api.BodyResponse[[]*dto.ArticleTypeListItemDTO], error) {
	articleTypes, err := h.service.List(ctx)
	if err != nil {
		log.Printf("list article types: %v", err)
		return nil, api.InternalError(err.Error())
	}
	items := make([]*dto.ArticleTypeListItemDTO, 0, len(articleTypes))
	for _, articleType := range articleTypes {
		items = append(items, &dto.ArticleTypeListItemDTO{
			ID: articleType.ID, Name: articleType.Name, Slug: articleType.Slug,
		})
	}
	return api.NewSuccessResponse(items), nil
}
