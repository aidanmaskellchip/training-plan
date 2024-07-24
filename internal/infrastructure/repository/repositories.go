package repository

import (
	"errors"
	"github.com/google/uuid"
	model2 "training-plan/internal/domain/model"

	"gorm.io/gorm"
)

var ErrNoRecord = errors.New("no matching record found")

type UserRepository interface {
	Create(user model2.User) error
	FindByID(id uuid.UUID) (user model2.User, err error)
}
type RunningProfileRepository interface {
	Create(profile model2.RunningProfile) error
}
type PlanRepository interface {
	Create(plan model2.Plan) error
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

func NewMockRepos() *Repositories {
	return &Repositories{
		UserRepository:           UserRepoMock{},
		RunningProfileRepository: RunningProfileRepoMock{},
		PlanRepository:           PlanRepoMock{},
	}
}
