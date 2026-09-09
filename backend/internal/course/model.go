// model.go
package course

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Course struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	Title       string         `gorm:"not null" json:"title"`
	Slug        string         `gorm:"uniqueIndex;not null" json:"slug"`
	Description string         `json:"description"`
	Language    string         `gorm:"not null" json:"language"`
	OrderIndex  int            `gorm:"not null;default:0" json:"order_index"`
	IsPublished bool           `gorm:"not null;default:false" json:"is_published"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Course) TableName() string { return "courses" }