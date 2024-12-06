package action

import (
	"github.com/google/uuid"
	activitytypefactory "training-plan/internal/domain/factory/activity_type_factory"
	"training-plan/internal/infrastructure/repository"
	"training-plan/internal/transport/request"
)

func EditUserActivityAction(actId string, data *request.EditUserActivityRequest, repos *repository.Repositories) (err error) {
	if err := data.Validate(); err != nil {
		return err
	}

	actUuid, err := uuid.Parse(actId)
	if err != nil {
		return err
	}

	userAct, err := repos.UserActivityRepository.FindByID(actUuid)
	if err != nil {
		return err
	}

	if data.Type != nil {
		at, err := activitytypefactory.NewActivityType(*data.Type)
		if err != nil {
			return err
		}

		userAct.Type = at.Type
	}

	if data.Pace != nil {
		userAct.Pace = *data.Pace
	}

	if data.Distance != nil {
		userAct.Distance = *data.Distance
	}

	if err := repos.UserActivityRepository.Update(userAct); err != nil {
		return err
	}

	return nil
}
