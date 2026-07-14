package articleinfra

import (
	"blog-2026ddd-server/internal/article/domain"
	"blog-2026ddd-server/internal/article/dto"
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

func (r *Repository) Create(
	ctx context.Context,
	article *domain.Article,
) error {
	return r.db.
		WithContext(ctx).
		Create(article).
		Error
}
func (r *Repository) Delete(ctx context.Context, id uint) error {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) Update(ctx context.Context, article *domain.Article) error {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) List(ctx context.Context) ([]*dto.ArticleListItem, error) {
	var articles []*dto.ArticleListItem
	//err := r.db.WithContext(ctx).Find(&articles).Error
	err := r.db.
		WithContext(ctx).
		Table("article").
		Scan(&articles).
		Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (r *Repository) FindByID(

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
