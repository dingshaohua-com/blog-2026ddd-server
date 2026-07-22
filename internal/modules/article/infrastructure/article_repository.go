package articleinfra

import (
	"blog-2026ddd-server/internal/modules/article/application"
	domain2 "blog-2026ddd-server/internal/modules/article/domain"
	"context"

	"gorm.io/gorm"
)

type ArticleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) *ArticleRepository {
	return &ArticleRepository{db: db}
}

func (r *ArticleRepository) List(ctx context.Context, query application.ListQuery) ([]*application.ListItem, int64, error) {
	var articles []*application.ListItem
	var total int64

	dbQuery := r.db.WithContext(ctx).Table("article")
	if err := dbQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := dbQuery.Order("create_time DESC").
		Order("id DESC").Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize).Find(&articles).Error; err != nil {
		return nil, 0, err
	}
	return articles, total, nil
}

func (r *ArticleRepository) GetByID(ctx context.Context, id int) (*domain2.Article, error) {
	var model articleModel
	if err := r.db.WithContext(ctx).First(&model, id).Error; err != nil {
		return nil, err
	}
	return model.toDomain(), nil
}

var _ domain2.ArticleRepository = (*ArticleRepository)(nil)
var _ application.ArticleQuery = (*ArticleRepository)(nil)
