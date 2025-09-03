package action

import (
	"github.com/google/uuid"
	"training-plan/internal/api/domain/plan/factory"
	"training-plan/internal/api/infrastructure/repository"
	"training-plan/internal/api/transport/request"
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
		at, err := factory.NewActivityType(*data.Type)
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

	if err := repos.Update(userAct); err != nil {
		return err
	}

	return nil
}
