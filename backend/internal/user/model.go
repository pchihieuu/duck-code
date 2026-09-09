package user

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User is the GORM model backing the users table.
type User struct {
	ID           uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	Email        string         `gorm:"uniqueIndex;not null" json:"email"`
	Username     string         `gorm:"uniqueIndex;not null" json:"username"`
	PasswordHash string         `gorm:"not null" json:"-"`
	DisplayName  string         `gorm:"not null" json:"display_name"`
	AvatarURL    *string        `json:"avatar_url,omitempty"`
	Role         string         `gorm:"not null;default:student" json:"role"`

	// Gamification (Phase 2, migration 000007) — 1-1 với user, đọc thường
	// xuyên (header), giữ thẳng trong bảng users thay vì tách bảng riêng.
	TotalXP          int        `gorm:"not null;default:0" json:"total_xp"`
	CurrentStreak    int        `gorm:"not null;default:0" json:"current_streak"`
	LongestStreak    int        `gorm:"not null;default:0" json:"longest_streak"`
	LastActivityDate *time.Time `gorm:"type:date" json:"last_activity_date,omitempty"`

	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

func (User) TableName() string {
	return "users"
}