package repository

import (
	"gorm.io/gorm"
	"training-plan/internal/data/model"
)

type PlanRepoMock struct {
	db *gorm.DB
}

func (ur PlanRepoMock) Create(plan model.Plan) error {
	return nil
}
