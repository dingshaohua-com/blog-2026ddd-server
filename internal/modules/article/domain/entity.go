package domain

import "time"

type Article struct {
	ID          int
	Title       string
	Description string
	TypeID      int
	CreatedAt   time.Time
	Content     string
}

type ArticleType struct {
	ID   int
	Name string
	Slug string
}
