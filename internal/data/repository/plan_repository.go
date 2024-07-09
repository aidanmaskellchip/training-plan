package repository

import (
	"gorm.io/gorm"
	"training-plan/internal/data/model"
)

type PlanRepo struct {
	db *gorm.DB
}

func (pr PlanRepo) Create(plan model.Plan) error {
	result := pr.db.Create(&plan)

	return result.Error
}
