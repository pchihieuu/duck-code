package profile

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var ErrNotFound = errors.New("profile not found")

type Repository interface {
	Create(ctx context.Context, p *Profile) error
	FindByUserID(ctx context.Context, userID uuid.UUID) (*Profile, error)
	Update(ctx context.Context, userID uuid.UUID, updates map[string]interface{}) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, p *Profile) error {
	return r.db.WithContext(ctx).Create(p).Error
}

func (r *repository) FindByUserID(ctx context.Context, userID uuid.UUID) (*Profile, error) {
	var p Profile
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).First(&p).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &p, nil
}

func (r *repository) Update(ctx context.Context, userID uuid.UUID, updates map[string]interface{}) error {
	if len(updates) == 0 {
		return nil
	}
	return r.db.WithContext(ctx).Model(&Profile{}).Where("user_id = ?", userID).Updates(updates).Error
}
