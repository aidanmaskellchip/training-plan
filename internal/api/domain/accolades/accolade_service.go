package accolades

import (
	"training-plan/internal/api/domain/model"
)

type AccoladeService interface {
	GetUserAccolades(userID string) ([]model.Accolade, error)
}
