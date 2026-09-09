package gamification

import (
	"context"
	"math"
	"time"

	"github.com/google/uuid"

	"backend/internal/user"
	apperrors "backend/pkg/errors"
)

type Service interface {
	AwardXP(ctx context.Context, userID uuid.UUID, amount int) error
	RecordStreak(ctx context.Context, userID uuid.UUID) error
	GetStatus(ctx context.Context, userID uuid.UUID) (*Response, error)
}

type service struct {
	users user.Repository
}

func NewService(users user.Repository) Service {
	return &service{users: users}
}

func (s *service) AwardXP(ctx context.Context, userID uuid.UUID, amount int) error {
	if amount <= 0 {
		return nil
	}
	if err := s.users.AddXP(ctx, userID, amount); err != nil {
		return apperrors.Internal(err)
	}
	return nil
}

// RecordStreak implements BR-052 (lazy-check, per §0.2 của guide Phase 2).
// Gọi 1 lần mỗi "hoạt động hợp lệ" (hoàn thành lesson/exercise/quiz) — gọi
// lại nhiều lần trong cùng 1 ngày là an toàn (no-op từ lần thứ 2).
func (s *service) RecordStreak(ctx context.Context, userID uuid.UUID) error {
	u, err := s.users.FindByID(ctx, userID)
	if err != nil {
		return apperrors.Internal(err)
	}

	today := time.Now().UTC().Truncate(24 * time.Hour)

	newStreak := 1
	switch {
	case u.LastActivityDate == nil:
		newStreak = 1
	case sameDay(*u.LastActivityDate, today):
		return nil // đã ghi nhận hôm nay rồi, không đổi gì
	case sameDay(*u.LastActivityDate, today.AddDate(0, 0, -1)):
		newStreak = u.CurrentStreak + 1
	default:
		newStreak = 1 // hở >= 2 ngày -> reset
	}

	longest := u.LongestStreak
	if newStreak > longest {
		longest = newStreak
	}

	if err := s.users.UpdateStreak(ctx, userID, newStreak, longest, today); err != nil {
		return apperrors.Internal(err)
	}
	return nil
}

func (s *service) GetStatus(ctx context.Context, userID uuid.UUID) (*Response, error) {
	u, err := s.users.FindByID(ctx, userID)
	if err != nil {
		return nil, apperrors.Internal(err)
	}
	return &Response{
		TotalXP:       u.TotalXP,
		Level:         level(u.TotalXP), // BR-051
		CurrentStreak: u.CurrentStreak,
		LongestStreak: u.LongestStreak,
	}, nil
}

// level implements BR-051: level = floor(sqrt(total_xp / 100)) + 1.
// Tuning knob — chỉnh số chia (100) để đổi tốc độ lên level.
func level(totalXP int) int {
	return int(math.Floor(math.Sqrt(float64(totalXP)/100.0))) + 1
}

func sameDay(a, b time.Time) bool {
	return a.Year() == b.Year() && a.YearDay() == b.YearDay()
}