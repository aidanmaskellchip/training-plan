package request

import "errors"

type CreateUserRequest struct {
	Username string `json:"username"`
}

func (c *CreateUserRequest) Validate() error {
	if c.Username == "" {
		return errors.New("username is required")
	}
	if len(c.Username) > 30 {
		return errors.New("username is too long")
	}

	return nil
}
