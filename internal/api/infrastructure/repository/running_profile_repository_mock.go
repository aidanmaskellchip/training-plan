package repository

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	runningprofile "training-plan/internal/api/domain/running_profile"
)

const MagicFailingRunningProfileId = "99999999-8888-1111-9999-111111111111"
const MagicFailingRunningProfileUserId = "11111111-1111-1111-1111-111111111111"

type RunningProfileRepoMock struct {
	_ *gorm.DB
}

func (ur RunningProfileRepoMock) Create(_ *runningprofile.Entity) error {
	return nil
}

func (ur RunningProfileRepoMock) FindByID(id uuid.UUID) (rp *runningprofile.Entity, err error) {
	if id == getMagicFailingID(MagicFailingRunningProfileId) {
		return rp, fmt.Errorf("running profile not found")
	}

	return getRpMock(), nil
}

func (ur RunningProfileRepoMock) FindByUserID(userID uuid.UUID) (rps []runningprofile.Entity, err error) {
	if userID == getMagicFailingID(MagicFailingUserId) {
		return rps, fmt.Errorf("user not found")
	}

	return rps, nil
}

func (ur RunningProfileRepoMock) FindLatestUserProfile(id uuid.UUID) (rp *runningprofile.Entity, err error) {
	if id == getMagicFailingID(MagicFailingUserId) || id == getMagicFailingID(MagicFailingRunningProfileUserId) {
		return rp, fmt.Errorf("running profile not found")
	}

	return getRpMock(), nil
}

func getRpMock() *runningprofile.Entity {
	rd, _ := json.Marshal([]int{1, 1, 1, 1, 1, 1, 1})

	return &runningprofile.Entity{
		ID:                  uuid.New(),
		UserID:              uuid.New(),
		GoalDistance:        "half-marathon",
		GoalTime:            1200,
		Terrain:             "road",
		Current5K:           nil,
		Current10K:          nil,
		CurrentHalfMarathon: nil,
		CurrentFullMarathon: nil,
		RunningDays:         rd,
		RunningDaysPerWeek:  3,
		LongRunDay:          6,
		CurrentAbility:      "intermediate",
		PlanLength:          12,
		StartDate:           nil,
		GoalDate:            nil,
	}
}
