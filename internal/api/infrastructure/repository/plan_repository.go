package repository

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"training-plan/internal/api/domain/plan"
	"training-plan/internal/api/infrastructure/database/model"
)

type PlanRepo struct {
	db *gorm.DB
}

func (pr PlanRepo) Create(plan *plan.Entity) error {
	planModel := &model.Plan{
		ID:           plan.ID,
		UserID:       plan.UserID,
		Length:       plan.Length,
		GoalDistance: plan.GoalDistance,
		Week1:        plan.Week1,
		Week2:        plan.Week2,
		Week3:        plan.Week3,
		Week4:        plan.Week4,
		Week5:        plan.Week5,
		Week6:        plan.Week6,
		Week7:        plan.Week7,
		Week8:        plan.Week8,
		Week9:        plan.Week9,
		Week10:       plan.Week10,
		Week11:       plan.Week11,
		Week12:       plan.Week12,
	}

	result := pr.db.Create(&planModel)

	return result.Error
}

func (pr PlanRepo) FindByID(id uuid.UUID) (p *plan.Entity, err error) {
	planModel := &model.Plan{}

	result := pr.db.First(&planModel, id)

	if result.RowsAffected == 0 {
		return p, fmt.Errorf("running profile not found")
	}

	if result.Error != nil {
		return p, result.Error
	}

	return planModel.ToDomainEntity(), nil
}

func (pr PlanRepo) FindLatestUserPlan(userID uuid.UUID) (p *plan.Entity, err error) {
	planModel := &model.Plan{}

	result := pr.db.
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Limit(1).
		Find(&planModel)

	if result.RowsAffected == 0 {
		return p, fmt.Errorf("plan not found")
	}

	if result.Error != nil {
		return p, result.Error
	}

	return planModel.ToDomainEntity(), nil
}
