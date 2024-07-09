package action

import (
	"training-plan/internal/data/model"
	"training-plan/internal/data/repository"
	"training-plan/internal/transport/request"
)

func CreateUserAction(data *request.CreateUserRequest, repos *repository.Repositories) (err error) {
	if err := data.Validate(); err != nil {
		return err
	}

	user := model.User{
		Username: data.Username,
	}

	if err := repos.UserRepository.Create(user); err != nil {
		return err
	}

	return nil
}
