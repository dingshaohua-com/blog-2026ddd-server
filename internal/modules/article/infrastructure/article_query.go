package articleinfra

import (
	"blog-2026ddd-server/internal/modules/article/application"
	"context"

	"gorm.io/gorm"
)

type ArticleQuery struct {
	db *gorm.DB
}

func NewArticleQuery(db *gorm.DB) *ArticleQuery {
	return &ArticleQuery{db: db}
}

func (r *ArticleQuery) List(ctx context.Context, query application.ListQuery) ([]*application.ArticleListItem, int64, error) {
	var articles []*application.ArticleListItem
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

var _ application.ArticleQuery = (*ArticleQuery)(nil)
