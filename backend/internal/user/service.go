package user

import (
	"context"

	"github.com/google/uuid"

	apperrors "backend/pkg/errors"
)

type Service interface {
	GetByID(ctx context.Context, id uuid.UUID) (*User, error)
	UpdateProfile(ctx context.Context, id uuid.UUID, req UpdateProfileRequest) (*User, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetByID(ctx context.Context, id uuid.UUID) (*User, error) {
	u, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if apperrors.Is(err, ErrNotFound) {
			return nil, apperrors.NotFound("user not found")
		}
		return nil, apperrors.Internal(err)
	}
	return u, nil
}

func (s *service) UpdateProfile(ctx context.Context, id uuid.UUID, req UpdateProfileRequest) (*User, error) {
	if err := s.repo.UpdateProfile(ctx, id, req.DisplayName, req.AvatarURL); err != nil {
		return nil, apperrors.Internal(err)
	}
	return s.GetByID(ctx, id)
}
