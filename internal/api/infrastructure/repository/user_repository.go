package repository

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"training-plan/internal/api/domain/user"
	"training-plan/internal/api/infrastructure/database/model"
)

type UserRepo struct {
	db *gorm.DB
}

func (ur UserRepo) Create(user *user.Entity) (*user.Entity, error) {
	userModel := model.User{
		Username: user.Username,
	}

	result := ur.db.Create(&userModel)

	return userModel.ToDomainEntity(), result.Error
}

func (ur UserRepo) FindByID(id uuid.UUID) (*user.Entity, error) {
	userModel := model.User{}

	result := ur.db.First(userModel, id)

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("user not found")
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return userModel.ToDomainEntity(), nil
}
