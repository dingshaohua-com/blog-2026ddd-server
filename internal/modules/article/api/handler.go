package api

import (
	"blog-2026ddd-server/internal/modules/article/api/dto"
	"blog-2026ddd-server/internal/modules/article/application"
	"blog-2026ddd-server/internal/modules/article/query"
	"blog-2026ddd-server/internal/shared/api"
	"context"
)

// ArticleHandler ---start--
type ArticleHandler struct {
	service *application.ArticleService
	query   *query.ArticleQuery
}

func NewArticleHandler(service *application.ArticleService, query *query.ArticleQuery) *ArticleHandler {
	return &ArticleHandler{
		service: service,
		query:   query,
	}
}

type ListRequest struct {
	api.Page
}

func (h *ArticleHandler) List(ctx context.Context, req *ListRequest) (*api.Body[api.PageResult[*dto.ArticleListItemDTO]], error) {
	result, err := h.query.List(ctx, query.ListQuery{
		Page: req.Page.Page, PageSize: req.Page.PageSize,
	})
	if err != nil {
		return nil, api.InternalError(err.Error())
	}
	items := dto.ToArticleDTOList(result.Items)
	page := api.Page{Page: result.Page, PageSize: result.PageSize}
	return api.NewBody(api.NewPageResult(items, result.Total, &page)), nil
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

func (h *ArticleTypeHandler) List(ctx context.Context, _ *struct{}) (*api.Body[[]*dto.ArticleTypeListItemDTO], error) {
	articleTypes, err := h.service.List(ctx)
	if err != nil {
		return nil, api.InternalError(err.Error())
	}
	items := make([]*dto.ArticleTypeListItemDTO, 0, len(articleTypes))
	for _, articleType := range articleTypes {
		items = append(items, &dto.ArticleTypeListItemDTO{
			ID: articleType.ID, Name: articleType.Name, Slug: articleType.Slug,
		})
	}
	return api.NewBody(items), nil
}
