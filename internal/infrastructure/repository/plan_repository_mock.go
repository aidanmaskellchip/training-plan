package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"training-plan/internal/domain/model"
)

type PlanRepoMock struct {
	db *gorm.DB
}

func (pr PlanRepoMock) Create(plan model.Plan) error {
	return nil
}

func (pr PlanRepoMock) FindByID(id uuid.UUID) (p model.Plan, err error) {
	return p, nil
}

func (pr PlanRepoMock) FindLatestUserPlan(id uuid.UUID) (p model.Plan, err error) {
	return p, nil
}
