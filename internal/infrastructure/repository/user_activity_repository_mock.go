package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"training-plan/internal/domain/model"
	valueobjects "training-plan/internal/domain/value_objects"
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

func (ur UserActivityRepoMock) GetFastestUserActivity(userID uuid.UUID) (stats valueobjects.ActivityStats, err error) {
	return
}

func (ur UserActivityRepoMock) GetLongestUserActivity(userID uuid.UUID) (stats valueobjects.ActivityStats, err error) {
	return
}

func (ur UserActivityRepoMock) GetFastestCommunityActivity() (stats valueobjects.ActivityStats, err error) {
	return
}

func (ur UserActivityRepoMock) GetLongestCommunityActivity() (stats valueobjects.ActivityStats, err error) {
	return
}
