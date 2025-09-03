package factory

import (
	"errors"
	vo "training-plan/internal/api/domain/plan/entities"
)

func NewActivityType(at string) (vo.ActivityType, error) {
	found := false
	for _, a := range vo.GetActivityTypeStrings() {
		if a == at {
			found = true
			break
		}
	}

	if !found {
		return vo.ActivityType{}, errors.New("invalid activity type")
	}

	return vo.ActivityType{Type: at}, nil
}
