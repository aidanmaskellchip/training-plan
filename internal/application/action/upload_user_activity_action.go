package action

import (
	"errors"
	"training-plan/internal/domain/factory"
	"training-plan/internal/domain/model"
	"training-plan/internal/infrastructure/repository"
	"training-plan/internal/transport/request"
)

func UploadUserActivityAction(data *request.UploadUserActivityRequest, repos *repository.Repositories) (err error) {
	if err := data.Validate(); err != nil {
		return err
	}

	_, err = repos.UserRepository.FindByID(data.UserID)
	if err != nil {
		return errors.New("user not found")
	}

	at, err := factory.NewActivityType(data.Type)
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
