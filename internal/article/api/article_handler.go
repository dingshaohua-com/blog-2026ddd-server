package api

import (
	"blog-2026ddd-server/internal/article/api/dto"
	"blog-2026ddd-server/internal/article/application"
	"blog-2026ddd-server/internal/shared/api"
	"context"
	"log"
)

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

func (h *ArticleHandler) List(ctx context.Context, req *ListRequest) (*api.PageBodyResponse[*dto.ArticleListItem], error) {
	result, err := h.service.List(ctx, application.ListQuery{
		Page: req.Page.Page, PageSize: req.Page.PageSize,
	})
	if err != nil {
		log.Printf("list articles: %v", err)
		return nil, api.InternalError("文章列表加载失败")
	}

	items := make([]*dto.ArticleListItem, 0, len(result.Items))
	for _, article := range result.Items {
		items = append(items, &dto.ArticleListItem{
			ID:          article.ID,
			Title:       article.Title,
			Description: article.Description,
			TypeID:      article.TypeID,
			CreateTime:  article.CreateTime,
		})
	}
	page := api.Page{Page: result.Page, PageSize: result.PageSize}
	return api.NewSuccessResponse(api.NewPageResult(items, result.Total, &page)), nil
}
