package repository

import (
	"gorm.io/gorm"
	"training-plan/internal/data/model"
)

type RunningProfileRepo struct {
	db *gorm.DB
}

func (rpr RunningProfileRepo) Create(profile model.RunningProfile) error {
	result := rpr.db.Create(&profile)

	return result.Error
}
