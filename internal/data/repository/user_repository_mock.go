package repository

import (
	"gorm.io/gorm"
	"training-plan/internal/data/model"
)

type UserRepoMock struct {
	db *gorm.DB
}

func (ur UserRepoMock) Create(user model.User) error {
	return nil
}

func (ur UserRepoMock) FindByID(id string) (user model.User, err error) {
	return user, nil
}
