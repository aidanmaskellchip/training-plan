package response

import (
	"time"
	"training-plan/internal/api/domain/model"
	"training-plan/internal/api/domain/plan/entities"
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

func NewFindRunningProfileResponse(p model.RunningProfile) (FindRunningProfileResponse, error) {
	rd, err := entities.RunningDaysFromJson(p.RunningDays)
	if err != nil {
		return FindRunningProfileResponse{}, err
	}

	return FindRunningProfileResponse{
		ID:                  p.ID.String(),
		UserID:              p.UserID.String(),
		GoalDistance:        p.GoalDistance,
		GoalTime:            p.GoalTime,
		Terrain:             p.Terrain,
		Current5K:           p.Current5K,
		Current10K:          p.Current10K,
		CurrentHalfMarathon: p.CurrentHalfMarathon,
		CurrentFullMarathon: p.CurrentFullMarathon,
		RunningDays:         rd.Days,
		RunningDaysPerWeek:  p.RunningDaysPerWeek,
		LongRunDay:          p.LongRunDay,
		CurrentAbility:      p.CurrentAbility,
		PlanLength:          p.PlanLength,
		StartDate:           p.StartDate,
		GoalDate:            p.GoalDate,
	}, nil
}
