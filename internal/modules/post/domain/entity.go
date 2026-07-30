package domain

import (
	"errors"
	"time"
)

// 领域错误
var (
	ErrPostNotFound       = errors.New("post 不存在")
	ErrPostContentEmpty   = errors.New("post 内容不能为空")
	ErrPostContentTooLong = errors.New("post 内容太长")
)

// Post 领域实体（也是业务的对象）
type Post struct {
	id        int
	content   PostContent
	createdAt time.Time
	updatedAt time.Time
}

func (p *Post) ID() int {
	return p.id
}
func (p *Post) Content() PostContent {
	return p.content
}
func (p *Post) CreatedAt() time.Time {
	return p.createdAt
}
func (p *Post) UpdatedAt() time.Time {
	return p.updatedAt
}

// NewPost 领域工厂方法
func NewPost(content PostContent) (*Post, error) {
	post := &Post{
		content: content,
	}
	return post, nil
}

// ChangeContent 领域行为，用于修改文章内容
func (p *Post) ChangeContent(content string) error {
	postContent, err := NewPostContent(content)
	if err != nil {
		return err
	}
	p.content = postContent
	return nil
}

// RestorePost 业务的（领域）对象重建方法
func RestorePost(id int, content PostContent, createdAt time.Time, updatedAt time.Time) *Post {
	return &Post{
		id:        id,
		content:   content,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}
