package api

import (
	"blog-2026ddd-server/internal/modules/post/api/dto"
	"blog-2026ddd-server/internal/modules/post/application"
	"blog-2026ddd-server/internal/modules/post/domain"
	sharedapi "blog-2026ddd-server/internal/shared/api"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type PostHandler struct {
	service *application.PostService
}

type GetPostRequest struct {
	ID int `path:"id" minimum:"1" doc:"文章 ID"`
}

func NewPostHandler(service *application.PostService) *PostHandler {
	return &PostHandler{
		service: service,
	}
}

func (h *PostHandler) List(ctx context.Context, _ *struct{}) (*sharedapi.BodyResponse[[]*dto.PostDTO], error) {
	result, err := h.service.List(ctx)
	if err != nil {
		log.Printf("list article types: %v", err)
		return nil, sharedapi.InternalError(err.Error())
	}
	post := dto.ToPostDTOList(result)
	fmt.Print(result, post)
	return sharedapi.NewSuccessResponse(post), nil
}

func (h *PostHandler) Get(
	ctx context.Context,
	req *GetPostRequest,
) (*sharedapi.BodyResponse[*dto.PostDTO], error) {
	post, err := h.service.Get(ctx, req.ID)
	if err != nil {
		if errors.Is(err, domain.ErrPostNotFound) {
			return nil, sharedapi.NewError(
				http.StatusNotFound,
				"文章不存在",
			)
		}

		log.Printf("get post: %v", err)
		return nil, sharedapi.InternalError("文章详情加载失败")
	}

	return sharedapi.NewSuccessResponse(dto.ToPostDTO(post)), nil
}
