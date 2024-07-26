package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"training-plan/internal/domain/model"
)

type UserActivityRepoMock struct {
	db *gorm.DB
}

func (ur UserActivityRepoMock) Create(ua model.UserActivity) error {
	return nil
}

func (ur UserActivityRepoMock) FindByID(id uuid.UUID) (ua model.UserActivity, err error) {
	return ua, nil
}
