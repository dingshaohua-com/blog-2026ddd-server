package articleinfra

import (
	domain2 "blog-2026ddd-server/internal/modules/article/domain"
	"context"

	"gorm.io/gorm"
)

type ArticleTypeRepository struct {
	db *gorm.DB
}

func NewArticleTypeRepository(db *gorm.DB) *ArticleTypeRepository {
	return &ArticleTypeRepository{db: db}
}

func (r *ArticleTypeRepository) List(ctx context.Context) ([]*domain2.ArticleType, error) {
	var models []articleTypeModel
	if err := r.db.WithContext(ctx).Find(&models).Error; err != nil {
		return nil, err
	}

	articleTypes := make([]*domain2.ArticleType, 0, len(models))
	for _, model := range models {
		articleTypes = append(articleTypes, model.toDomain())
	}
	return articleTypes, nil
}

var _ domain2.ArticleTypeRepository = (*ArticleTypeRepository)(nil)
