package repository

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	vo "training-plan/internal/api/domain/plan/entities"
	"training-plan/internal/api/domain/user_activity"
)

type UserActivityRepoMock struct {
	_ *gorm.DB
}

func (ur UserActivityRepoMock) Create(entity *useractivity.Entity) error {
	return nil
}

func (ur UserActivityRepoMock) FindByID(id uuid.UUID) (ua *useractivity.Entity, err error) {
	return &useractivity.Entity{
		ID:       id,
		UserID:   getMagicFailingID("11111111-2222-3333-4444-555555555555"),
		Type:     "EasyRun",
		Distance: 5.0,
		Pace:     5.0,
	}, nil
}

func (ur UserActivityRepoMock) GetFastestUserActivity(userID uuid.UUID) (stats useractivity.ActivityStats, err error) {
	if userID == getMagicFailingID(MagicFailingUserId) {
		return stats, fmt.Errorf("error")
	}
	return
}

func (ur UserActivityRepoMock) GetLongestUserActivity(userID uuid.UUID) (stats useractivity.ActivityStats, err error) {
	if userID == getMagicFailingID(MagicFailingUserId) {
		return stats, fmt.Errorf("error")
	}
	return
}

func (ur UserActivityRepoMock) GetFastestCommunityActivity() (stats useractivity.ActivityStats, err error) {
	return
}

func (ur UserActivityRepoMock) GetLongestCommunityActivity() (stats useractivity.ActivityStats, err error) {
	return
}

func (ur UserActivityRepoMock) GetMostCommonActivityType(userID uuid.UUID) (t vo.ActivityType, err error) {
	if userID == getMagicFailingID(MagicFailingUserId) {
		return t, fmt.Errorf("error")
	}
	return
}

func (ur UserActivityRepoMock) Update(*useractivity.Entity) error {
	return nil
}
