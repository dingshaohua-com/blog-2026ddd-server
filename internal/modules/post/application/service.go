package application

import (
	"blog-2026ddd-server/internal/modules/post/domain"
	"context"
)

type PostService struct {
	repo domain.PostRepository
}

func NewPostService(repo domain.PostRepository) *PostService {
	return &PostService{
		repo: repo,
	}
}

func (s *PostService) List(c context.Context) ([]*domain.Post, error) {
	res, err := s.repo.List(c)
	return res, err
}
