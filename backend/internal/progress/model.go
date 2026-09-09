package progress

import (
	"time"

	"github.com/google/uuid"
)

type Progress struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	UserID      uuid.UUID  `gorm:"type:uuid;not null;index" json:"user_id"`
	LessonID    uuid.UUID  `gorm:"type:uuid;not null" json:"lesson_id"`
	Status      string     `gorm:"not null;default:not_started" json:"status"` // not_started | in_progress | completed
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

func (Progress) TableName() string { return "user_progress" }