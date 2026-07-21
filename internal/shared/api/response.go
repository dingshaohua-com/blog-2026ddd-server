package api

const (
	CodeSuccess = 0
	CodeFailure = -1
)

type Response[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

type BodyResponse[T any] struct {
	Body Response[T]
}

type PageBodyResponse[T any] = BodyResponse[PageResult[T]]

func NewSuccessResponse[T any](data T) *BodyResponse[T] {
	return &BodyResponse[T]{
		Body: Response[T]{
			Code: CodeSuccess,
			Msg:  "success",
			Data: data,
		},
	}
}
