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

func (ur UserActivityRepoMock) GetFastestUserActivity(userID uuid.UUID) (stats model.ActivityStats, err error) {
	return
}

func (ur UserActivityRepoMock) GetLongestUserActivity(userID uuid.UUID) (stats model.ActivityStats, err error) {
	return
}

func (ur UserActivityRepoMock) GetFastestCommunityActivity() (stats model.ActivityStats, err error) {
	return
}

func (ur UserActivityRepoMock) GetLongestCommunityActivity() (stats model.ActivityStats, err error) {
	return
}
