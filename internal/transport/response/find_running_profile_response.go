package response

import (
	"time"
)

type FindRunningProfileResponse struct {
	ID                  string    `json:"id"`
	UserID              string    `json:"user_id"`
	GoalDistance        string    `json:"goal_distance"`
	GoalTime            int       `json:"goal_time"`
	Terrain             string    `json:"terrain"`
	Current5K           int       `json:"current_5_k"`
	Current10K          int       `json:"current_10_k"`
	CurrentHalfMarathon int       `json:"current_half_marathon"`
	CurrentFullMarathon int       `json:"current_full_marathon"`
	RunningDays         []int     `json:"running_days"`
	RunningDaysPerWeek  int       `json:"running_days_per_week"`
	LongRunDay          int       `json:"long_run_day"`
	CurrentAbility      string    `json:"current_ability"`
	PlanLength          int       `json:"plan_length"`
	StartDate           time.Time `json:"start_date"`
	GoalDate            time.Time `json:"goal_date"`
}
