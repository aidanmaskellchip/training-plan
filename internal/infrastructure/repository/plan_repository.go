package repository

import (
	"gorm.io/gorm"
	"training-plan/internal/domain/model"
)

type PlanRepo struct {
	db *gorm.DB
}

func (pr PlanRepo) Create(plan model.Plan) error {
	result := pr.db.Create(&plan)

	return result.Error
}
