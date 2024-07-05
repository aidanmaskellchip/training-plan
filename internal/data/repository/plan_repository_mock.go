package repository

import (
	"gorm.io/gorm"
)

type PlanRepoMock struct {
	db *gorm.DB
}

func (ur PlanRepoMock) Create() error {
	return nil
}
