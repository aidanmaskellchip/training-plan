package query

import (
	"training-plan/internal/data/domain"
	"training-plan/internal/data/model"
	"training-plan/internal/data/repository"
)

func FindUserQuery(id *string, repos *repository.Repositories) (user model.User, err error) {
	userID := domain.NewUserID(*id)

	if user, err = repos.UserRepository.FindByID(userID.ID); err != nil {
		return user, err
	}

	return user, nil
}
