package application

import (
	"blog-2026ddd-server/internal/article/domain"
	"blog-2026ddd-server/internal/article/dto"
	"blog-2026ddd-server/internal/shared/api"
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
	param api.Page,
) (api.PageResult[*dto.ArticleListItem], error) {
	articles, err := s.repo.ListArticles(ctx, param)
	return articles, err
}

func (s *Service) GetArticle(

	ctx context.Context,
	id uint,

) (*domain.Article, error) {

	return s.repo.GetArticleByID(ctx, id)

}
