package query

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	repository2 "training-plan/internal/api/infrastructure/repository"
)

func TestFindRunningProfileQuery(t *testing.T) {
	t.Parallel()

	repos := repository2.NewMockRepos()

	tests := []struct {
		name string
		rpId string
		err  error
	}{
		{
			name: "Valid request",
			rpId: "99999991-8888-1111-9999-111111111111",
			err:  nil,
		},
		{
			name: "Valid request",
			rpId: repository2.MagicFailingRunningProfileId,
			err:  fmt.Errorf("running profile not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			_, err := FindRunningProfileQuery(&tt.rpId, repos)
			assert.Equal(t, tt.err, err)
		})
	}
}
