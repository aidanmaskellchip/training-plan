package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"training-plan/internal/domain/model"
)

type UserRepoMock struct {
	db *gorm.DB
}

func (ur UserRepoMock) Create(user model.User) error {
	return nil
}

func (ur UserRepoMock) FindByID(id uuid.UUID) (user model.User, err error) {
	return user, nil
}
