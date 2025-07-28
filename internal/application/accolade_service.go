package application

import (
	"training-plan/internal/domain/model"
)

type AccoladeService interface {
	GetUserAccolades(userID string) ([]model.Accolade, error)
}
