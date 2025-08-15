package valueobjects

import "github.com/google/uuid"

type UserID struct {
	ID uuid.UUID
}

func NewUserID(id string) UserID {
	return UserID{
		ID: uuid.MustParse(id),
	}
}
