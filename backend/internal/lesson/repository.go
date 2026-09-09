// repository.go
package lesson

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var ErrNotFound = errors.New("lesson not found")

type Repository interface {
	Create(ctx context.Context, l *Lesson) error
	FindByID(ctx context.Context, id uuid.UUID) (*Lesson, error)
	ListByCourseID(ctx context.Context, courseID uuid.UUID) ([]Lesson, error)
	Update(ctx context.Context, id uuid.UUID, updates map[string]interface{}) error
	SoftDelete(ctx context.Context, id uuid.UUID) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, l *Lesson) error {
	if l.ID == uuid.Nil {
		l.ID = uuid.New()
	}
	return r.db.WithContext(ctx).Create(l).Error
}

func (r *repository) FindByID(ctx context.Context, id uuid.UUID) (*Lesson, error) {
	var l Lesson
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&l).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &l, nil
}

func (r *repository) ListByCourseID(ctx context.Context, courseID uuid.UUID) ([]Lesson, error) {
	var lessons []Lesson
	err := r.db.WithContext(ctx).Where("course_id = ?", courseID).Order("order_index ASC").Find(&lessons).Error
	return lessons, err
}

func (r *repository) Update(ctx context.Context, id uuid.UUID, updates map[string]interface{}) error {
	if len(updates) == 0 {
		return nil
	}
	return r.db.WithContext(ctx).Model(&Lesson{}).Where("id = ?", id).Updates(updates).Error
}

func (r *repository) SoftDelete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&Lesson{}, "id = ?", id).Error
}