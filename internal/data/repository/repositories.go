package repository

import (
	"errors"

	"gorm.io/gorm"
)

var ErrNoRecord = errors.New("no matching record found")

type Repositories struct {
	User interface {
		Create() error
	}
	RunningProfile interface {
		Create() error
	}
	Plan interface {
		Create() error
	}
}

func NewRepos(db *gorm.DB) *Repositories {
	return &Repositories{
		User: UserRepo{
			db: db,
		},
		RunningProfile: RunningProfileRepo{
			db: db,
		},
		Plan: PlanRepo{
			db: db,
		},
	}
}

func NewMockRepos() Repositories {
	return Repositories{
		User:           UserRepoMock{},
		RunningProfile: RunningProfileRepoMock{},
		Plan:           PlanRepoMock{},
	}
}
