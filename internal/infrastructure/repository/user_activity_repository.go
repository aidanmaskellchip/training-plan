package repository

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"training-plan/internal/domain/model"
	valueobjects "training-plan/internal/domain/value_objects"
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

func (ur UserActivityRepo) GetFastestUserActivity(userID uuid.UUID) (stats valueobjects.ActivityStats, err error) {
	err = ur.db.Table("user_activities").
		Where("user_id", "=", userID).
		Select("min(pace) as Pace, user_id as UserID, distance as Distance").
		Row().
		Scan(&stats)

	return
}

func (ur UserActivityRepo) GetLongestUserActivity(userID uuid.UUID) (stats valueobjects.ActivityStats, err error) {
	err = ur.db.Table("user_activities").
		Where("user_id", "=", userID).
		Select("max(distance) as Distance, user_id as UserID, pace as Pace").
		Row().
		Scan(&stats)

	return
}

func (ur UserActivityRepo) GetFastestCommunityActivity() (stats valueobjects.ActivityStats, err error) {
	err = ur.db.Table("user_activities").
		Select("min(pace) as Pace, user_id as UserID, distance as Distance").
		Row().
		Scan(&stats)

	return
}

func (ur UserActivityRepo) GetLongestCommunityActivity() (stats valueobjects.ActivityStats, err error) {
	err = ur.db.Table("user_activities").
		Select("max(distance) as Distance, user_id as UserID, pace as Pace").
		Row().
		Scan(&stats)

	return
}
