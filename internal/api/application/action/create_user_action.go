package action

import (
	"training-plan/internal/api/domain/user"
	"training-plan/internal/api/infrastructure/repository"
	"training-plan/internal/api/transport/request"
)

func CreateUserAction(data *request.CreateUserRequest, repos *repository.Repositories) (*user.Entity, error) {
	if err := data.Validate(); err != nil {
		return nil, err
	}

	userEnt := &user.Entity{
		Username: data.Username,
	}

	res, err := repos.UserRepository.Create(userEnt)
	if err != nil {
		return nil, err
	}

	return res, nil
}
