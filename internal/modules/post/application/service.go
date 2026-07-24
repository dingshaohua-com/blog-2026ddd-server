package application

import (
	"blog-2026ddd-server/internal/modules/post/domain"
	"context"
	"time"
)

type PostService struct {
	repo domain.PostRepository
}

func NewPostService(repo domain.PostRepository) *PostService {
	return &PostService{
		repo: repo,
	}
}

func (s *PostService) Create(
	ctx context.Context,
	content string,
) (*domain.Post, error) {
	post, err := domain.NewPost(content, time.Now())
	if err != nil {
		return nil, err
	}

	return s.repo.Create(ctx, post)
}

func (s *PostService) Update(
	ctx context.Context,
	id int,
	content string,
) error {
	post, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if err := post.ChangeContent(content); err != nil {
		return err
	}

	return s.repo.Update(ctx, post)
}

func (s *PostService) List(c context.Context) ([]*domain.Post, error) {
	res, err := s.repo.List(c)
	return res, err
}

func (s *PostService) Get(
	ctx context.Context,
	id int,
) (*domain.Post, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *PostService) Delete(
	ctx context.Context,
	id int,
) error {
	_, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	return s.repo.Delete(ctx, id)
}
