package progress

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"

	"backend/internal/activity"
	"backend/internal/gamification"
	"backend/internal/lesson"
	apperrors "backend/pkg/errors"
)

type Service interface {
	CompleteLesson(ctx context.Context, userID, lessonID uuid.UUID) error
	GetCourseProgress(ctx context.Context, userID, courseID uuid.UUID) ([]LessonProgressView, error)
}

type service struct {
	repo            Repository
	lessonRepo      lesson.Repository
	activitySvc     activity.Service
	gamificationSvc gamification.Service
}

func NewService(
	repo Repository,
	lessonRepo lesson.Repository,
	activitySvc activity.Service,
	gamificationSvc gamification.Service,
) Service {
	return &service{
		repo:            repo,
		lessonRepo:      lessonRepo,
		activitySvc:     activitySvc,
		gamificationSvc: gamificationSvc,
	}
}

// CompleteLesson: idempotent (gọi lại nhiều lần không lỗi, không cộng XP 2
// lần), rồi ghi activity + XP + streak theo đúng thứ tự trong guide §5.2.
func (s *service) CompleteLesson(ctx context.Context, userID, lessonID uuid.UUID) error {
	// 1. Idempotency check
	existing, err := s.repo.FindByUserAndLesson(ctx, userID, lessonID)
	if err != nil && !errors.Is(err, ErrNotFound) {
		return apperrors.Internal(err)
	}
	if existing != nil && existing.Status == "completed" {
		return nil // đã hoàn thành từ trước, không làm lại
	}

	// 2. Ghi nhận hoàn thành
	now := time.Now()
	if err := s.repo.Upsert(ctx, &Progress{
		UserID: userID, LessonID: lessonID,
		Status: "completed", CompletedAt: &now,
	}); err != nil {
		return apperrors.Internal(err)
	}

	// 3. Ghi activity (nguồn cho streak + lịch sử)
	if err := s.activitySvc.Record(ctx, userID, "lesson_completed", &lessonID); err != nil {
		return apperrors.Internal(err)
	}

	// 4. XP thưởng hoàn thành lesson (BR-050 — 10 XP, tune sau)
	if err := s.gamificationSvc.AwardXP(ctx, userID, 10); err != nil {
		return apperrors.Internal(err)
	}

	// 5. Streak
	return s.gamificationSvc.RecordStreak(ctx, userID)
}

// GetCourseProgress implements unlock check (BR-041): lesson đầu tiên luôn
// unlock, lesson sau chỉ unlock nếu lesson ngay trước đã completed.
func (s *service) GetCourseProgress(ctx context.Context, userID, courseID uuid.UUID) ([]LessonProgressView, error) {
	lessons, err := s.lessonRepo.ListByCourseID(ctx, courseID) // đã sort theo order_index
	if err != nil {
		return nil, apperrors.Internal(err)
	}
	progresses, err := s.repo.ListByUserAndCourse(ctx, userID, courseID)
	if err != nil {
		return nil, apperrors.Internal(err)
	}

	completed := map[uuid.UUID]bool{}
	for _, p := range progresses {
		if p.Status == "completed" {
			completed[p.LessonID] = true
		}
	}

	views := make([]LessonProgressView, len(lessons))
	prevCompleted := true // lesson đầu tiên luôn unlock
	for i, l := range lessons {
		views[i] = LessonProgressView{
			LessonID:  l.ID,
			Completed: completed[l.ID],
			Unlocked:  prevCompleted,
		}
		prevCompleted = completed[l.ID]
	}
	return views, nil
}