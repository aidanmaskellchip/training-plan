package query

import (
	"training-plan/internal/data/domain"
	"training-plan/internal/data/repository"
	"training-plan/internal/transport/response"
)

func FindUserQuery(id *string, repos *repository.Repositories) (res *response.FindUserResponse, err error) {
	userID := domain.NewUserID(*id)

	user, err := repos.UserRepository.FindByID(userID.ID)
	if err != nil {
		return res, err
	}

	return &response.FindUserResponse{
		ID:        user.ID.String(),
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
	}, nil
}
