package repository

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	model2 "training-plan/internal/api/domain/model"
	"training-plan/internal/api/domain/value_objects"
)

type UserActivityRepo struct {
	db *gorm.DB
}

func (ur UserActivityRepo) Update(act model2.UserActivity) error {
	result := ur.db.Where("user_id = ?", act.UserID).Save(act)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("activity not found")
	}

	return nil
}

func (ur UserActivityRepo) Create(ua model2.UserActivity) error {
	result := ur.db.Create(&ua)

	return result.Error
}

func (ur UserActivityRepo) FindByID(id uuid.UUID) (ua model2.UserActivity, err error) {
	result := ur.db.First(&ua, id)

	if result.RowsAffected == 0 {
		return ua, fmt.Errorf("activity not found")
	}

	if result.Error != nil {
		return ua, result.Error
	}

	return ua, nil
}

func (ur UserActivityRepo) GetFastestUserActivity(userID uuid.UUID) (stats model2.ActivityStats, err error) {
	err = ur.db.Table("user_activities").
		Select("type as Type, pace as Pace, user_id as UserID, distance as Distance").
		Where("user_id = ?", userID).
		Order("pace ASC").
		Row().
		Scan(&stats.Type, &stats.Pace, &stats.UserID, &stats.Distance)

	stats.Title = valueobjects.STATS_TYPE_FASTEST_USER

	return
}

func (ur UserActivityRepo) GetLongestUserActivity(userID uuid.UUID) (stats model2.ActivityStats, err error) {
	err = ur.db.Table("user_activities").
		Select("type as Type, pace as Pace, user_id as UserID, distance as Distance").
		Where("user_id = ?", userID).
		Order("distance DESC").
		Row().
		Scan(&stats.Type, &stats.Pace, &stats.UserID, &stats.Distance)

	stats.Title = valueobjects.STATS_TYPE_LONGEST_USER

	return
}

func (ur UserActivityRepo) GetFastestCommunityActivity() (stats model2.ActivityStats, err error) {
	err = ur.db.Table("user_activities").
		Select("type as Type, pace as Pace, user_id as UserID, distance as Distance").
		Order("pace ASC").
		Row().
		Scan(&stats.Type, &stats.Pace, &stats.UserID, &stats.Distance)

	stats.Title = valueobjects.STATS_TYPE_FASTEST_COMMUNITY

	return
}

func (ur UserActivityRepo) GetLongestCommunityActivity() (stats model2.ActivityStats, err error) {
	err = ur.db.Table("user_activities").
		Select("type as Type, pace as Pace, user_id as UserID, distance as Distance").
		Order("distance DESC").
		Row().
		Scan(&stats.Type, &stats.Pace, &stats.UserID, &stats.Distance)

	stats.Title = valueobjects.STATS_TYPE_LONGEST_COMMUNITY

	return
}

func (ur UserActivityRepo) GetMostCommonActivityType(userID uuid.UUID) (t valueobjects.ActivityType, err error) {
	var mostCommon string

	err = ur.db.Table("user_activities").
		Select("type").
		Where("user_id = ?", userID).
		Group("type").
		Order("COUNT(type) DESC").
		Limit(1).
		Pluck("type", &mostCommon).Error

	if err == nil {
		t = valueobjects.FromActivityType(mostCommon)
	}

	fmt.Println(t.Type)

	return
}
