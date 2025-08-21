package action

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	repository2 "training-plan/internal/api/infrastructure/repository"
	"training-plan/internal/api/transport/request"
)

func TestCreatePlanAction(t *testing.T) {
	t.Parallel()

	repos := repository2.NewMockRepos()

	tests := []struct {
		name    string
		request request.CreatePlanRequest
		err     error
	}{
		{
			name: "Valid request",
			request: request.CreatePlanRequest{
				UserID: "99999991-8888-1111-9999-111111111111",
			},
			err: nil,
		},
		{
			name: "User not found",
			request: request.CreatePlanRequest{
				UserID: repository2.MagicFailingUserId,
			},
			err: fmt.Errorf("user not found"),
		},
		{
			name: "Running profile not found",
			request: request.CreatePlanRequest{
				UserID: repository2.MagicFailingRunningProfileUserId,
			},
			err: fmt.Errorf("running profile not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			_, err := CreatePlanAction(&tt.request, repos)
			assert.Equal(t, tt.err, err)
		})
	}
}
