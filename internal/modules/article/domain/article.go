package domain

import "time"

type Article struct {
	ID          int
	Title       string
	Description string
	TypeID      string
	CreatedAt   time.Time
	Content     string
}
