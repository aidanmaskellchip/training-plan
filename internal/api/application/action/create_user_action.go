package action

import (
	"training-plan/internal/api/domain/model"
	"training-plan/internal/api/infrastructure/repository"
	"training-plan/internal/api/transport/request"
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
