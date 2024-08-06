package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	vo "training-plan/internal/domain/value_objects"
)

type Plan struct {
	ID           uuid.UUID     `json:"id" gorm:"type:uuid;primaryKey;"`
	UserID       uuid.UUID     `json:"user_id" gorm:"type:uuid;not null;"`
	Length       int           `json:"length" gorm:"type:int;not null;"`
	GoalDistance int           `json:"goal_distance" gorm:"type:int;not null;"`
	Week1        []vo.Activity `json:"week_1" gorm:"type:json;"`
	Week2        []vo.Activity `json:"week_2" gorm:"type:json;"`
	Week3        []vo.Activity `json:"week_3" gorm:"type:json;"`
	Week4        []vo.Activity `json:"week_4" gorm:"type:json;"`
	Week5        []vo.Activity `json:"week_5" gorm:"type:json;"`
	Week6        []vo.Activity `json:"week_6" gorm:"type:json;"`
	Week7        []vo.Activity `json:"week_7" gorm:"type:json;"`
	Week8        []vo.Activity `json:"week_8" gorm:"type:json;"`
	Week9        []vo.Activity `json:"week_9" gorm:"type:json;"`
	Week10       []vo.Activity `json:"week_10" gorm:"type:json;"`
	Week11       []vo.Activity `json:"week_11" gorm:"type:json;"`
	Week12       []vo.Activity `json:"week_12" gorm:"type:json;"`
	gorm.Model
}

func (p *Plan) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New()

	return
}
