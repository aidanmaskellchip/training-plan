package repository

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"training-plan/internal/domain/model"
)

type UserActivityRepo struct {
	db *gorm.DB
}

func (ur UserActivityRepo) Create(ua model.UserActivity) error {
	result := ur.db.Create(&ua)

	return result.Error
}

func (ur UserActivityRepo) FindByID(id uuid.UUID) (ua model.UserActivity, err error) {
	result := ur.db.First(&ua, id)

	if result.RowsAffected == 0 {
		return ua, fmt.Errorf("activity not found")
	}

	if result.Error != nil {
		return ua, result.Error
	}

	return ua, nil
}
