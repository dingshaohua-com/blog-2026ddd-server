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

func (s *ArticleTypeService) List(ctx context.Context) ([]*ArticleTypeListItem, error) {
	articleTypes, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}

	items := make([]*ArticleTypeListItem, 0, len(articleTypes))
	for _, articleType := range articleTypes {
		items = append(items, &ArticleTypeListItem{
			ID: articleType.ID, Name: articleType.Name, Slug: articleType.Slug,
		})
	}
	return items, nil
}
