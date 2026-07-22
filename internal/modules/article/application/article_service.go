package application

import (
	domain2 "blog-2026ddd-server/internal/modules/article/domain"
	"context"
)

type ArticleService struct {
	repo  domain2.ArticleRepository
	query ArticleQuery
}

func NewArticleService(repo domain2.ArticleRepository, query ArticleQuery) *ArticleService {
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

func (s *ArticleService) GetByID(ctx context.Context, id int) (*domain2.Article, error) {
	return s.repo.GetByID(ctx, id)
}
