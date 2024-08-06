package query

import (
	vo "training-plan/internal/domain/value_objects"
	"training-plan/internal/infrastructure/repository"
	"training-plan/internal/transport/response"
)

func FindUserRunningProfilesQuery(id *string, repos *repository.Repositories) (res []response.FindRunningProfileResponse, err error) {
	userID := vo.NewUserID(*id)

	rps, err := repos.RunningProfileRepository.FindByUserID(userID.ID)
	if err != nil {
		return res, err
	}

	for _, rp := range rps {
		runningDays, err := vo.FromJson(rp.RunningDays)
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
