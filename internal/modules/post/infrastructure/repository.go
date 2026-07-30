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

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) Create(ctx context.Context, post *domain.Post) (*domain.Post, error) {
	model := PostModel{
		Content:   post.Content().String(),
		CreatedAt: post.CreatedAt(),
	}
	if err := r.db.WithContext(ctx).Create(&model).Error; err != nil {
		return nil, err
	}
	return model.toDomain()
}

func (r *PostRepository) Update(ctx context.Context, post *domain.Post) error {
	result := r.db.
		WithContext(ctx).
		Model(&PostModel{}).
		Where("id = ?", post.ID()).
		Update("content", post.Content().String())
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domain.ErrPostNotFound
	}
	return nil
}

func (r *PostRepository) Delete(ctx context.Context, id int) error {
	result := r.db.
		WithContext(ctx).
		Delete(&PostModel{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domain.ErrPostNotFound
	}
	return nil
}

func (r *PostRepository) List(ctx context.Context) ([]*domain.Post, error) {
	var models []PostModel
	if err := r.db.WithContext(ctx).Find(&models).Error; err != nil {
		return nil, err
	}
	return toDomainList(models)
}

func (r *PostRepository) FindByID(ctx context.Context, id int) (*domain.Post, error) {
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
	return model.toDomain()
}
