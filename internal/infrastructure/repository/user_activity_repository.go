package repository

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"training-plan/internal/domain/model"
	vo "training-plan/internal/domain/value_objects"
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

func (ur UserActivityRepo) GetFastestUserActivity(userID uuid.UUID) (stats model.ActivityStats, err error) {
	err = ur.db.Table("user_activities").
		Select("pace as Pace, user_id as UserID, distance as Distance").
		Where("user_id = ?", userID).
		Order("pace ASC").
		Row().
		Scan(&stats.Pace, &stats.UserID, &stats.Distance)

	stats.Type = vo.STATS_TYPE_FASTEST_USER

	return
}

func (ur UserActivityRepo) GetLongestUserActivity(userID uuid.UUID) (stats model.ActivityStats, err error) {
	err = ur.db.Table("user_activities").
		Select("pace as Pace, user_id as UserID, distance as Distance").
		Where("user_id = ?", userID).
		Order("distance DESC").
		Row().
		Scan(&stats.Pace, &stats.UserID, &stats.Distance)

	stats.Type = vo.STATS_TYPE_LONGEST_USER

	return
}

func (ur UserActivityRepo) GetFastestCommunityActivity() (stats model.ActivityStats, err error) {
	err = ur.db.Table("user_activities").
		Select("pace as Pace, user_id as UserID, distance as Distance").
		Order("pace ASC").
		Row().
		Scan(&stats.Pace, &stats.UserID, &stats.Distance)

	stats.Type = vo.STATS_TYPE_FASTEST_COMMUNITY

	return
}

func (ur UserActivityRepo) GetLongestCommunityActivity() (stats model.ActivityStats, err error) {
	err = ur.db.Table("user_activities").
		Select("pace as Pace, user_id as UserID, distance as Distance").
		Order("distance DESC").
		Row().
		Scan(&stats.Pace, &stats.UserID, &stats.Distance)

	stats.Type = vo.STATS_TYPE_LONGEST_COMMUNITY

	return
}
