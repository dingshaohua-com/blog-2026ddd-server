package query

import (
	"context"

	"gorm.io/gorm"
)

// ArticleQuery ---start---
type ArticleService struct {
	db *gorm.DB
}

func NewArticleService(db *gorm.DB) *ArticleService {
	return &ArticleService{db: db}
}

//func (r *ArticleQuery) List(ctx context.Context, query application.ListQuery) ([]*application.ArticleListItem, int64, error) {
//	var articles []*application.ArticleListItem
//	var total int64
//
//	dbQuery := r.db.WithContext(ctx).Table("article")
//	if err := dbQuery.Count(&total).Error; err != nil {
//		return nil, 0, err
//	}
//	if err := dbQuery.Order("created_at DESC").
//		Order("id DESC").Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize).Find(&articles).Error; err != nil {
//		return nil, 0, err
//	}
//	return articles, total, nil
//}

func (r *ArticleService) ListDao(ctx context.Context, query ListQuery) ([]*ArticleListItem, int64, error) {
	var articles []*ArticleListItem
	var total int64

	dbQuery := r.db.WithContext(ctx).Table("article")
	if err := dbQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 改成 JOIN 查询，SELECT 明确指定字段，避免 id 等字段冲突
	if err := dbQuery.
		Select("article.id, article.title, article.description, article.type_id, article_type.name as type_name, article.created_at").
		Joins("LEFT JOIN article_type ON article.type_id = article_type.id").
		Order("article.created_at DESC").
		Order("article.id DESC").
		Limit(query.PageSize).
		Offset((query.Page - 1) * query.PageSize).
		Find(&articles).Error; err != nil {
		return nil, 0, err
	}
	return articles, total, nil
}

func (s *ArticleService) List(
	ctx context.Context,
	query ListQuery,
) (ListResult, error) {
	query = query.Normalize()
	articles, total, err := s.ListDao(ctx, query)
	if err != nil {
		return ListResult{}, err
	}
	if articles == nil {
		articles = make([]*ArticleListItem, 0)
	}
	return ListResult{
		Items: articles, Total: total, Page: query.Page, PageSize: query.PageSize,
	}, nil
}
