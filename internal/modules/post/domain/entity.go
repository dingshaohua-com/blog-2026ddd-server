package domain

import (
	"errors"
	"strings"
	"time"
	"unicode/utf8"
)

const MaxPostContentLength = 300

// 领域错误
var (
	ErrPostNotFound       = errors.New("post 不存在")
	ErrPostContentEmpty   = errors.New("post 内容不能为空")
	ErrPostContentTooLong = errors.New("post 内容太长")
)

// 业务的对象/领域实体
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

// 值对象
type PostContent struct {
	value string
}

func NewPostContent(value string) (PostContent, error) {
	value = strings.TrimSpace(value)
	switch {
	case value == "":
		return PostContent{}, ErrPostContentEmpty
	case utf8.RuneCountInString(value) > MaxPostContentLength:
		return PostContent{}, ErrPostContentTooLong
	}
	return PostContent{value: value}, nil
}
func (c PostContent) String() string {
	return c.value
}

// 领域行为：修改文章内容
func (p *Post) ChangeContent(content string) error {
	postContent, err := NewPostContent(content)
	if err != nil {
		return err
	}
	p.content = postContent
	return nil
}

// 领域行为：创建领域对象的工厂方法，用于创建一篇新文章，（没错，领域实体的构造函数本身属于一个领域行为）
func NewPost(content PostContent, now time.Time) (*Post, error) {
	post := &Post{
		createdAt: now,
		content:   content,
	}

	return post, nil
}

// 业务的（领域）对象重建方法（不是业务行为）：用于 Infrastructure 从持久化数据恢复领域对象，为了让 Repository 在查询数据库后，能够构造私有字段的领域对象
func RestorePost(
	id int,
	content PostContent,
	createdAt time.Time,
	updatedAt time.Time,
) *Post {
	return &Post{
		id:        id,
		content:   content,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}
