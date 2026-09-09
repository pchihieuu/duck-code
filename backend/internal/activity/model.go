// model.go
package activity

import (
	"time"

	"github.com/google/uuid"
)

type Log struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	UserID      uuid.UUID  `gorm:"type:uuid;not null;index:idx_activity_log_user_occurred" json:"user_id"`
	Type        string     `gorm:"type:varchar(30);not null" json:"type"`
	ReferenceID *uuid.UUID `gorm:"type:uuid" json:"reference_id,omitempty"`
	OccurredAt  time.Time  `gorm:"not null;index:idx_activity_log_user_occurred" json:"occurred_at"`
}

func (Log) TableName() string { return "activity_log" }