package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Plan struct {
	ID     uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;"`
	UserID uuid.UUID `json:"user_id" gorm:"type:uuid;not null;"`
	gorm.Model
}
