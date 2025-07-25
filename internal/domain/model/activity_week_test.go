package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
	vo "training-plan/internal/domain/value_objects"
)

func TestGetDayByIndex(t *testing.T) {
	t.Parallel()

	aw := ActivityWeek{
		Mon: vo.Activity{Distance: 1.0},
		Tue: vo.Activity{Distance: 2.0},
		Wed: vo.Activity{Distance: 3.0},
		Thu: vo.Activity{Distance: 4.0},
		Fri: vo.Activity{Distance: 5.0},
		Sat: vo.Activity{Distance: 6.0},
		Sun: vo.Activity{Distance: 7.0},
	}

	tests := []struct {
		name      string
		index     int
		expected  *vo.Activity
		expectErr bool
	}{
		{"Monday", 0, &aw.Mon, false},
		{"Tuesday", 1, &aw.Tue, false},
		{"Wednesday", 2, &aw.Wed, false},
		{"Thursday", 3, &aw.Thu, false},
		{"Friday", 4, &aw.Fri, false},
		{"Saturday", 5, &aw.Sat, false},
		{"Sunday", 6, &aw.Sun, false},
		{"Invalid Index", 7, nil, true},
		{"Negative Index", -1, nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			activity, err := aw.GetDayByIndex(tt.index)
			if tt.expectErr {
				assert.Error(t, err)
				assert.Nil(t, activity)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, activity)
			}
		})
	}
}

func TestGetEasyRunDay(t *testing.T) {
	t.Parallel()

	aw := ActivityWeek{} // The content of ActivityWeek doesn't matter for this test

	tests := []struct {
		name        string
		runningDays vo.RunningDays
		longRunDay  int
		expectedDay int
		expectErr   bool
	}{
		{
			name:        "Easy run day found (Tuesday)",
			runningDays: vo.RunningDays{Days: []int{1, 1, 1, 1, 1, 1, 1}}, // All running days
			longRunDay:  0, // Monday is long run
			expectedDay: 1,
			expectErr:   false,
		},
		{
			name:        "No easy run day found (all unavailable)",
			runningDays: vo.RunningDays{Days: []int{0, 0, 0, 0, 0, 0, 0}}, // No running days
			longRunDay:  -1, // No specific long run day
			expectedDay: 0,
			expectErr:   true,
		},
		{
			name:        "No easy run day found (only long run day is available)",
			runningDays: vo.RunningDays{Days: []int{0, 0, 0, 1, 0, 0, 0}}, // Only Thursday is a running day
			longRunDay:  3, // Thursday is the long run day
			expectedDay: 0,
			expectErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			day, err := aw.GetEasyRunDay(tt.runningDays, tt.longRunDay)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedDay, day)
			}
		})
	}
}
