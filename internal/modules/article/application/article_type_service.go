package application

import (
	"blog-2026ddd-server/internal/modules/article/domain"
	"context"
)

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
