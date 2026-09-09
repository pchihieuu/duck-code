// repository.go
package activity

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, l *Log) error
	ListRecent(ctx context.Context, userID uuid.UUID, limit int) ([]Log, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, l *Log) error {
	if l.ID == uuid.Nil {
		l.ID = uuid.New()
	}
	return r.db.WithContext(ctx).Create(l).Error
}

func (r *repository) ListRecent(ctx context.Context, userID uuid.UUID, limit int) ([]Log, error) {
	var logs []Log
	err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Order("occurred_at DESC").
		Limit(limit).
		Find(&logs).Error
	return logs, err
}