package domain

import "context"

type ArticleTypeRepository interface {
	List(ctx context.Context) ([]*ArticleType, error)
}
