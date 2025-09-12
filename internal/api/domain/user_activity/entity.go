package useractivity

import (
	"github.com/google/uuid"
	vo "training-plan/internal/api/domain/plan/entities"
)

type Entity struct {
	ID        uuid.UUID    `json:"id"`
	UserID    uuid.UUID    `json:"user_id"`
	Type      string       `json:"type"`
	Distance  float32      `json:"distance"`
	Pace      float32      `json:"pace"`
	Intervals vo.Intervals `json:"intervals"`
}
