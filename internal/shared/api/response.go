package api

type BodyResponse[T any] struct {
	Body T
}

func NewBodyResponse[T any](body T) *BodyResponse[T] {
	return &BodyResponse[T]{
		Body: body,
	}
}
