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
	// UpdateRole is admin-only in practice (enforced by middleware.RequireRole
	// at the handler layer, not here) — the repository just persists it.
	UpdateRole(ctx context.Context, id uuid.UUID, role string) error
	// List supports the admin user list (Phase 6), simple offset pagination.
	List(ctx context.Context, limit, offset int) ([]User, int64, error)
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

func (r *repository) UpdateRole(ctx context.Context, id uuid.UUID, role string) error {
	return r.db.WithContext(ctx).Model(&User{}).Where("id = ?", id).Update("role", role).Error
}

func (r *repository) List(ctx context.Context, limit, offset int) ([]User, int64, error) {
	var users []User
	var total int64

	if err := r.db.WithContext(ctx).Model(&User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := r.db.WithContext(ctx).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&users).Error; err != nil {
		return nil, 0, err
	}
	return users, total, nil
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
