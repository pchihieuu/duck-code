package user

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ErrNotFound is returned when a lookup finds no matching row.
var ErrNotFound = errors.New("user not found")

// Repository defines the persistence contract for users, so the service
// layer never depends on GORM directly (easier to mock/test).
type Repository interface {
	Create(ctx context.Context, u *User) error
	FindByID(ctx context.Context, id uuid.UUID) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByUsername(ctx context.Context, username string) (*User, error)
	UpdateProfile(ctx context.Context, id uuid.UUID, displayName, avatarURL string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, u *User) error {
	return r.db.WithContext(ctx).Create(u).Error
}

func (r *repository) FindByID(ctx context.Context, id uuid.UUID) (*User, error) {
	return r.findOne(ctx, "id = ?", id)
}

func (r *repository) FindByEmail(ctx context.Context, email string) (*User, error) {
	return r.findOne(ctx, "email = ?", email)
}

func (r *repository) FindByUsername(ctx context.Context, username string) (*User, error) {
	return r.findOne(ctx, "username = ?", username)
}

func (r *repository) UpdateProfile(ctx context.Context, id uuid.UUID, displayName, avatarURL string) error {
	updates := map[string]interface{}{}
	if displayName != "" {
		updates["display_name"] = displayName
	}
	if avatarURL != "" {
		updates["avatar_url"] = avatarURL
	}
	if len(updates) == 0 {
		return nil
	}
	return r.db.WithContext(ctx).Model(&User{}).Where("id = ?", id).Updates(updates).Error
}

func (r *repository) findOne(ctx context.Context, cond string, args ...interface{}) (*User, error) {
	var u User
	err := r.db.WithContext(ctx).Where(cond, args...).First(&u).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &u, nil
}
