package request

import (
	"errors"
)

type CreatePlanRequest struct {
	UserID string `json:"user_id"`
}

func (c *CreatePlanRequest) Validate() error {
	if c.UserID == "" {
		return errors.New("user id is required")
	}

	return nil
}
