package dto

import (
	"blog-2026ddd-server/internal/modules/article/application"
	"time"
)

type ArticleListItemDTO struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	TypeID      int       `json:"typeId"`
	TypeName    string    `json:"typeName"` // 新增
	CreatedAt   time.Time `json:"createdAt"`
}

type ArticleTypeListItemDTO struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// ToArticleDTO 2. 封装单个对象的转换：PO -> DTO
func ToArticleDTO(article *application.ArticleListItem) *ArticleListItemDTO {
	if article == nil {
		return nil
	}
	return &ArticleListItemDTO{
		ID:          article.ID,
		Title:       article.Title,
		Description: article.Description,
		TypeID:      article.TypeID,
		TypeName:    article.TypeName,
		CreatedAt:   article.CreatedAt,
	}
}

// ToArticleDTOList 3. 封装切片/列表的批量转换：[]PO -> []DTO
func ToArticleDTOList(pos []*application.ArticleListItem) []*ArticleListItemDTO {
	list := make([]*ArticleListItemDTO, 0, len(pos))
	for _, po := range pos {
		list = append(list, ToArticleDTO(po))
	}
	return list
}
