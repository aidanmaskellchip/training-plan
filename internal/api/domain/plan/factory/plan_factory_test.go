package factory

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	model2 "training-plan/internal/api/domain/model"
	vo "training-plan/internal/api/domain/plan/entities"
)

func TestNewPlan(t *testing.T) {
	t.Parallel()

	rd, _ := json.Marshal([]int{1, 1, 1, 1, 1, 1, 1})

	tests := []struct {
		name           string
		runningProfile model2.RunningProfile
		expectedErr    error
	}{
		{
			name: "Successful plan creation",
			runningProfile: model2.RunningProfile{
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
			},
			expectedErr: nil,
		},
		{
			name: "Invalid running days JSON",
			runningProfile: model2.RunningProfile{
				ID:                 uuid.New(),
				UserID:             uuid.New(),
				GoalDistance:       "half-marathon",
				RunningDays:        []byte("invalid json"),
				RunningDaysPerWeek: 3,
				LongRunDay:         6,
				CurrentAbility:     "intermediate",
				PlanLength:         12,
			},
			expectedErr: fmt.Errorf("invalid character 'i' looking for beginning of value"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewPlan(tt.runningProfile)
			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestSetLongRuns(t *testing.T) {
	t.Parallel()

	weeks := make([]model2.ActivityWeek, 12)
	for i := range weeks {
		weeks[i] = model2.ActivityWeek{}
	}

	err := setLongRuns(&weeks, 6, 12)
	assert.NoError(t, err)

	// Test with an invalid planLength
	err = setLongRuns(&weeks, 6, 99)
	assert.Error(t, err)
}

func TestSetEasyRuns(t *testing.T) {
	t.Parallel()

	weeks := make([]model2.ActivityWeek, 12)
	for i := range weeks {
		weeks[i] = model2.ActivityWeek{}
	}

	rd, _ := vo.RunningDaysFromJson([]byte("[1,1,1,1,1,1,1]"))
	err := setEasyRuns(&weeks, rd, 6, 12)
	assert.NoError(t, err)

	// Test with an invalid planLength
	err = setEasyRuns(&weeks, rd, 6, 99)
	assert.Error(t, err)
}

func TestSetThresholdRuns(t *testing.T) {
	t.Parallel()

	weeks := make([]model2.ActivityWeek, 12)
	setThresholdRuns(&weeks)

	// Add assertions to check if threshold runs are set correctly
}
