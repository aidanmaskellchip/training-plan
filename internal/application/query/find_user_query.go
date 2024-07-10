package query

import (
	"training-plan/internal/data/model"
	"training-plan/internal/data/repository"
	"training-plan/internal/transport/request"
)

func FindUserQuery(data *request.FindUserRequest, repos *repository.Repositories) (user model.User, err error) {
	if err := data.Validate(); err != nil {
		return user, nil
	}

	if user, err = repos.UserRepository.FindByID(data.UserID); err != nil {
		return user, err
	}

	return user, nil
}
