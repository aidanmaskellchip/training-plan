package action

import (
	planfactory "training-plan/internal/domain/factory/plan_factory"
	vo "training-plan/internal/domain/value_objects"
	"training-plan/internal/infrastructure/repository"
	"training-plan/internal/transport/request"
)

func CreatePlanAction(data *request.CreatePlanRequest, repos *repository.Repositories) (err error) {
	if err = data.Validate(); err != nil {
		return err
	}

	userID := vo.NewUserID(data.UserID)

	_, err = repos.UserRepository.FindByID(userID.ID)
	if err != nil {
		return err
	}

	rp, err := repos.RunningProfileRepository.FindLatestUserProfile(userID.ID)
	if err != nil {
		return err
	}

	plan, err := planfactory.NewPlan(rp)
	if err != nil {
		return err
	}

	if err := repos.PlanRepository.Create(plan); err != nil {
		return err
	}

	return nil
}
