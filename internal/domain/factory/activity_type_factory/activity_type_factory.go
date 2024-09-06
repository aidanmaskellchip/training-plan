package activitytypefactory

import (
	"errors"
	vo "training-plan/internal/domain/value_objects"
)

func NewActivityType(at string) (vo.ActivityType, error) {
	err := errors.New("invalid activity type")
	for _, a := range vo.GetTypeStrings() {
		if a == at {
			err = nil
		}
	}

	if err != nil {
		return vo.ActivityType{}, err
	}

	return vo.ActivityType{Type: at}, nil
}
