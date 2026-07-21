package api

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type ErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`

	status int
}

func (e *ErrorResponse) Error() string {
	return e.Msg
}

func (e *ErrorResponse) GetStatus() int {
	return e.status
}

func NewError(status int, msg string) *ErrorResponse {
	return &ErrorResponse{
		Code:   CodeFailure,
		Msg:    msg,
		Data:   nil,
		status: status,
	}
}

func NewCodeError(status, code int, msg string) *ErrorResponse {
	return &ErrorResponse{
		Code:   code,
		Msg:    msg,
		Data:   nil,
		status: status,
	}
}

func InternalError(messages ...string) *ErrorResponse {
	msg := "服务器内部错误"
	if len(messages) > 0 && messages[0] != "" {
		msg += "：" + messages[0]
	}

	return NewError(http.StatusInternalServerError, msg)
}

func ConfigureHumaErrors() {
	huma.NewError = func(status int, msg string, _ ...error) huma.StatusError {
		if status >= http.StatusInternalServerError {
			return InternalError()
		}
		return NewError(status, normalizeErrorMessage(status, msg))
	}

	huma.NewErrorWithContext = func(
		_ huma.Context,
		status int,
		msg string,
		errs ...error,
	) huma.StatusError {
		return huma.NewError(status, msg, errs...)
	}
}

func normalizeErrorMessage(status int, msg string) string {
	switch status {
	case http.StatusBadRequest:
		return "请求参数错误"
	case http.StatusUnauthorized:
		return "未登录或登录已失效"
	case http.StatusForbidden:
		return "无权执行该操作"
	case http.StatusNotFound:
		return "请求的资源不存在"
	case http.StatusUnprocessableEntity:
		return "参数校验失败"
	default:
		if msg != "" {
			return msg
		}
		return http.StatusText(status)
	}
}
