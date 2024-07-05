package repository

import (
	"gorm.io/gorm"
)

type PlanRepo struct {
	db *gorm.DB
}

func (ur PlanRepo) Create() error {
	return nil
}
