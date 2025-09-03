package query

import (
	"training-plan/internal/api/domain/plan/entities"
	"training-plan/internal/api/domain/value_objects"
	"training-plan/internal/api/infrastructure/repository"
	"training-plan/internal/api/transport/response"
)

func FindUserRunningProfilesQuery(id *string, repos *repository.Repositories) (res []response.FindRunningProfileResponse, err error) {
	userID := valueobjects.NewUserID(*id)

	rps, err := repos.FindByUserID(userID.ID)
	if err != nil {
		return res, err
	}

	for _, rp := range rps {
		runningDays, err := entities.RunningDaysFromJson(rp.RunningDays)
		if err != nil {
			return res, err
		}

		res = append(res, response.FindRunningProfileResponse{
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
		})
	}

	return res, nil
}
