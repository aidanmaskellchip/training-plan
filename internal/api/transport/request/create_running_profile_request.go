package request

import (
	"errors"
	"github.com/google/uuid"
	"training-plan/internal/api/domain/value_objects"
)

type CreateRunningProfileRequest struct {
	UserID              uuid.UUID `json:"user_id"`
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
	StartDate           string    `json:"start_date"`
	GoalDate            string    `json:"goal_date"`
}

func (c *CreateRunningProfileRequest) Validate() error {
	var n int
	for _, v := range c.RunningDays {
		if v == 1 {
			n++
		}
	}

	if n < c.RunningDaysPerWeek {
		return errors.New("not enough running days selected to allocate all runs")
	}

	if c.RunningDaysPerWeek < 2 || c.RunningDaysPerWeek > 5 {
		return errors.New("invalid number of running days selected")
	}

	if c.RunningDays[c.LongRunDay] != 1 {
		return errors.New("long run day not able to be allocated from running days")
	}

	if c.PlanLength != 12 {
		return errors.New("plan length not valid")
	}

	err := errors.New("invalid terrain type")
	for _, v := range valueobjects.GetTerrainStrings() {
		if v == c.Terrain {
			err = nil
		}
	}

	return err
}
