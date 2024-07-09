package repository

import (
	"gorm.io/gorm"
	"training-plan/internal/data/model"
)

type UserRepo struct {
	db *gorm.DB
}

func (ur UserRepo) Create(user model.User) error {
	result := ur.db.Create(&user)

	return result.Error
}
