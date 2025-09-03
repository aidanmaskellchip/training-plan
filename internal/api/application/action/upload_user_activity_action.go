package action

import (
	"errors"
	"training-plan/internal/api/domain/model"
	"training-plan/internal/api/domain/plan/factory"
	"training-plan/internal/api/infrastructure/repository"
	"training-plan/internal/api/transport/request"
)

func UploadUserActivityAction(data *request.UploadUserActivityRequest, repos *repository.Repositories) (err error) {
	if err := data.Validate(); err != nil {
		return err
	}

	_, err = repos.UserRepository.FindByID(data.UserID)
	if err != nil {
		return errors.New("user not found")
	}

	at, err := factory.activitytypefactory.NewActivityType(data.Type)
	if err != nil {
		return err
	}

	ua := &model.UserActivity{
		UserID:   data.UserID,
		Type:     at.Type,
		Distance: data.Distance,
		Pace:     data.Pace,
	}

	if err := repos.UserActivityRepository.Create(*ua); err != nil {
		return err
	}

	return nil
}
