package api

import (
	"blog-2026ddd-server/internal/modules/post/api/dto"
	"blog-2026ddd-server/internal/modules/post/application"
	"blog-2026ddd-server/internal/shared/api"
	"context"
	"log"
)

type PostHandler struct {
	service *application.PostService
}

func NewPostHandler(service *application.PostService) *PostHandler {
	return &PostHandler{
		service: service,
	}
}

func (h *PostHandler) List(ctx context.Context, _ *struct{}) (*api.BodyResponse[[]*dto.PostDTO], error) {
	result, err := h.service.List(ctx)
	if err != nil {
		log.Printf("list article types: %v", err)
		return nil, api.InternalError(err.Error())
	}
	post := dto.ToPostDTOList(result)
	return api.NewSuccessResponse(post), nil
}
