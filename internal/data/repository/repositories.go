package repository

import (
	"errors"
	"training-plan/internal/data/model"

	"gorm.io/gorm"
)

var ErrNoRecord = errors.New("no matching record found")

type UserRepository interface {
	Create(user model.User) error
}
type RunningProfileRepository interface {
	Create(profile model.RunningProfile) error
}
type PlanRepository interface {
	Create(plan model.Plan) error
}

type Repositories struct {
	UserRepository
	RunningProfileRepository
	PlanRepository
}

func NewRepos(db *gorm.DB) *Repositories {
	return &Repositories{
		UserRepository: UserRepo{
			db: db,
		},
		RunningProfileRepository: RunningProfileRepo{
			db: db,
		},
		PlanRepository: PlanRepo{
			db: db,
		},
	}
}

func NewMockRepos() Repositories {
	return Repositories{
		UserRepository:           UserRepoMock{},
		RunningProfileRepository: RunningProfileRepoMock{},
		PlanRepository:           PlanRepoMock{},
	}
}
