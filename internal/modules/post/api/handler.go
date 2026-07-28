package api

import (
	"blog-2026ddd-server/internal/modules/post/api/dto"
	"blog-2026ddd-server/internal/modules/post/application"
	"blog-2026ddd-server/internal/modules/post/domain"
	sharedApi "blog-2026ddd-server/internal/shared/api"
	"context"
	"net/http"
)

type PostHandler struct {
	service *application.PostService
}

var errorMappings = []sharedApi.ErrorMapping{
	{
		Target: domain.ErrPostNotFound,
		Status: http.StatusNotFound,
	},
	{
		Target: domain.ErrPostContentEmpty,
		Status: http.StatusUnprocessableEntity,
	},
	{
		Target: domain.ErrPostContentTooLong,
		Status: http.StatusUnprocessableEntity,
	},
}

func NewPostHandler(service *application.PostService) *PostHandler {
	return &PostHandler{
		service: service,
	}
}

func (h *PostHandler) List(ctx context.Context, _ *struct{}) (*sharedApi.BodyResponse[[]*dto.PostDTO], error) {
	result, err := h.service.List(ctx)
	if err != nil {
		return nil, sharedApi.MapError(
			err,
			"获取失败",
			errorMappings...,
		)
	}
	post := dto.ToPostDTOList(result)
	return sharedApi.NewSuccessResponse(post), nil
}

func (h *PostHandler) Get(
	ctx context.Context,
	req *dto.GetPostRequest,
) (*sharedApi.BodyResponse[*dto.PostDTO], error) {
	post, err := h.service.Get(ctx, req.ID)
	if err != nil {
		return nil, sharedApi.MapError(
			err,
			"获取失败",
			errorMappings...,
		)
	}

	return sharedApi.NewSuccessResponse(dto.ToPostDTO(post)), nil
}

func (h *PostHandler) Create(ctx context.Context, req *dto.CreatePostRequest) (*sharedApi.BodyResponse[*dto.PostDTO], error) {
	post, err := h.service.Create(ctx, req.Body.Content)
	if err != nil {
		return nil, sharedApi.MapError(
			err,
			"插入失败",
			errorMappings...,
		)
	}
	return sharedApi.NewSuccessResponse(dto.ToPostDTO(post)), err
}

func (h *PostHandler) Update(ctx context.Context, req *dto.UpdatePostRequest) (*sharedApi.BodyResponse[any], error) {
	err := h.service.Update(ctx, req.ID, req.Body.Content)
	if err != nil {
		return nil, sharedApi.MapError(
			err,
			"更新失败",
			errorMappings...,
		)
	}
	return sharedApi.NewEmptySuccessResponse(), nil
}

func (h *PostHandler) Delete(ctx context.Context, req *dto.DeletePostRequest) (*sharedApi.BodyResponse[any], error) {
	err := h.service.Delete(ctx, req.ID)
	if err != nil {
		return nil, sharedApi.MapError(
			err,
			"删除失败",
			errorMappings...,
		)
	}

	return sharedApi.NewEmptySuccessResponse(), nil
}
