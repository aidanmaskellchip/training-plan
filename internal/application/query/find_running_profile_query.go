package query

import (
	vo "training-plan/internal/domain/value_objects"
	"training-plan/internal/infrastructure/repository"
	"training-plan/internal/transport/response"
)

func FindRunningProfileQuery(id *string, repos *repository.Repositories) (res *response.FindRunningProfileResponse, err error) {
	uuid := vo.NewUUID(*id)

	rp, err := repos.RunningProfileRepository.FindByID(uuid.ID)
	if err != nil {
		return res, err
	}

	runningDays, err := vo.FromJson(rp.RunningDays)
	if err != nil {
		return res, err
	}

	return &response.FindRunningProfileResponse{
		ID:                  rp.ID.String(),
		UserID:              rp.UserID.String(),
		GoalDistance:        rp.GoalDistance,
		GoalTime:            rp.GoalTime,
		Terrain:             rp.Terrain,
		Current5K:           rp.Current5K,
		Current10K:          rp.Current10K,
		CurrentHalfMarathon: rp.CurrentHalfMarathon,
		CurrentFullMarathon: rp.CurrentFullMarathon,
		RunningDays:         runningDays.Days,
		RunningDaysPerWeek:  rp.RunningDaysPerWeek,
		LongRunDay:          rp.LongRunDay,
		CurrentAbility:      rp.CurrentAbility,
		PlanLength:          rp.PlanLength,
		StartDate:           rp.StartDate,
		GoalDate:            rp.GoalDate,
	}, nil
}
