package articleinfra

import (
	"blog-2026ddd-server/internal/modules/article/domain"
	"context"

	"gorm.io/gorm"
)

type ArticleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) *ArticleRepository {
	return &ArticleRepository{db: db}
}

func (r *ArticleRepository) GetByID(ctx context.Context, id int) (*domain.Article, error) {
	var model articleModel
	if err := r.db.WithContext(ctx).First(&model, id).Error; err != nil {
		return nil, err
	}
	return model.toDomain(), nil
}

var _ domain.ArticleRepository = (*ArticleRepository)(nil)
