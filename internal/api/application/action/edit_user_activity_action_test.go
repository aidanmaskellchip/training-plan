package action

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"training-plan/internal/api/domain/value_objects"
	"training-plan/internal/api/infrastructure/repository"
	"training-plan/internal/api/transport/request"
)

func TestEditUserActivityAction(t *testing.T) {
	t.Parallel()

	repos := repository.NewMockRepos()

	tests := []struct {
		name    string
		pace    float32
		dist    float32
		actType string
		actID   string
		err     error
	}{
		{
			name:    "Valid Upload",
			actType: valueobjects.EasyRun.Type,
			dist:    5.00,
			pace:    5.00,
			actID:   "9D8AC610-566D-4EF0-9C22-186B2A5ED793",
			err:     nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := request.EditUserActivityRequest{
				Type:     &tt.actType,
				Distance: &tt.dist,
				Pace:     &tt.pace,
			}

			err := EditUserActivityAction(tt.actID, &req, repos)
			assert.Equal(t, tt.err, err)
		})
	}
}
