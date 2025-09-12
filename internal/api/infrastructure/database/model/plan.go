package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"training-plan/internal/api/domain/plan"
	"training-plan/internal/api/domain/plan/entities"
)

type Plan struct {
	ID           uuid.UUID             `json:"id" gorm:"type:uuid;primaryKey;"`
	UserID       uuid.UUID             `json:"user_id" gorm:"type:uuid;not null;"`
	Length       int                   `json:"length" gorm:"type:int;not null;"`
	GoalDistance string                `json:"goal_distance" gorm:"type:string;not null;"`
	Week1        entities.ActivityWeek `json:"week_1" gorm:"type:json;"`
	Week2        entities.ActivityWeek `json:"week_2" gorm:"type:json;"`
	Week3        entities.ActivityWeek `json:"week_3" gorm:"type:json;"`
	Week4        entities.ActivityWeek `json:"week_4" gorm:"type:json;"`
	Week5        entities.ActivityWeek `json:"week_5" gorm:"type:json;"`
	Week6        entities.ActivityWeek `json:"week_6" gorm:"type:json;"`
	Week7        entities.ActivityWeek `json:"week_7" gorm:"type:json;"`
	Week8        entities.ActivityWeek `json:"week_8" gorm:"type:json;"`
	Week9        entities.ActivityWeek `json:"week_9" gorm:"type:json;"`
	Week10       entities.ActivityWeek `json:"week_10" gorm:"type:json;"`
	Week11       entities.ActivityWeek `json:"week_11" gorm:"type:json;"`
	Week12       entities.ActivityWeek `json:"week_12" gorm:"type:json;"`
	gorm.Model
}

func (p *Plan) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New()

	return
}

func (p *Plan) ToDomainEntity() *plan.Entity {
	return &plan.Entity{
		ID:           p.ID,
		UserID:       p.UserID,
		Length:       p.Length,
		GoalDistance: p.GoalDistance,
		Week1:        p.Week1,
		Week2:        p.Week2,
		Week3:        p.Week3,
		Week4:        p.Week4,
		Week5:        p.Week5,
		Week6:        p.Week6,
		Week7:        p.Week7,
		Week8:        p.Week8,
		Week9:        p.Week9,
		Week10:       p.Week10,
		Week11:       p.Week11,
		Week12:       p.Week12,
	}
}
