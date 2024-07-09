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
