package repository

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	entities2 "training-plan/internal/api/domain/activity/entities"
	"training-plan/internal/api/domain/plan/entities"
	"training-plan/internal/api/domain/user_activity"
	"training-plan/internal/api/infrastructure/database/model"
)

type UserActivityRepo struct {
	db *gorm.DB
}

func (ur UserActivityRepo) Update(act *useractivity.Entity) error {
	actModel := model.UserActivity{
		ID:        act.ID,
		UserID:    act.UserID,
		Type:      act.Type,
		Distance:  act.Distance,
		Pace:      act.Pace,
		Intervals: act.Intervals,
	}

	result := ur.db.Where("user_id = ?", act.UserID).Save(actModel)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("activity not found")
	}

	return nil
}

func (ur UserActivityRepo) Create(act *useractivity.Entity) error {
	actModel := model.UserActivity{
		ID:        act.ID,
		UserID:    act.UserID,
		Type:      act.Type,
		Distance:  act.Distance,
		Pace:      act.Pace,
		Intervals: act.Intervals,
	}

	result := ur.db.Create(&actModel)

	return result.Error
}

func (ur UserActivityRepo) FindByID(id uuid.UUID) (*useractivity.Entity, error) {
	ua := &model.UserActivity{}

	result := ur.db.First(&ua, id)

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("activity not found")
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return ua.ToDomainEntity(), nil
}

func (ur UserActivityRepo) GetFastestUserActivity(userID uuid.UUID) (stats useractivity.ActivityStats, err error) {
	err = ur.db.Table("user_activities").
		Select("type as Type, pace as Pace, user_id as UserID, distance as Distance").
		Where("user_id = ?", userID).
		Order("pace ASC").
		Row().
		Scan(&stats.Type, &stats.Pace, &stats.UserID, &stats.Distance)

	stats.Title = entities2.StatsTypeFastestUser

	return
}

func (ur UserActivityRepo) GetLongestUserActivity(userID uuid.UUID) (stats useractivity.ActivityStats, err error) {
	err = ur.db.Table("user_activities").
		Select("type as Type, pace as Pace, user_id as UserID, distance as Distance").
		Where("user_id = ?", userID).
		Order("distance DESC").
		Row().
		Scan(&stats.Type, &stats.Pace, &stats.UserID, &stats.Distance)

	stats.Title = entities2.StatsTypeLongestUser

	return
}

func (ur UserActivityRepo) GetFastestCommunityActivity() (stats useractivity.ActivityStats, err error) {
	err = ur.db.Table("user_activities").
		Select("type as Type, pace as Pace, user_id as UserID, distance as Distance").
		Order("pace ASC").
		Row().
		Scan(&stats.Type, &stats.Pace, &stats.UserID, &stats.Distance)

	stats.Title = entities2.StatsTypeFastestCommunity

	return
}

func (ur UserActivityRepo) GetLongestCommunityActivity() (stats useractivity.ActivityStats, err error) {
	err = ur.db.Table("user_activities").
		Select("type as Type, pace as Pace, user_id as UserID, distance as Distance").
		Order("distance DESC").
		Row().
		Scan(&stats.Type, &stats.Pace, &stats.UserID, &stats.Distance)

	stats.Title = entities2.StatsTypeLongestCommunity

	return
}

func (ur UserActivityRepo) GetMostCommonActivityType(userID uuid.UUID) (t entities.ActivityType, err error) {
	var mostCommon string

	err = ur.db.Table("user_activities").
		Select("type").
		Where("user_id = ?", userID).
		Group("type").
		Order("COUNT(type) DESC").
		Limit(1).
		Pluck("type", &mostCommon).Error

	if err == nil {
		t = entities.FromActivityType(mostCommon)
	}

	fmt.Println(t.Type)

	return
}
