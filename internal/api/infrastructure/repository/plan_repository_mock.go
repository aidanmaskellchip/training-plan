package repository

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"training-plan/internal/api/domain/plan"
)

type PlanRepoMock struct {
	_ *gorm.DB
}

func (pr PlanRepoMock) Create(_ *plan.Entity) error {
	return nil
}

func (pr PlanRepoMock) FindByID(_ uuid.UUID) (p *plan.Entity, err error) {
	return p, nil
}

func (pr PlanRepoMock) FindLatestUserPlan(id uuid.UUID) (p *plan.Entity, err error) {
	if id == getMagicFailingID(MagicFailingUserId) {
		return p, fmt.Errorf("plan not found")
	}

	return p, nil
}
