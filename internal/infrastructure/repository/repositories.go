package repository

import (
	"errors"
	"github.com/google/uuid"
	"training-plan/internal/domain/model"

	"gorm.io/gorm"
)

var ErrNoRecord = errors.New("no matching record found")

type UserRepository interface {
	Create(user model.User) error
	FindByID(id uuid.UUID) (user model.User, err error)
}
type UserActivityRepository interface {
	Create(ua model.UserActivity) error
	FindByID(id uuid.UUID) (ua model.UserActivity, err error)
}
type RunningProfileRepository interface {
	Create(profile model.RunningProfile) error
	FindByID(id uuid.UUID) (rp model.RunningProfile, err error)
}
type PlanRepository interface {
	Create(plan model.Plan) error
}

type Repositories struct {
	UserRepository
	UserActivityRepository
	RunningProfileRepository
	PlanRepository
}

func NewRepos(db *gorm.DB) *Repositories {
	return &Repositories{
		UserRepository: UserRepo{
			db: db,
		},
		UserActivityRepository: UserActivityRepo{
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
		UserActivityRepository:   UserActivityRepoMock{},
		RunningProfileRepository: RunningProfileRepoMock{},
		PlanRepository:           PlanRepoMock{},
	}
}
