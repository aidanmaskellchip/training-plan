package user

import (
	"github.com/google/uuid"
	"time"
)

type Entity struct {
	ID        uuid.UUID  `json:"id"`
	Username  string     `json:"username"`
	CreatedAt *time.Time `json:"created_at"`
}
