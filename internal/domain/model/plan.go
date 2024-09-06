package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	vo "training-plan/internal/domain/value_objects"
)

type Plan struct {
	ID           uuid.UUID       `json:"id" gorm:"type:uuid;primaryKey;"`
	UserID       uuid.UUID       `json:"user_id" gorm:"type:uuid;not null;"`
	Length       int             `json:"length" gorm:"type:int;not null;"`
	GoalDistance vo.GoalDistance `json:"goal_distance" gorm:"type:string;not null;"`
	Week1        ActivityWeek    `json:"week_1" gorm:"type:json;"`
	Week2        ActivityWeek    `json:"week_2" gorm:"type:json;"`
	Week3        ActivityWeek    `json:"week_3" gorm:"type:json;"`
	Week4        ActivityWeek    `json:"week_4" gorm:"type:json;"`
	Week5        ActivityWeek    `json:"week_5" gorm:"type:json;"`
	Week6        ActivityWeek    `json:"week_6" gorm:"type:json;"`
	Week7        ActivityWeek    `json:"week_7" gorm:"type:json;"`
	Week8        ActivityWeek    `json:"week_8" gorm:"type:json;"`
	Week9        ActivityWeek    `json:"week_9" gorm:"type:json;"`
	Week10       ActivityWeek    `json:"week_10" gorm:"type:json;"`
	Week11       ActivityWeek    `json:"week_11" gorm:"type:json;"`
	Week12       ActivityWeek    `json:"week_12" gorm:"type:json;"`
	gorm.Model
}

func (p *Plan) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New()

	return
}
