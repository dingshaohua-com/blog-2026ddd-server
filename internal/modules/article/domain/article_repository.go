package domain

import "context"

type ArticleRepository interface {
	GetByID(ctx context.Context, id int) (*Article, error)
}
