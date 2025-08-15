package repository

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"training-plan/internal/api/domain/model"
)

type UserRepoMock struct {
	_ *gorm.DB
}

const MagicFailingUserId = "99999999-8888-1111-9999-111111111111"

func (ur UserRepoMock) Create(user model.User) (*model.User, error) {
	return &user, nil
}

func (ur UserRepoMock) FindByID(id uuid.UUID) (user model.User, err error) {
	if id == getMagicFailingID(MagicFailingUserId) {
		return user, fmt.Errorf("user not found")
	}

	return user, nil
}

func getMagicFailingID(id string) uuid.UUID {
	res, _ := uuid.Parse(id)

	return res
}
