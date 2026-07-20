package articleinfra

import (
	"blog-2026ddd-server/internal/article/domain"
	"blog-2026ddd-server/internal/article/dto"
	"blog-2026ddd-server/internal/infrastructure"
	"blog-2026ddd-server/internal/shared/api"
	"context"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// func (r *Repository) Create(
//
//	ctx context.Context,
//	article *domain.Article,
//
//	) error {
//		return r.db.
//			WithContext(ctx).
//			Create(article).
//			Error
//	}
func (r *Repository) Delete(ctx context.Context, id uint) error {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) Update(ctx context.Context, article *domain.Article) error {
	//TODO implement me
	panic("implement me")
}

// func (r *Repository) ListArticles(ctx context.Context, param api.Page) ([]*dto.ArticleListItem, error) {
func (r *Repository) ListArticles(ctx context.Context, param api.Page) (api.PageResult[*dto.ArticleListItem], error) {

	//var articles []*dto.ArticleListItem
	//err := r.db.WithContext(ctx).Find(&articles).Error
	//err := r.db.
	//	WithContext(ctx).
	//	Table("article").
	//	Scan(&articles).
	//	Error
	//if err != nil {
	//	return nil, err
	//}
	//return articles, nil

	var articles []*dto.ArticleListItem
	var total int64

	// 1. 统计总数
	if err := r.db.Table("article").Count(&total).Error; err != nil {
		return api.PageResult[*dto.ArticleListItem]{}, err
	}

	// 2. 调用上面的 Paginate 关联查询
	if err := r.db.Table("article").Scopes(infrastructure.Paginate(&param)).Find(&articles).Error; err != nil {
		return api.PageResult[*dto.ArticleListItem]{}, err
	}

	// 3. 返回结构化结果
	return api.NewPageResult(articles, total, &param), nil
}

func (r *Repository) GetArticleByID(

	ctx context.Context,
	id uint,

) (*domain.Article, error) {

	var article domain.Article
	err := r.db.
		WithContext(ctx).
		First(&article, id).
		Error
	if err != nil {
		return nil, err
	}
	return &article, nil

}
