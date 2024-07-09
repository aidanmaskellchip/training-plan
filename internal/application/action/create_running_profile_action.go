package action

import (
	"training-plan/internal/data/model"
	"training-plan/internal/data/repository"
	"training-plan/internal/transport/request"
)

func CreateRunningProfileAction(data *request.CreateRunningProfileRequest, repos *repository.Repositories) (err error) {
	if err := data.Validate(); err != nil {
		return err
	}

	rp := model.RunningProfile{
		UserID:              data.UserID,
		GoalDistance:        data.GoalDistance,
		GoalTime:            data.GoalTime,
		Terrain:             data.Terrain,
		Current5K:           data.Current5K,
		Current10K:          data.Current10K,
		CurrentHalfMarathon: data.CurrentHalfMarathon,
		CurrentFullMarathon: data.CurrentFullMarathon,
		RunningDays:         data.RunningDays,
		RunningDaysPerWeek:  data.RunningDaysPerWeek,
		LongRunDay:          data.LongRunDay,
		CurrentAbility:      data.CurrentAbility,
		PlanLength:          data.PlanLength,
		StartDate:           data.StartDate,
		GoalDate:            data.GoalDate,
	}

	if err := repos.RunningProfileRepository.Create(rp); err != nil {
		return err
	}

	return nil
}
