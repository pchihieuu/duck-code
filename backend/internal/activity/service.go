// service.go
package activity

import (
	"context"
	"time"

	"github.com/google/uuid"

	apperrors "backend/pkg/errors"
)

// Service.Record là điểm vào DUY NHẤT để module khác ghi activity — đúng
// rule ở SKILL.md #7: không module nào khác tự Create thẳng vào bảng
// activity_log.
type Service interface {
	Record(ctx context.Context, userID uuid.UUID, activityType string, refID *uuid.UUID) error
	ListRecent(ctx context.Context, userID uuid.UUID, limit int) ([]Response, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) Record(ctx context.Context, userID uuid.UUID, activityType string, refID *uuid.UUID) error {
	log := &Log{UserID: userID, Type: activityType, ReferenceID: refID, OccurredAt: time.Now()}
	if err := s.repo.Create(ctx, log); err != nil {
		return apperrors.Internal(err)
	}
	return nil
}

func (s *service) ListRecent(ctx context.Context, userID uuid.UUID, limit int) ([]Response, error) {
	if limit <= 0 || limit > 100 {
		limit = 20
	}
	logs, err := s.repo.ListRecent(ctx, userID, limit)
	if err != nil {
		return nil, apperrors.Internal(err)
	}
	out := make([]Response, len(logs))
	for i, l := range logs {
		out[i] = ToResponse(&l)
	}
	return out, nil
}