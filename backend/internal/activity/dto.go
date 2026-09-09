// dto.go
package activity

import (
	"time"

	"github.com/google/uuid"
)

type Response struct {
	ID          uuid.UUID  `json:"id"`
	Type        string     `json:"type"`
	ReferenceID *uuid.UUID `json:"reference_id,omitempty"`
	OccurredAt  time.Time  `json:"occurred_at"`
}

func ToResponse(l *Log) Response {
	return Response{ID: l.ID, Type: l.Type, ReferenceID: l.ReferenceID, OccurredAt: l.OccurredAt}
}