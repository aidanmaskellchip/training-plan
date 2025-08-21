package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type RunningProfile struct {
	ID                  uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;"`
	UserID              uuid.UUID `json:"user_id" gorm:"type:uuid;not null;"`
	GoalDistance        string    `json:"goal_distance" gorm:"type:text;"`
	GoalTime            int       `json:"goal_time" gorm:"type:int;"`
	Terrain             string    `json:"terrain" gorm:"type:text;"`
	Current5K           int       `json:"current_5_k" gorm:"type:int;"`
	Current10K          int       `json:"current_10_k" gorm:"type:int;"`
	CurrentHalfMarathon int       `json:"current_half_marathon" gorm:"type:int;"`
	CurrentFullMarathon int       `json:"current_full_marathon" gorm:"type:int;"`
	RunningDays         []byte    `json:"running_days" gorm:"type:json;"`
	RunningDaysPerWeek  int       `json:"running_days_per_week" gorm:"type:int;"`
	LongRunDay          int       `json:"long_run_day" gorm:"type:int;"`
	CurrentAbility      string    `json:"current_ability" gorm:"type:text;"`
	PlanLength          int       `json:"plan_length" gorm:"type:int;"`
	StartDate           time.Time `json:"start_date" gorm:"type:date;"`
	GoalDate            time.Time `json:"goal_date" gorm:"type:date;"`
	gorm.Model
}

func (rp *RunningProfile) BeforeCreate(tx *gorm.DB) (err error) {
	rp.ID = uuid.New()

	return
}
