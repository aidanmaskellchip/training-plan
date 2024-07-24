package repository

import (
	"gorm.io/gorm"
	"training-plan/internal/domain/model"
)

type RunningProfileRepoMock struct {
	db *gorm.DB
}

func (ur RunningProfileRepoMock) Create(profile model.RunningProfile) error {
	return nil
}
