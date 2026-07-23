package dto

import (
	"blog-2026ddd-server/internal/modules/post/domain"
	"time"
)

// 1. 封装结构体
type PostDTO struct {
	ID         int       `json:"id"`
	Content    string    `json:"content"`
	CreateTime time.Time `json:"createTime"`
}

// 2. 封装单个对象的转换：PO -> DTO
func toPostDTO(po *domain.Post) *PostDTO {
	if po == nil {
		return nil
	}
	return &PostDTO{
		ID:         po.ID,
		Content:    po.Content,
		CreateTime: po.CreateTime,
	}
}

// 3. 封装切片/列表的批量转换：[]PO -> []DTO
func ToPostDTOList(pos []*domain.Post) []*PostDTO {
	list := make([]*PostDTO, 0, len(pos))
	for _, po := range pos {
		list = append(list, toPostDTO(po))
	}
	return list
}
