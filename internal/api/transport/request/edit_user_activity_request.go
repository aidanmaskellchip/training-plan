package request

import (
	"errors"
	"training-plan/internal/api/domain/value_objects"
)

type EditUserActivityRequest struct {
	Type     *string  `json:"type"`
	Distance *float32 `json:"distance"`
	Pace     *float32 `json:"pace"`
}

func (u *EditUserActivityRequest) Validate() error {
	err := errors.New("invalid activity type: " + *u.Type)

	if u.Type != nil {
		for _, v := range valueobjects.GetActivityTypeStrings() {
			if v == *u.Type {
				err = nil
			}
		}
	}

	if *u.Distance == 0 {
		return errors.New("distance is invalid")
	}
	if *u.Pace == 0 {
		return errors.New("pace is invalid")
	}

	return err
}
