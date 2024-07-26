package repository

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"training-plan/internal/domain/model"
)

type RunningProfileRepo struct {
	db *gorm.DB
}

func (rpr RunningProfileRepo) Create(profile model.RunningProfile) error {
	result := rpr.db.Create(&profile)

	return result.Error
}

func (rpr RunningProfileRepo) FindByID(id uuid.UUID) (rp model.RunningProfile, err error) {
	result := rpr.db.First(&rp, id)

	if result.RowsAffected == 0 {
		return rp, fmt.Errorf("running profile not found")
	}

	if result.Error != nil {
		return rp, result.Error
	}

	return rp, nil
}
