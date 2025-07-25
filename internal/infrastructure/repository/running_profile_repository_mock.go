package repository

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
	"training-plan/internal/domain/model"
)

const MagicFailingRunningProfileId = "99999999-8888-1111-9999-111111111111"
const MagicFailingRunningProfileUserId = "11111111-1111-1111-1111-111111111111"

type RunningProfileRepoMock struct {
	db *gorm.DB
}

func (ur RunningProfileRepoMock) Create(_ model.RunningProfile) error {
	return nil
}

func (ur RunningProfileRepoMock) FindByID(id uuid.UUID) (rp model.RunningProfile, err error) {
	if id == getMagicFailingID(MagicFailingRunningProfileId) {
		return rp, fmt.Errorf("running profile not found")
	}

	return getRpMock(), nil
}

func (rpr RunningProfileRepoMock) FindByUserID(userID uuid.UUID) (rps []model.RunningProfile, err error) {
	if userID == getMagicFailingID(MagicFailingUserId) {
		return rps, fmt.Errorf("user not found")
	}

	return rps, nil
}

func (rpr RunningProfileRepoMock) FindLatestUserProfile(id uuid.UUID) (rp model.RunningProfile, err error) {
	if id == getMagicFailingID(MagicFailingUserId) || id == getMagicFailingID(MagicFailingRunningProfileUserId) {
		return rp, fmt.Errorf("running profile not found")
	}

	return getRpMock(), nil
}

func getRpMock() model.RunningProfile {
	rd, _ := json.Marshal([]int{1, 1, 1, 1, 1, 1, 1})

	return model.RunningProfile{
		ID:                  uuid.New(),
		UserID:              uuid.New(),
		GoalDistance:        "half-marathon",
		GoalTime:            1200,
		Terrain:             "road",
		Current5K:           0,
		Current10K:          0,
		CurrentHalfMarathon: 0,
		CurrentFullMarathon: 0,
		RunningDays:         rd,
		RunningDaysPerWeek:  3,
		LongRunDay:          6,
		CurrentAbility:      "intermediate",
		PlanLength:          12,
		StartDate:           time.Time{},
		GoalDate:            time.Time{},
	}
}
