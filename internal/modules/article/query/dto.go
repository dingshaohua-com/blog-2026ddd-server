package query

import (
	"time"
)

type ArticleListItem struct {
	ID          int
	Title       string
	Description string
	TypeID      int
	TypeName    string
	CreatedAt   time.Time
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
