package repository

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
	"training-plan/internal/api/domain/user"
)

type UserRepoMock struct {
	_ *gorm.DB
}

const MagicFailingUserId = "99999999-8888-1111-9999-111111111111"

func (ur UserRepoMock) Create(ent *user.Entity) (*user.Entity, error) {
	return &user.Entity{
		ID:       uuid.New(),
		Username: ent.Username,
	}, nil
}

func (ur UserRepoMock) FindByID(id uuid.UUID) (entity *user.Entity, err error) {
	if id == getMagicFailingID(MagicFailingUserId) {
		return &user.Entity{
			ID:        id,
			CreatedAt: nil,
		}, fmt.Errorf("user not found")
	}

	now := time.Now()
	return &user.Entity{
		ID:        id,
		Username:  "test-user",
		CreatedAt: &now,
	}, nil
}

func getMagicFailingID(id string) uuid.UUID {
	res, _ := uuid.Parse(id)

	return res
}
