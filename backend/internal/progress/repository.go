// package progress

// import (
// 	"context"
// 	"errors"
// 	"time"

// 	"github.com/google/uuid"
// 	"gorm.io/gorm"
// )

// var ErrNotFound = errors.New("progress not found")

// type Repository interface {
// 	Upsert(ctx context.Context, p *Progress) error
// 	FindByUserAndLesson(ctx context.Context, userID, lessonID uuid.UUID) (*Progress, error)
// 	ListByUserAndCourse(ctx context.Context, userID, courseID uuid.UUID) ([]Progress, error)
// }

// type repository struct {
// 	db *gorm.DB
// }

// func NewRepository(db *gorm.DB) Repository {
// 	return &repository{db: db}
// }

// // Upsert dùng ON CONFLICT (user_id, lesson_id) — cần UNIQUE constraint đã có
// // ở migration 000005. completed_at/updated_at luôn được ghi đè khi conflict.
// func (r *repository) Upsert(ctx context.Context, p *Progress) error {
// 	if p.ID == uuid.Nil {
// 		p.ID = uuid.New()
// 	}
// 	now := time.Now()
// 	p.UpdatedAt = now
// 	return r.db.WithContext(ctx).
// 		Clauses(gorm.OnConflict{
// 			Columns:   []gorm.Column{{Name: "user_id"}, {Name: "lesson_id"}},
// 			DoUpdates: gorm.AssignmentColumns([]string{"status", "completed_at", "updated_at"}),
// 		}).
// 		Create(p).Error
// }

// func (r *repository) FindByUserAndLesson(ctx context.Context, userID, lessonID uuid.UUID) (*Progress, error) {
// 	var p Progress
// 	err := r.db.WithContext(ctx).
// 		Where("user_id = ? AND lesson_id = ?", userID, lessonID).
// 		First(&p).Error
// 	if err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return nil, ErrNotFound
// 		}
// 		return nil, err
// 	}
// 	return &p, nil
// }

// // ListByUserAndCourse join qua lessons để lọc theo course_id — cần
// // lessons.course_id (đã có index idx_lessons_course_id).
// func (r *repository) ListByUserAndCourse(ctx context.Context, userID, courseID uuid.UUID) ([]Progress, error) {
// 	var items []Progress
// 	err := r.db.WithContext(ctx).
// 		Joins("JOIN lessons ON lessons.id = user_progress.lesson_id").
// 		Where("user_progress.user_id = ? AND lessons.course_id = ?", userID, courseID).
// 		Find(&items).Error
// 	return items, err
// }

package progress

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var ErrNotFound = errors.New("progress not found")

type Repository interface {
	Upsert(ctx context.Context, p *Progress) error
	FindByUserAndLesson(ctx context.Context, userID, lessonID uuid.UUID) (*Progress, error)
	ListByUserAndCourse(ctx context.Context, userID, courseID uuid.UUID) ([]Progress, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

// Upsert dùng ON CONFLICT (user_id, lesson_id) — cần UNIQUE constraint đã có
// ở migration 000005. completed_at/updated_at luôn được ghi đè khi conflict.
func (r *repository) Upsert(ctx context.Context, p *Progress) error {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	now := time.Now()
	p.UpdatedAt = now
	return r.db.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "user_id"}, {Name: "lesson_id"}},
			DoUpdates: clause.AssignmentColumns([]string{"status", "completed_at", "updated_at"}),
		}).
		Create(p).Error
}

func (r *repository) FindByUserAndLesson(ctx context.Context, userID, lessonID uuid.UUID) (*Progress, error) {
	var p Progress
	err := r.db.WithContext(ctx).
		Where("user_id = ? AND lesson_id = ?", userID, lessonID).
		First(&p).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &p, nil
}

// ListByUserAndCourse join qua lessons để lọc theo course_id — cần
// lessons.course_id (đã có index idx_lessons_course_id).
func (r *repository) ListByUserAndCourse(ctx context.Context, userID, courseID uuid.UUID) ([]Progress, error) {
	var items []Progress
	err := r.db.WithContext(ctx).
		Joins("JOIN lessons ON lessons.id = user_progress.lesson_id").
		Where("user_progress.user_id = ? AND lessons.course_id = ?", userID, courseID).
		Find(&items).Error
	return items, err
}