package valueobjects

import "github.com/google/uuid"

type UUID struct {
	ID uuid.UUID
}

func NewUUID(id string) UUID {
	return UUID{
		ID: uuid.MustParse(id),
	}
}
