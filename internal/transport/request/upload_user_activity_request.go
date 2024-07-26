package request

import (
	"errors"
	"github.com/google/uuid"
	valueobjects "training-plan/internal/domain/value_objects"
)

type UploadUserActivityRequest struct {
	UserID   uuid.UUID `json:"user_id"`
	Type     string    `json:"type"`
	Distance float32   `json:"distance"`
	Pace     float32   `json:"pace"`
	//Todo: add intervals
}

func (u *UploadUserActivityRequest) Validate() error {
	err := errors.New("invalid activity type: " + u.Type)
	for _, v := range valueobjects.GetTypeStrings() {
		if v == u.Type {
			err = nil
		}
	}

	if u.Distance == 0 {
		return errors.New("distance is invalid")
	}
	if u.Pace == 0 {
		return errors.New("pace is invalid")
	}

	return err
}
