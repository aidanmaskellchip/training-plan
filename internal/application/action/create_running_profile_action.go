package action

import (
	"errors"
	"time"
	"training-plan/internal/domain/model"
	valueobjects "training-plan/internal/domain/value_objects"
	"training-plan/internal/infrastructure/repository"
	"training-plan/internal/transport/request"
)

func CreateRunningProfileAction(data *request.CreateRunningProfileRequest, repos *repository.Repositories) (err error) {
	if err := data.Validate(); err != nil {
		return err
	}

	_, err = repos.UserRepository.FindByID(data.UserID)
	if err != nil {
		return errors.New("user not found")
	}

	startDate, err := time.Parse("2006-01-02", data.StartDate)
	if err != nil {
		return err
	}

	goalDate, err := time.Parse("2006-01-02", data.GoalDate)
	if err != nil {
		return err
	}

	rd := valueobjects.NewRunningDays(data.RunningDays)
	rdJson, err := rd.ToJson()
	if err != nil {
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
		RunningDays:         rdJson,
		RunningDaysPerWeek:  data.RunningDaysPerWeek,
		LongRunDay:          data.LongRunDay,
		CurrentAbility:      data.CurrentAbility,
		PlanLength:          data.PlanLength,
		StartDate:           startDate,
		GoalDate:            goalDate,
	}

	if err := repos.RunningProfileRepository.Create(rp); err != nil {
		return err
	}

	return nil
}
