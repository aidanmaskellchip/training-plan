package action

import (
	"training-plan/internal/api/domain/model"
	"training-plan/internal/api/domain/plan/factory"
	vo "training-plan/internal/api/domain/value_objects"
	"training-plan/internal/api/infrastructure/repository"
	"training-plan/internal/api/transport/request"
)

func CreatePlanAction(data *request.CreatePlanRequest, repos *repository.Repositories) (plan model.Plan, err error) {
	if err = data.Validate(); err != nil {
		return plan, err
	}

	userID := vo.NewUserID(data.UserID)

	_, err = repos.UserRepository.FindByID(userID.ID)
	if err != nil {
		return plan, err
	}

	rp, err := repos.FindLatestUserProfile(userID.ID)
	if err != nil {
		return plan, err
	}

	plan, err = factory.planfactory.NewPlan(rp)
	if err != nil {
		return plan, err
	}

	if err := repos.PlanRepository.Create(plan); err != nil {
		return plan, err
	}

	res, err := repos.FindLatestUserPlan(userID.ID)
	if err != nil {
		return plan, err
	}

	return res, nil
}
