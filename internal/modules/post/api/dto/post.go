package dto

import (
	"blog-2026ddd-server/internal/modules/post/domain"
	"time"
)

// 1. 封装结构体
type PostDTO struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// 2. 封装单个对象的转换：PO -> DTO
func ToPostDTO(post *domain.Post) *PostDTO {
	if post == nil {
		return nil
	}
	return &PostDTO{
		ID:        post.ID(),
		Content:   post.Content(),
		CreatedAt: post.CreatedAt(),
		UpdatedAt: post.UpdatedAt(),
	}
}

// 3. 封装切片/列表的批量转换：[]PO -> []DTO
func ToPostDTOList(pos []*domain.Post) []*PostDTO {
	list := make([]*PostDTO, 0, len(pos))
	for _, po := range pos {
		list = append(list, ToPostDTO(po))
	}
	return list
}
