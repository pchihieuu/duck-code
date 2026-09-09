// service.go
package lesson

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"backend/internal/course"
	apperrors "backend/pkg/errors"
)

type Service interface {
	Create(ctx context.Context, req CreateRequest) (*Lesson, error)
	GetByID(ctx context.Context, id uuid.UUID) (*Lesson, error)
	ListByCourseID(ctx context.Context, courseID uuid.UUID) ([]Lesson, error)
	Update(ctx context.Context, id uuid.UUID, req UpdateRequest) (*Lesson, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type service struct {
	repo       Repository
	courseRepo course.Repository
}

func NewService(repo Repository, courseRepo course.Repository) Service {
	return &service{repo: repo, courseRepo: courseRepo}
}

func (s *service) Create(ctx context.Context, req CreateRequest) (*Lesson, error) {
	if _, err := s.courseRepo.FindByID(ctx, req.CourseID); err != nil {
		if errors.Is(err, course.ErrNotFound) {
			return nil, apperrors.BadRequest("course_id does not exist", err)
		}
		return nil, apperrors.Internal(err)
	}

	l := &Lesson{
		ID: uuid.New(), CourseID: req.CourseID, Title: req.Title, Slug: req.Slug,
		ContentMDX: req.ContentMDX, OrderIndex: req.OrderIndex,
	}
	if err := s.repo.Create(ctx, l); err != nil {
		return nil, apperrors.Internal(err)
	}
	return l, nil
}

func (s *service) GetByID(ctx context.Context, id uuid.UUID) (*Lesson, error) {
	l, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, apperrors.NotFound("lesson not found")
		}
		return nil, apperrors.Internal(err)
	}
	return l, nil
}

func (s *service) ListByCourseID(ctx context.Context, courseID uuid.UUID) ([]Lesson, error) {
	lessons, err := s.repo.ListByCourseID(ctx, courseID)
	if err != nil {
		return nil, apperrors.Internal(err)
	}
	return lessons, nil
}

func (s *service) Update(ctx context.Context, id uuid.UUID, req UpdateRequest) (*Lesson, error) {
	if _, err := s.repo.FindByID(ctx, id); err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, apperrors.NotFound("lesson not found")
		}
		return nil, apperrors.Internal(err)
	}

	updates := map[string]interface{}{}
	if req.Title != nil {
		updates["title"] = *req.Title
	}
	if req.ContentMDX != nil {
		updates["content_mdx"] = *req.ContentMDX
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