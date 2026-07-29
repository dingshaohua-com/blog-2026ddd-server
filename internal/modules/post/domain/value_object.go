package domain

import (
	"strings"
	"unicode/utf8"
)

// PostContent 值对象
type PostContent struct {
	value string
}

const MaxPostContentLength = 300

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
