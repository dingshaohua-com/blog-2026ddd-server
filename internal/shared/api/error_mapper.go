package api

import (
	"errors"
	"log"
)

type ErrorMapping struct {
	Target error
	Status int
}

// MapError 领域错误 → HTTP 状态码
func MapError(
	internalMessage string,
	err error,
	mappings ...ErrorMapping,
) error {
	for _, mapping := range mappings {
		if errors.Is(err, mapping.Target) {
			return NewError(
				mapping.Status,
				err.Error(),
			)
		}
	}

	log.Printf("operation failed: %v", err)
	return InternalError(internalMessage + err.Error())
}

// handlePostError 将领域错误转换为 HTTP 错误。
// Handler 不需要知道 GORM 等基础设施错误。
//func handlePostError(err error) error {
//	switch {
//	case errors.Is(err, domain.ErrPostNotFound):
//		return sharedApi.NewError(
//			http.StatusNotFound,
//			"文章不存在",
//		)
//
//	case errors.Is(err, domain.ErrPostContentEmpty):
//		return sharedApi.NewError(
//			http.StatusUnprocessableEntity,
//			"文章内容不能为空",
//		)
//
//	case errors.Is(err, domain.ErrPostContentTooLong):
//		return sharedApi.NewError(
//			http.StatusUnprocessableEntity,
//			"文章内容不能超过 100 个字符",
//		)
//
//	default:
//		log.Printf("post operation failed: %v", err)
//		return sharedApi.InternalError("文章操作失败")
//	}
//}

//if err != nil {
//	return nil, handlePostError(err)
//}
