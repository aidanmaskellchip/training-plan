package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"training-plan/internal/data/domain"
)

type Plan struct {
	ID           uuid.UUID         `json:"id" gorm:"type:uuid;primaryKey;"`
	UserID       uuid.UUID         `json:"user_id" gorm:"type:uuid;not null;"`
	Length       int               `json:"length" gorm:"type:int;not null;"`
	GoalDistance int               `json:"goal_distance" gorm:"type:int;not null;"`
	Week1        []domain.Activity `json:"week_1" gorm:"type:json;"`
	Week2        []domain.Activity `json:"week_2" gorm:"type:json;"`
	Week3        []domain.Activity `json:"week_3" gorm:"type:json;"`
	Week4        []domain.Activity `json:"week_4" gorm:"type:json;"`
	Week5        []domain.Activity `json:"week_5" gorm:"type:json;"`
	Week6        []domain.Activity `json:"week_6" gorm:"type:json;"`
	Week7        []domain.Activity `json:"week_7" gorm:"type:json;"`
	Week8        []domain.Activity `json:"week_8" gorm:"type:json;"`
	Week9        []domain.Activity `json:"week_9" gorm:"type:json;"`
	Week10       []domain.Activity `json:"week_10" gorm:"type:json;"`
	Week11       []domain.Activity `json:"week_11" gorm:"type:json;"`
	Week12       []domain.Activity `json:"week_12" gorm:"type:json;"`
	gorm.Model
}
