package repository

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	runningprofile "training-plan/internal/api/domain/running_profile"
	"training-plan/internal/api/infrastructure/database/model"
)

type RunningProfileRepo struct {
	db *gorm.DB
}

func (rpr RunningProfileRepo) Create(profile *runningprofile.Entity) error {
	runningProfile := model.RunningProfile{
		ID:                  profile.ID,
		UserID:              profile.UserID,
		GoalDistance:        profile.GoalDistance,
		GoalTime:            profile.GoalTime,
		Terrain:             profile.Terrain,
		Current5K:           *profile.Current5K,
		Current10K:          *profile.Current10K,
		CurrentHalfMarathon: *profile.CurrentHalfMarathon,
		CurrentFullMarathon: *profile.CurrentFullMarathon,
		RunningDays:         profile.RunningDays,
		RunningDaysPerWeek:  profile.RunningDaysPerWeek,
		LongRunDay:          profile.LongRunDay,
		CurrentAbility:      profile.CurrentAbility,
		PlanLength:          profile.PlanLength,
		StartDate:           *profile.StartDate,
		GoalDate:            *profile.GoalDate,
	}

	result := rpr.db.Create(&runningProfile)

	return result.Error
}

func (rpr RunningProfileRepo) FindByID(id uuid.UUID) (*runningprofile.Entity, error) {
	rp := &model.RunningProfile{}

	result := rpr.db.First(rp, id)

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("running profile not found")
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return rp.ToDomainEntity(), nil
}

func (rpr RunningProfileRepo) FindByUserID(userID uuid.UUID) (ents []runningprofile.Entity, err error) {
	var rps []model.RunningProfile

	result := rpr.db.Where("user_id = ?", userID).Find(&rps)

	if result.RowsAffected == 0 {
		return ents, fmt.Errorf("running profiles not found")
	}

	if result.Error != nil {
		return ents, result.Error
	}

	for _, rp := range rps {
		ents = append(ents, *rp.ToDomainEntity())
	}

	return ents, nil
}

func (rpr RunningProfileRepo) FindLatestUserProfile(id uuid.UUID) (*runningprofile.Entity, error) {
	rp := model.RunningProfile{}

	result := rpr.db.Order("created_at desc").Where("user_id = ?", id).Find(&rp)

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("running profile not found")
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return rp.ToDomainEntity(), nil
}
