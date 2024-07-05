package repository

import (
	"gorm.io/gorm"
)

type RunningProfileRepo struct {
	db *gorm.DB
}

func (ur RunningProfileRepo) Create() error {
	return nil
}