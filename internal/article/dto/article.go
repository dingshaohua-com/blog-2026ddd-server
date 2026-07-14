package dto

import "time"

type ArticleListItem struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	TypeID      string    `json:"typeId"`
	CreateTime  time.Time `json:"createTime"`
}
