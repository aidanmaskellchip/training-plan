package repository

import (
	"gorm.io/gorm"
	"training-plan/internal/domain/model"
)

type PlanRepoMock struct {
	db *gorm.DB
}

func (ur PlanRepoMock) Create(plan model.Plan) error {
	return nil
}
