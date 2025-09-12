package repository

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"training-plan/internal/api/domain/plan"
	vo "training-plan/internal/api/domain/plan/entities"
	runningprofile "training-plan/internal/api/domain/running_profile"
	"training-plan/internal/api/domain/user"
	"training-plan/internal/api/domain/user_activity"
)

var ErrNoRecord = errors.New("no matching record found")

type UserRepository interface {
	Create(user *user.Entity) (*user.Entity, error)
	FindByID(id uuid.UUID) (*user.Entity, error)
}
type UserActivityRepository interface {
	Create(ua *useractivity.Entity) error
	FindByID(id uuid.UUID) (ua *useractivity.Entity, err error)
	GetFastestUserActivity(userID uuid.UUID) (stats useractivity.ActivityStats, err error)
	GetLongestUserActivity(userID uuid.UUID) (stats useractivity.ActivityStats, err error)
	GetFastestCommunityActivity() (stats useractivity.ActivityStats, err error)
	GetLongestCommunityActivity() (stats useractivity.ActivityStats, err error)
	GetMostCommonActivityType(userID uuid.UUID) (t vo.ActivityType, err error)
	Update(act *useractivity.Entity) error
}
type RunningProfileRepository interface {
	Create(profile *runningprofile.Entity) error
	FindByUserID(id uuid.UUID) (rps []runningprofile.Entity, err error)
	FindLatestUserProfile(id uuid.UUID) (rps *runningprofile.Entity, err error)
	FindByID(id uuid.UUID) (rp *runningprofile.Entity, err error)
}
type PlanRepository interface {
	Create(plan *plan.Entity) error
	FindByID(id uuid.UUID) (p *plan.Entity, err error)
	FindLatestUserPlan(id uuid.UUID) (p *plan.Entity, err error)
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
