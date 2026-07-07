package application

import (
	"blog-2026ddd-server/internal/article/domain"
	"blog-2026ddd-server/internal/article/dto"
	"context"
)

type Service struct {
	repo domain.Repository
}

func NewService(repo domain.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetArticles(
	ctx context.Context,
) (*domain.Article, error) {
	var count int64
	var results []dto.ArticleResponse

	s.repo.List()

}

func (s *Service) GetArticle(

	ctx context.Context,
	id uint,

) (*domain.Article, error) {

	return s.repo.FindByID(ctx, id)

}
