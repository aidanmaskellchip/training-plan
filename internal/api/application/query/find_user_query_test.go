package query

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	repository2 "training-plan/internal/api/infrastructure/repository"
)

func TestFindUserQuery(t *testing.T) {
	t.Parallel()

	repos := repository2.NewMockRepos()

	tests := []struct {
		name   string
		userId string
		err    error
	}{
		{
			name:   "Valid request",
			userId: "99999991-8888-1111-9999-111111111111",
			err:    nil,
		},
		{
			name:   "Valid request",
			userId: repository2.MagicFailingUserId,
			err:    fmt.Errorf("user not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			_, err := FindUserQuery(&tt.userId, repos)
			assert.Equal(t, tt.err, err)
		})
	}
}
