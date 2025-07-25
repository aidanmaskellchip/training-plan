
package activitytypefactory

import (
	"github.com/stretchr/testify/assert"
	"testing"
	vo "training-plan/internal/domain/value_objects"
)

func TestNewActivityType(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		activityType string
		expectedErr  string
	}{
		{
			name:         "Valid activity type",
			activityType: "easy_run",
			expectedErr:  "",
		},
		{
			name:         "Invalid activity type",
			activityType: "invalid",
			expectedErr:  "invalid activity type",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			activityType, err := NewActivityType(tt.activityType)

			if tt.expectedErr != "" {
				assert.EqualError(t, err, tt.expectedErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, vo.ActivityType{Type: tt.activityType}, activityType)
			}
		})
	}
}
