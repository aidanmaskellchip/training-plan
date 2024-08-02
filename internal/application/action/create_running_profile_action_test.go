package action

//
//import (
//	"github.com/stretchr/testify/assert"
//	"testing"
//	"training-plan/internal/infrastructure/repository"
//	"training-plan/internal/transport/request"
//)
//
//func TestCreateRunningProfileAction(t *testing.T) {
//	t.Parallel()
//
//	repos := repository.NewMockRepos()
//
//	tests := []struct {
//		name    string
//		request request.CreateRunningProfileRequest
//		err     error
//	}{
//		{
//			name:    "Valid Running Profile",
//			request: request.CreateRunningProfileRequest{},
//			err:     nil,
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			err := CreateRunningProfileAction(&tt.request, repos)
//			assert.Equal(t, tt.err, err)
//		})
//	}
//}
