package action

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	valueobjects "training-plan/internal/domain/value_objects"
	"training-plan/internal/infrastructure/event"
	"training-plan/internal/infrastructure/repository"
	"training-plan/internal/transport/request"
)

func TestUploadUserActivityAction(t *testing.T) {
	t.Parallel()

	repos := repository.NewMockRepos()

	tests := []struct {
		name    string
		request request.UploadUserActivityRequest
		err     error
	}{
		{
			name: "Valid Upload",
			request: request.UploadUserActivityRequest{
				UserID:   uuid.New(),
				Type:     valueobjects.EasyRun.Type,
				Distance: 5.00,
				Pace:     5.00,
			},
			err: nil,
		},
		{
			name: "Invalid type",
			request: request.UploadUserActivityRequest{
				UserID:   uuid.New(),
				Type:     "invalid run type",
				Distance: 5.00,
				Pace:     5.00,
			},
			err: errors.New("invalid activity type: invalid run type"),
		},
		{
			name: "Invalid distance",
			request: request.UploadUserActivityRequest{
				UserID:   uuid.New(),
				Type:     valueobjects.EasyRun.Type,
				Distance: 0,
				Pace:     5.00,
			},
			err: errors.New("distance is invalid"),
		},
		{
			name: "Invalid pace",
			request: request.UploadUserActivityRequest{
				UserID:   uuid.New(),
				Type:     valueobjects.EasyRun.Type,
				Distance: 5.00,
				Pace:     0,
			},
			err: errors.New("pace is invalid"),
		},
		{
			name: "User not found",
			request: request.UploadUserActivityRequest{
				Type:     valueobjects.EasyRun.Type,
				Distance: 5.00,
				Pace:     1.00,
			},
			err: errors.New("user not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.request.UserID == uuid.Nil {
				id, _ := uuid.Parse(repository.MagicFailingUserId)
				tt.request.UserID = id
			}

			err := UploadUserActivityAction(&tt.request, repos, event.NewEventBus(), event.NewChannels())
			assert.Equal(t, tt.err, err)
		})
	}
}
