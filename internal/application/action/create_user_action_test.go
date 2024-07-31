package action

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"training-plan/internal/infrastructure/repository"
	"training-plan/internal/transport/request"
)

func TestCreateUserAction(t *testing.T) {
	t.Parallel()

	repos := repository.NewMockRepos()

	tests := []struct {
		name    string
		request request.CreateUserRequest
		err     error
	}{
		{
			name:    "Valid User",
			request: request.CreateUserRequest{Username: "ValidUser"},
			err:     nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CreateUserAction(&tt.request, repos)
			assert.Equal(t, tt.err, err)
		})
	}
}
