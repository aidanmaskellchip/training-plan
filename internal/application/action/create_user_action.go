package action

import (
	"training-plan/internal/domain/model"
	"training-plan/internal/infrastructure/repository"
	"training-plan/internal/transport/request"
)

func CreateUserAction(data *request.CreateUserRequest, repos *repository.Repositories) (*model.User, error) {
	if err := data.Validate(); err != nil {
		return nil, err
	}

	user := model.User{
		Username: data.Username,
	}

	res, err := repos.UserRepository.Create(user)
	if err != nil {
		return nil, err
	}

	return res, nil
}
