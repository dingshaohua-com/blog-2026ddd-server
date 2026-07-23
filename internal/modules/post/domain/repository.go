package domain

import "context"

type PostRepository interface {
	List(ctx context.Context) ([]*Post, error)
}
