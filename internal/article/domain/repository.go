package domain

import "context"

type Repository interface {
	Create(ctx context.Context, article *Article) error
	Update(ctx context.Context, article *Article) error
	Delete(ctx context.Context, id uint) error
	FindByID(ctx context.Context, id uint) (*Article, error)
	List(ctx context.Context) ([]*Article, error)
}
