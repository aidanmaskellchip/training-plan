package action

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"training-plan/internal/api/infrastructure/repository"
	"training-plan/internal/api/transport/request"
)

func TestCreateRunningProfileAction(t *testing.T) {
	t.Parallel()

	repos := repository.NewMockRepos()

	tests := []struct {
		name    string
		request request.CreateRunningProfileRequest
		err     error
	}{
		{
			name: "Valid Running Profile",
			request: request.CreateRunningProfileRequest{
				UserID:              uuid.New(),
				GoalDistance:        "500",
				GoalTime:            30.00,
				Terrain:             "road",
				Current5K:           35.00,
				Current10K:          0,
				CurrentHalfMarathon: 0,
				CurrentFullMarathon: 0,
				RunningDays:         []int{0, 1, 1, 0, 1, 1, 1},
				RunningDaysPerWeek:  3,
				LongRunDay:          6,
				CurrentAbility:      "intermediate",
				PlanLength:          12,
				StartDate:           "2024-07-10",
				GoalDate:            "2024-12-10",
			},
			err: nil,
		},
		{
			name: "Invalid long run day",
			request: request.CreateRunningProfileRequest{
				UserID:              uuid.New(),
				GoalDistance:        "500",
				GoalTime:            30.00,
				Terrain:             "road",
				Current5K:           35.00,
				Current10K:          0,
				CurrentHalfMarathon: 0,
				CurrentFullMarathon: 0,
				RunningDays:         []int{0, 1, 1, 0, 1, 1, 1},
				RunningDaysPerWeek:  3,
				LongRunDay:          0,
				CurrentAbility:      "intermediate",
				PlanLength:          12,
				StartDate:           "2024-07-10",
				GoalDate:            "2024-12-10",
			},
			err: errors.New("long run day not able to be allocated from running days"),
		},
		{
			name: "Invalid running days per week",
			request: request.CreateRunningProfileRequest{
				UserID:              uuid.New(),
				GoalDistance:        "500",
				GoalTime:            30.00,
				Terrain:             "road",
				Current5K:           35.00,
				Current10K:          0,
				CurrentHalfMarathon: 0,
				CurrentFullMarathon: 0,
				RunningDays:         []int{0, 1, 1, 1, 1, 1, 1},
				RunningDaysPerWeek:  6,
				LongRunDay:          6,
				CurrentAbility:      "intermediate",
				PlanLength:          12,
				StartDate:           "2024-07-10",
				GoalDate:            "2024-12-10",
			},
			err: errors.New("invalid number of running days selected"),
		},
		{
			name: "Invalid running days",
			request: request.CreateRunningProfileRequest{
				UserID:              uuid.New(),
				GoalDistance:        "500",
				GoalTime:            30.00,
				Terrain:             "road",
				Current5K:           35.00,
				Current10K:          0,
				CurrentHalfMarathon: 0,
				CurrentFullMarathon: 0,
				RunningDays:         []int{0, 1, 1, 0, 1, 0, 1},
				RunningDaysPerWeek:  5,
				LongRunDay:          6,
				CurrentAbility:      "intermediate",
				PlanLength:          12,
				StartDate:           "2024-07-10",
				GoalDate:            "2024-12-10",
			},
			err: errors.New("not enough running days selected to allocate all runs"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CreateRunningProfileAction(&tt.request, repos)
			assert.Equal(t, tt.err, err)
		})
	}
}
