package domain

import (
	"blog-2026ddd-server/internal/article/dto"
	"context"
)

type Repository interface {
	Create(ctx context.Context, article *Article) error
	Delete(ctx context.Context, id uint) error
	Update(ctx context.Context, article *Article) error
	List(ctx context.Context) ([]*dto.ArticleListItem, error)
	FindByID(ctx context.Context, id uint) (*Article, error)
}
