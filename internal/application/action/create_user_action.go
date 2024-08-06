package action

import (
	"github.com/google/uuid"
	"training-plan/internal/domain/model"
	"training-plan/internal/infrastructure/repository"
	"training-plan/internal/transport/request"
)

func CreateUserAction(data *request.CreateUserRequest, repos *repository.Repositories) (err error) {
	if err := data.Validate(); err != nil {
		return err
	}

	user := model.User{
		ID:       uuid.New(),
		Username: data.Username,
	}

	if err := repos.UserRepository.Create(user); err != nil {
		return err
	}

	return nil
}
