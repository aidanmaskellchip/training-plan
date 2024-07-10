package action

import (
	"training-plan/internal/data/domain"
	"training-plan/internal/data/model"
	"training-plan/internal/data/repository"
	"training-plan/internal/transport/request"
)

func CreatePlanAction(data *request.CreatePlanRequest, repos *repository.Repositories) (user model.User, err error) {
	if err := data.Validate(); err != nil {
		return user, nil
	}

	userID := domain.NewUserID(data.UserID)

	if user, err = repos.UserRepository.FindByID(userID.ID); err != nil {
		return user, err
	}

	//plan := model.Plan{}
	//if err := repos.PlanRepository.Create(plan); err != nil {
	//	return err
	//}
	//
	return user, nil
}
