package action

import (
	planfactory "training-plan/internal/domain/factory/plan_factory"
	"training-plan/internal/domain/model"
	vo "training-plan/internal/domain/value_objects"
	"training-plan/internal/infrastructure/repository"
	"training-plan/internal/transport/request"
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

	rp, err := repos.RunningProfileRepository.FindLatestUserProfile(userID.ID)
	if err != nil {
		return plan, err
	}

	plan, err = planfactory.NewPlan(rp)
	if err != nil {
		return plan, err
	}

	if err := repos.PlanRepository.Create(plan); err != nil {
		return plan, err
	}

	// TODO: how to cast model struct prop types when retrieving from DB
	res, err := repos.PlanRepository.FindLatestUserPlan(userID.ID)
	if err != nil {
		return plan, err
	}

	return res, nil
}
