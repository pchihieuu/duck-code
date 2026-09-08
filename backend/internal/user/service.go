package user

import (
	"context"

	"github.com/google/uuid"

	apperrors "backend/pkg/errors"
)

type Service interface {
	GetByID(ctx context.Context, id uuid.UUID) (*User, error)
	UpdateProfile(ctx context.Context, id uuid.UUID, req UpdateProfileRequest) (*User, error)
	// UpdateRole promotes/demotes a user. actorID is the admin performing the
	// action — passed in so the service can block self-demotion, not for
	// authorization (that's already enforced by middleware.RequireRole before
	// the handler is even reached).
	UpdateRole(ctx context.Context, actorID, targetID uuid.UUID, role string) (*User, error)
	List(ctx context.Context, page, pageSize int) ([]User, int64, error)
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

func (s *service) UpdateRole(ctx context.Context, actorID, targetID uuid.UUID, role string) (*User, error) {
	// Block self-demotion so an admin can never accidentally lock themselves
	// (or, worse, the last remaining admin) out of admin-only routes.
	if actorID == targetID {
		return nil, apperrors.BadRequest("cannot change your own role", nil)
	}

	if _, err := s.GetByID(ctx, targetID); err != nil {
		return nil, err // already mapped to apperrors.NotFound / Internal
	}

	if err := s.repo.UpdateRole(ctx, targetID, role); err != nil {
		return nil, apperrors.Internal(err)
	}
	return s.GetByID(ctx, targetID)
}

func (s *service) List(ctx context.Context, page, pageSize int) ([]User, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	users, total, err := s.repo.List(ctx, pageSize, (page-1)*pageSize)
	if err != nil {
		return nil, 0, apperrors.Internal(err)
	}
	return users, total, nil
}
