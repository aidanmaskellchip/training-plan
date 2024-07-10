package request

import (
	"errors"
)

type FindUserRequest struct {
	UserID string `json:"user_id"`
}

func (c *FindUserRequest) Validate() error {
	if c.UserID == "" {
		return errors.New("user id is required")
	}

	return nil
}
