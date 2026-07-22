package application

import (
	"context"
	"time"
)

type ArticleListItem struct {
	ID          int
	Title       string
	Description string
	TypeID      string
	CreateTime  time.Time
}

type ListQuery struct {
	Page     int
	PageSize int
}

func (q ListQuery) Normalize() ListQuery {
	if q.Page <= 0 {
		q.Page = 1
	}
	if q.PageSize <= 0 {
		q.PageSize = 10
	} else if q.PageSize > 100 {
		q.PageSize = 100
	}
	return q
}

type ListResult struct {
	Items    []*ArticleListItem
	Total    int64
	Page     int
	PageSize int
}

type ArticleQuery interface {
	List(ctx context.Context, query ListQuery) ([]*ArticleListItem, int64, error)
}
