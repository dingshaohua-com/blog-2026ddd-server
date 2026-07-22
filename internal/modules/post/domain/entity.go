package domain

import "time"

type Post struct {
	ID         int64
	Content    string
	CreateTime time.Time
}
