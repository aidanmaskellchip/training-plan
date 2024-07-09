package repository

import (
	"gorm.io/gorm"
	"training-plan/internal/data/model"
)

type RunningProfileRepoMock struct {
	db *gorm.DB
}

func (ur RunningProfileRepoMock) Create(profile model.RunningProfile) error {
	return nil
}
