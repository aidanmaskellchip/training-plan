package query

import (
	vo "training-plan/internal/domain/value_objects"
	"training-plan/internal/infrastructure/repository"
	"training-plan/internal/transport/response"
)

func FindUserQuery(id *string, repos *repository.Repositories) (res *response.FindUserResponse, err error) {
	userID := vo.NewUserID(*id)

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
