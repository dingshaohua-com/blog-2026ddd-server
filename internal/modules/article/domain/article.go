package domain

import "time"

type Article struct {
	ID          int
	Title       string
	Description string
	TypeID      string
	CreateTime  time.Time
	Content     string
}
