package infrastructure

import (
	"blog-2026ddd-server/internal/modules/post/domain"
	"time"
)

type PostModel struct {
	//ID         int64     `gorm:"column:id;primaryKey"`
	//Content    string    `gorm:"column:content"`
	//CreateTime time.Time `gorm:"column:create_time"`

	ID         int       `gorm:"column:id;primaryKey;autoIncrement"`
	Content    string    `gorm:"column:content"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime"`
}

func (PostModel) TableName() string { return "post" }

func (m PostModel) toDomain() *domain.Post {
	return &domain.Post{ID: m.ID, Content: m.Content, CreateTime: m.CreateTime}
}

// 封装在 PO 层（或专门的 convert 包里）
func toDomainList(models []PostModel) []*domain.Post {
	if len(models) == 0 {
		return nil
	}
	res := make([]*domain.Post, 0, len(models))
	for _, m := range models {
		res = append(res, m.toDomain()) // 调用单个转换
	}
	return res
}
