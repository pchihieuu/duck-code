package profile

import (
	"context"

	"github.com/google/uuid"

	apperrors "backend/pkg/errors"
)

type Service interface {
	// GetOrCreate returns the user's profile, lazily creating a row with
	// defaults on first access — so registration doesn't need to know
	// anything about this module (no coupling to internal/auth).
	GetOrCreate(ctx context.Context, userID uuid.UUID) (*Profile, error)
	Update(ctx context.Context, userID uuid.UUID, req UpdateRequest) (*Profile, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetOrCreate(ctx context.Context, userID uuid.UUID) (*Profile, error) {
	p, err := s.repo.FindByUserID(ctx, userID)
	if err == nil {
		return p, nil
	}
	if !apperrors.Is(err, ErrNotFound) {
		return nil, apperrors.Internal(err)
	}

	p = &Profile{UserID: userID, Timezone: "UTC", ShowOnLeaderboard: true}
	if err := s.repo.Create(ctx, p); err != nil {
		return nil, apperrors.Internal(err)
	}
	return p, nil
}

func (s *service) Update(ctx context.Context, userID uuid.UUID, req UpdateRequest) (*Profile, error) {
	// Ensure a row exists before updating (handles a user whose profile row
	// was never lazily created yet, e.g. an account created before this
	// module existed).
	if _, err := s.GetOrCreate(ctx, userID); err != nil {
		return nil, err
	}

	updates := map[string]interface{}{}
	if req.Bio != nil {
		updates["bio"] = *req.Bio
	}
	if req.GithubURL != nil {
		updates["github_url"] = *req.GithubURL
	}
	if req.WebsiteURL != nil {
		updates["website_url"] = *req.WebsiteURL
	}
	if req.Timezone != nil {
		updates["timezone"] = *req.Timezone
	}
	if req.PreferredLanguage != nil {
		updates["preferred_language"] = *req.PreferredLanguage
	}
	if req.ShowOnLeaderboard != nil {
		updates["show_on_leaderboard"] = *req.ShowOnLeaderboard
	}

	if err := s.repo.Update(ctx, userID, updates); err != nil {
		return nil, apperrors.Internal(err)
	}
	return s.repo.FindByUserID(ctx, userID)
}
