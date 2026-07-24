package domain

import "context"

type PostRepository interface {
	Create(ctx context.Context, post *Post) (*Post, error)
	FindByID(ctx context.Context, id int) (*Post, error)
	List(ctx context.Context) ([]*Post, error)
	Update(ctx context.Context, post *Post) error
	Delete(ctx context.Context, id int) error
}
