package repository

import (
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func (ur UserRepo) Create() error {
	return nil
}