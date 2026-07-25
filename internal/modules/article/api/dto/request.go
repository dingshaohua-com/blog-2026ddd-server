package dto

import "time"

type ArticleListItem struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	TypeID      string    `json:"typeId"`
	CreatedAt   time.Time `json:"createdAt"`
}

type ArticleTypeListItem struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}
