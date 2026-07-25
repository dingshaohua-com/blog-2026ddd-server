package infrastructure

import (
	"blog-2026ddd-server/internal/modules/article/domain"
	"time"
)

type ArticleModel struct {
	ID          int       `gorm:"column:id;primaryKey"`
	Title       string    `gorm:"column:title"`
	Description string    `gorm:"column:description"`
	TypeID      int       `gorm:"column:type_id"`
	TypeName    string    `gorm:"column:type_name"` // 新增（连表查询用到）
	CreatedAt   time.Time `gorm:"column:created_at"`
	Content     string    `gorm:"column:content"`
}

func (ArticleModel) TableName() string { return "article" }

func (m ArticleModel) toDomain() *domain.Article {
	return &domain.Article{
		ID: m.ID, Title: m.Title, Description: m.Description,
		TypeID: m.TypeID, CreatedAt: m.CreatedAt, Content: m.Content,
	}
}

type articleTypeModel struct {
	ID   int    `gorm:"column:id;primaryKey"`
	Name string `gorm:"column:name;type:varchar(50);not null"`
	Slug string `gorm:"column:slug;type:varchar(50);uniqueIndex;not null"`
}

func (articleTypeModel) TableName() string { return "article_type" }

func (m articleTypeModel) toDomain() *domain.ArticleType {
	return &domain.ArticleType{ID: m.ID, Name: m.Name, Slug: m.Slug}
}
