// service.go
package course

import (
	"context"
	"errors"

	"github.com/google/uuid"

	apperrors "backend/pkg/errors"
)

type Service interface {
	Create(ctx context.Context, req CreateRequest) (*Course, error)
	GetBySlug(ctx context.Context, slug string) (*Course, error)
	ListPublished(ctx context.Context) ([]Course, error)
	ListAll(ctx context.Context) ([]Course, error)
	Update(ctx context.Context, id uuid.UUID, req UpdateRequest) (*Course, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) Create(ctx context.Context, req CreateRequest) (*Course, error) {
	if _, err := s.repo.FindBySlug(ctx, req.Slug); err == nil {
		return nil, apperrors.Conflict("slug already exists")
	} else if !errors.Is(err, ErrNotFound) {
		return nil, apperrors.Internal(err)
	}

	c := &Course{
		ID: uuid.New(), Title: req.Title, Slug: req.Slug,
		Description: req.Description, Language: req.Language, OrderIndex: req.OrderIndex,
	}
	if err := s.repo.Create(ctx, c); err != nil {
		return nil, apperrors.Internal(err)
	}
	return c, nil
}

func (s *service) GetBySlug(ctx context.Context, slug string) (*Course, error) {
	c, err := s.repo.FindBySlug(ctx, slug)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, apperrors.NotFound("course not found")
		}
		return nil, apperrors.Internal(err)
	}
	return c, nil
}

func (s *service) ListPublished(ctx context.Context) ([]Course, error) {
	courses, err := s.repo.ListPublished(ctx)
	if err != nil {
		return nil, apperrors.Internal(err)
	}
	return courses, nil
}

func (s *service) ListAll(ctx context.Context) ([]Course, error) {
	courses, err := s.repo.ListAll(ctx)
	if err != nil {
		return nil, apperrors.Internal(err)
	}
	return courses, nil
}

func (s *service) Update(ctx context.Context, id uuid.UUID, req UpdateRequest) (*Course, error) {
	if _, err := s.repo.FindByID(ctx, id); err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, apperrors.NotFound("course not found")
		}
		return nil, apperrors.Internal(err)
	}

	updates := map[string]interface{}{}
	if req.Title != nil {
		updates["title"] = *req.Title
	}
	if req.Description != nil {
		updates["description"] = *req.Description
	}
	if req.Language != nil {
		updates["language"] = *req.Language
	}
	if req.OrderIndex != nil {
		updates["order_index"] = *req.OrderIndex
	}
	if req.IsPublished != nil {
		updates["is_published"] = *req.IsPublished
	}
	if err := s.repo.Update(ctx, id, updates); err != nil {
		return nil, apperrors.Internal(err)
	}
	return s.repo.FindByID(ctx, id)
}

func (s *service) Delete(ctx context.Context, id uuid.UUID) error {
	if err := s.repo.SoftDelete(ctx, id); err != nil {
		return apperrors.Internal(err)
	}
	return nil
}