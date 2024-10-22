package action

import (
	"github.com/google/uuid"
	"training-plan/internal/domain/model"
	"training-plan/internal/infrastructure/repository"
	"training-plan/internal/transport/request"
)

func CreateUserAction(data *request.CreateUserRequest, repos *repository.Repositories) (id string, err error) {
	if err := data.Validate(); err != nil {
		return "", err
	}

	userId := uuid.New()
	user := model.User{
		ID:       userId,
		Username: data.Username,
	}

	if err := repos.UserRepository.Create(user); err != nil {
		return "", err
	}

	return userId.String(), nil
}
