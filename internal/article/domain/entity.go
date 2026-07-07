package domain

import "time"

type Article struct {
	ID          int       `gorm:"column:id" json:"id"`
	Title       string    `gorm:"column:title" json:"title"`
	Description string    `gorm:"column:description" json:"description"`
	TypeID      string    `gorm:"column:type_id" json:"typeId"`
	CreateTime  time.Time `gorm:"column:create_time" json:"createTime"`
	Content     string    `gorm:"column:content" json:"content"`
}
