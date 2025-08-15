package repository

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	model2 "training-plan/internal/api/domain/model"
	vo "training-plan/internal/api/domain/value_objects"
)

var ErrNoRecord = errors.New("no matching record found")

type UserRepository interface {
	Create(user model2.User) (*model2.User, error)
	FindByID(id uuid.UUID) (user model2.User, err error)
}
type UserActivityRepository interface {
	Create(ua model2.UserActivity) error
	FindByID(id uuid.UUID) (ua model2.UserActivity, err error)
	GetFastestUserActivity(userID uuid.UUID) (stats model2.ActivityStats, err error)
	GetLongestUserActivity(userID uuid.UUID) (stats model2.ActivityStats, err error)
	GetFastestCommunityActivity() (stats model2.ActivityStats, err error)
	GetLongestCommunityActivity() (stats model2.ActivityStats, err error)
	GetMostCommonActivityType(userID uuid.UUID) (t vo.ActivityType, err error)
	Update(act model2.UserActivity) error
}
type RunningProfileRepository interface {
	Create(profile model2.RunningProfile) error
	FindByUserID(id uuid.UUID) (rps []model2.RunningProfile, err error)
	FindLatestUserProfile(id uuid.UUID) (rps model2.RunningProfile, err error)
	FindByID(id uuid.UUID) (rp model2.RunningProfile, err error)
}
type PlanRepository interface {
	Create(plan model2.Plan) error
	FindByID(id uuid.UUID) (p model2.Plan, err error)
	FindLatestUserPlan(id uuid.UUID) (p model2.Plan, err error)
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
