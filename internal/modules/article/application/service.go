package application

import (
	"blog-2026ddd-server/internal/modules/article/domain"
	"context"
)

// ArticleService ---start---
type ArticleService struct {
	repo  domain.ArticleRepository
	query ArticleQuery
}

func NewArticleService(repo domain.ArticleRepository, query ArticleQuery) *ArticleService {
	return &ArticleService{
		repo:  repo,
		query: query,
	}
}

func (s *ArticleService) List(
	ctx context.Context,
	query ListQuery,
) (ListResult, error) {
	query = query.Normalize()
	articles, total, err := s.query.List(ctx, query)
	if err != nil {
		return ListResult{}, err
	}
	if articles == nil {
		articles = make([]*ArticleListItem, 0)
	}
	return ListResult{
		Items: articles, Total: total, Page: query.Page, PageSize: query.PageSize,
	}, nil
}

func (s *ArticleService) GetByID(ctx context.Context, id int) (*domain.Article, error) {
	return s.repo.GetByID(ctx, id)
}

// ArticleTypeService ---start---
type ArticleTypeService struct {
	repo domain.ArticleTypeRepository
}

func NewArticleTypeService(repo domain.ArticleTypeRepository) *ArticleTypeService {
	return &ArticleTypeService{
		repo: repo,
	}
}

func (s *ArticleTypeService) List(ctx context.Context) ([]*domain.ArticleType, error) {
	return s.repo.List(ctx)
}
