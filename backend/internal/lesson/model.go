// model.go
package lesson

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Lesson struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	CourseID    uuid.UUID      `gorm:"type:uuid;not null;index" json:"course_id"`
	Title       string         `gorm:"not null" json:"title"`
	Slug        string         `gorm:"not null" json:"slug"`
	ContentMDX  string         `gorm:"not null;default:''" json:"-"`
	OrderIndex  int            `gorm:"not null;default:0" json:"order_index"`
	IsPublished bool           `gorm:"not null;default:false" json:"is_published"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Lesson) TableName() string { return "lessons" }