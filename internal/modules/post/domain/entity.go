package domain

import "time"

type Post struct {
	ID         int
	Content    string
	CreateTime time.Time
}
