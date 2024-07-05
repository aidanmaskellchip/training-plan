package repository

import (
	"gorm.io/gorm"
)

type UserRepoMock struct {
	db *gorm.DB
}

func (ur UserRepoMock) Create() error {
	return nil
}
