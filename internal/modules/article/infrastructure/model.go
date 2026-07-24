package articleinfra

import (
	domain2 "blog-2026ddd-server/internal/modules/article/domain"
	"time"
)

type articleModel struct {
	ID          int       `gorm:"column:id;primaryKey"`
	Title       string    `gorm:"column:title"`
	Description string    `gorm:"column:description"`
	TypeID      string    `gorm:"column:type_id"`
	CreatedAt   time.Time `gorm:"column:create_time"`
	Content     string    `gorm:"column:content"`
}

func (articleModel) TableName() string { return "article" }

func (m articleModel) toDomain() *domain2.Article {
	return &domain2.Article{
		ID: m.ID, Title: m.Title, Description: m.Description,
		TypeID: m.TypeID, CreatedAt: m.CreatedAt, Content: m.Content,
	}
}

type articleTypeModel struct {
	ID   uint   `gorm:"column:id;primaryKey"`
	Name string `gorm:"column:name;type:varchar(50);not null"`
	Slug string `gorm:"column:slug;type:varchar(50);uniqueIndex;not null"`
}

func (articleTypeModel) TableName() string { return "article_type" }

func (m articleTypeModel) toDomain() *domain2.ArticleType {
	return &domain2.ArticleType{ID: m.ID, Name: m.Name, Slug: m.Slug}
}
