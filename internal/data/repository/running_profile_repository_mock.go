package repository

import (
	"gorm.io/gorm"
)

type RunningProfileRepoMock struct {
	db *gorm.DB
}

func (ur RunningProfileRepoMock) Create() error {
	return nil
}
