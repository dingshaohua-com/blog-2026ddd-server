package application

import (
	"blog-2026ddd-server/internal/modules/article/domain"
	"context"
)

// ArticleService ---start---
type ArticleService struct {
	repo domain.ArticleRepository
}

func NewArticleService(repo domain.ArticleRepository) *ArticleService {
	return &ArticleService{
		repo: repo,
	}
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
