package repository

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"training-plan/internal/api/domain/model"
)

type PlanRepo struct {
	db *gorm.DB
}

func (pr PlanRepo) Create(plan model.Plan) error {
	result := pr.db.Create(&plan)

	return result.Error
}

func (pr PlanRepo) FindByID(id uuid.UUID) (p model.Plan, err error) {
	result := pr.db.First(&p, id)

	if result.RowsAffected == 0 {
		return p, fmt.Errorf("running profile not found")
	}

	if result.Error != nil {
		return p, result.Error
	}

	return p, nil
}

func (pr PlanRepo) FindLatestUserPlan(userID uuid.UUID) (p model.Plan, err error) {
	result := pr.db.
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Limit(1).
		Find(&p)

	if result.RowsAffected == 0 {
		return p, fmt.Errorf("plan not found")
	}

	if result.Error != nil {
		return p, result.Error
	}

	return p, nil
}
