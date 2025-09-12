package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	vo "training-plan/internal/api/domain/plan/entities"
	useractivity "training-plan/internal/api/domain/user_activity"
)

type UserActivity struct {
	ID        uuid.UUID    `json:"id" gorm:"type:uuid;primaryKey;"`
	UserID    uuid.UUID    `json:"user_id" gorm:"type:uuid;not null;"`
	Type      string       `json:"type" gorm:"type:varchar(255);not null"`
	Distance  float32      `json:"distance" gorm:"numeric;not null"`
	Pace      float32      `json:"pace" gorm:"numeric;not null"`
	Intervals vo.Intervals `json:"intervals" gorm:"json"`
	gorm.Model
}

func (ua *UserActivity) BeforeCreate(tx *gorm.DB) (err error) {
	ua.ID = uuid.New()

	return
}

func (ua *UserActivity) ToDomainEntity() *useractivity.Entity {
	return &useractivity.Entity{
		ID:        ua.ID,
		UserID:    ua.UserID,
		Type:      ua.Type,
		Distance:  ua.Distance,
		Pace:      ua.Pace,
		Intervals: ua.Intervals,
	}
}
