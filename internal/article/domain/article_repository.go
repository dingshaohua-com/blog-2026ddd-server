package domain

import "context"

type ArticleRepository interface {
	GetByID(ctx context.Context, id uint) (*Article, error)
}
