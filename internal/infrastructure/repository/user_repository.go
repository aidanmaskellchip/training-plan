package repository

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"training-plan/internal/domain/model"
)

type UserRepo struct {
	db *gorm.DB
}

func (ur UserRepo) Create(user model.User) (*model.User, error) {
	result := ur.db.Create(&user)

	return &user, result.Error
}

func (ur UserRepo) FindByID(id uuid.UUID) (user model.User, err error) {
	result := ur.db.First(&user, id)

	if result.RowsAffected == 0 {
		return user, fmt.Errorf("user not found")
	}

	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}
