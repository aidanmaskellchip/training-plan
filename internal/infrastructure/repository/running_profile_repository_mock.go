package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"training-plan/internal/domain/model"
)

type RunningProfileRepoMock struct {
	db *gorm.DB
}

func (ur RunningProfileRepoMock) Create(_ model.RunningProfile) error {
	return nil
}

func (ur RunningProfileRepoMock) FindByID(_ uuid.UUID) (rp model.RunningProfile, err error) {
	return rp, nil
}

func (rpr RunningProfileRepoMock) FindByUserID(userID uuid.UUID) (rps []model.RunningProfile, err error) {
	return rps, nil
}
