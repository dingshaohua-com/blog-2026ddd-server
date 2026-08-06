package dto

import (
	"blog-2026ddd-server/internal/modules/article/query"
	"time"
)

type ArticleListItemDTO struct {
	ID          int       `json:"id" doc:"ID"`
	Title       string    `json:"title" doc:"标题"`
	Description string    `json:"description" doc:"简介描述"`
	TypeID      int       `json:"typeId" doc:"文章类型ID"`
	TypeName    string    `json:"typeName" doc:"文章类型名称"` // 新增
	CreatedAt   time.Time `json:"createdAt" doc:"创建时间"`
}

type ArticleTypeListItemDTO struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// ToArticleDTO 2. 封装单个对象的转换：PO -> DTO
func ToArticleDTO(article *query.ArticleListItem) ArticleListItemDTO {
	//if article == nil {
	//	return nil
	//}
	return ArticleListItemDTO{
		ID:          article.ID,
		Title:       article.Title,
		Description: article.Description,
		TypeID:      article.TypeID,
		TypeName:    article.TypeName,
		CreatedAt:   article.CreatedAt,
	}
}

// ToArticleDTOList 3. 封装切片/列表的批量转换：[]PO -> []DTO
func ToArticleDTOList(pos []*query.ArticleListItem) []ArticleListItemDTO {
	list := make([]ArticleListItemDTO, 0, len(pos))
	for _, po := range pos {
		list = append(list, ToArticleDTO(po))
	}
	return list
}
