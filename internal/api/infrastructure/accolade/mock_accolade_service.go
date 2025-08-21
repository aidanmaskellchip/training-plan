package accolade

import (
	"fmt"
	"training-plan/internal/api/domain/model"
)

const MagicFailingUserID = "00000000-0000-0000-0000-000000000001"

type MockAccoladeService struct{}

func (m *MockAccoladeService) GetUserAccolades(userID string) ([]model.Accolade, error) {
	if userID == MagicFailingUserID {
		return nil, fmt.Errorf("failed to retrieve accolades for user %s", userID)
	}

	return []model.Accolade{
		{Title: "5 Goals Created", Description: "Created 5 training goals"},
		{Title: "Sub 20 Minute 5K", Description: "Completed a 5K in under 20 minutes"},
	}, nil
}
