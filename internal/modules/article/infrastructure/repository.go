package articleinfra

import (
	"blog-2026ddd-server/internal/modules/article/domain"
	"context"

	"gorm.io/gorm"
)

// ArticleRepository ---start---
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

// 显式检查-编译期接口实现检查：请 Go 编译器确认 *ArticleRepository 实现了 domain.ArticleRepository 接口
// 如果 repo 没实现接口，编译器同样会报错。
var _ domain.ArticleRepository = (*ArticleRepository)(nil)

// ArticleTypeRepository ---start---
type ArticleTypeRepository struct {
	db *gorm.DB
}

func NewArticleTypeRepository(db *gorm.DB) *ArticleTypeRepository {
	return &ArticleTypeRepository{db: db}
}

func (r *ArticleTypeRepository) List(ctx context.Context) ([]*domain.ArticleType, error) {
	var models []articleTypeModel
	if err := r.db.WithContext(ctx).Find(&models).Error; err != nil {
		return nil, err
	}

	articleTypes := make([]*domain.ArticleType, 0, len(models))
	for _, model := range models {
		articleTypes = append(articleTypes, model.toDomain())
	}
	return articleTypes, nil
}
