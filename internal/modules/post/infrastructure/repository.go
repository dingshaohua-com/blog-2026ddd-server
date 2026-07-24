package infrastructure

import (
	"blog-2026ddd-server/internal/modules/post/domain"
	"context"
	"errors"

	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func (r *PostRepository) Create(ctx context.Context, post *domain.Post) (*domain.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (r *PostRepository) Update(ctx context.Context, post *domain.Post) error {
	//TODO implement me
	panic("implement me")
}

func (r *PostRepository) Delete(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) List(ctx context.Context) ([]*domain.Post, error) {
	// 1. 数据库持久层（DAO / ORM 层）：查出来的是和数据库表字段一一对应的模型
	var models []PostModel
	if err := r.db.WithContext(ctx).Find(&models).Error; err != nil {
		return nil, err
	}
	// 2. 赋值转化,把数据库模型（PO）转换为业务模型（Domain）
	posts := toDomainList(models)
	// 3. 把转换好的 posts 返回给上层（Service / Biz 层）使用
	return posts, nil

}

func (r *PostRepository) FindByID(
	ctx context.Context,
	id int,
) (*domain.Post, error) {
	var model PostModel

	err := r.db.
		WithContext(ctx).
		Where("id = ?", id).
		First(&model).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, domain.ErrPostNotFound
	}

	if err != nil {
		return nil, err
	}

	return model.toDomain(), nil
}
