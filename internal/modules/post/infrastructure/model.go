package infrastructure

import (
	"blog-2026ddd-server/internal/modules/post/domain"
	"time"
)

// PostModel 才是 PO，domain.Post 不是 PO，而是领域实体（Domain Entity）。
type PostModel struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement"`
	Content   string    `gorm:"column:content"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (PostModel) TableName() string { return "post" }

func (m PostModel) toDomain() (*domain.Post, error) {
	content, err := domain.NewPostContent(m.Content)
	if err != nil {
		return nil, err
	}
	return domain.RestorePost(m.ID, content, m.CreatedAt, m.UpdatedAt), nil
}

// 封装在 PO 层（或专门的 convert 包里）
func toDomainList(models []PostModel) ([]*domain.Post, error) {
	res := make([]*domain.Post, 0, len(models))
	for _, m := range models {
		post, err := m.toDomain()
		if err != nil {
			return nil, err
		}
		res = append(res, post) // 调用单个转换
	}
	return res, nil
}
