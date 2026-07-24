package domain

import (
	"errors"
	"strings"
	"time"
	"unicode/utf8"
)

const MaxPostContentLength = 100

var (
	ErrPostNotFound       = errors.New("post 不存在")
	ErrPostContentEmpty   = errors.New("post 内容不能为空")
	ErrPostContentTooLong = errors.New("post 内容太长")
)

type Post struct {
	id        int
	content   string
	createdAt time.Time
	updatedAt time.Time
}

func (p *Post) ID() int {
	return p.id
}

func (p *Post) Content() string {
	return p.content
}
func (p *Post) CreatedAt() time.Time {
	return p.createdAt
}
func (p *Post) UpdatedAt() time.Time {
	return p.updatedAt
}

func normalizeContent(content string) (string, error) {
	content = strings.TrimSpace(content)
	switch {
	case content == "":
		return "", ErrPostContentEmpty
	case utf8.RuneCountInString(content) > MaxPostContentLength:
		return "", ErrPostContentTooLong
	default:
		return content, nil
	}
}

// ChangeContent 修改文章内容。
func (p *Post) ChangeContent(content string) error {
	content, err := normalizeContent(content)
	if err != nil {
		return err
	}
	p.content = content
	return nil
}

// NewPost 构造函数，用于创建一篇新文章。
func NewPost(content string, now time.Time) (*Post, error) {
	post := &Post{
		createdAt: now,
	}

	if err := post.ChangeContent(content); err != nil {
		return nil, err
	}

	return post, nil
}

func RestorePost(
	id int,
	content string,
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
