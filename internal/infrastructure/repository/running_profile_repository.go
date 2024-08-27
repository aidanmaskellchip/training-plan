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

func (rpr RunningProfileRepo) FindByUserID(userID uuid.UUID) (rps []model.RunningProfile, err error) {
	result := rpr.db.Where("user_id = ?", userID).Find(&rps)

	if result.RowsAffected == 0 {
		return rps, fmt.Errorf("running profiles not found")
	}

	if result.Error != nil {
		return rps, result.Error
	}

	return rps, nil
}

func (rpr RunningProfileRepo) FindLatestUserProfile(id uuid.UUID) (rp model.RunningProfile, err error) {
	result := rpr.db.Order("created_at desc").Where("user_id = ?", id).Find(&rp)

	if result.RowsAffected == 0 {
		return rp, fmt.Errorf("running profile not found")
	}

	if result.Error != nil {
		return rp, result.Error
	}

	return rp, nil
}
