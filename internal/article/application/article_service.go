package application

import (
	"blog-2026ddd-server/internal/article/domain"
	"context"
)

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
		articles = make([]*ListItem, 0)
	}
	return ListResult{
		Items: articles, Total: total, Page: query.Page, PageSize: query.PageSize,
	}, nil
}

func (s *ArticleService) GetByID(ctx context.Context, id uint) (*domain.Article, error) {
	return s.repo.GetByID(ctx, id)
}
