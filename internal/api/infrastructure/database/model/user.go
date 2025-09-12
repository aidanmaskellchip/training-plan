package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"training-plan/internal/api/domain/user"
)

type User struct {
	ID              uuid.UUID        `json:"id" gorm:"type:uuid;primaryKey;"`
	Username        string           `json:"username" gorm:"type:varchar(255);not null"`
	RunningProfiles []RunningProfile `json:"runningProfiles" gorm:"foreignKey:UserID"`
	UserActivities  []UserActivity   `json:"userActivities" gorm:"foreignKey:UserID"`
	Plans           []Plan           `json:"plans" gorm:"foreignKey:UserID"`
	gorm.Model
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()

	return
}

func (u *User) ToDomainEntity() *user.Entity {
	return &user.Entity{
		ID:        u.ID,
		Username:  u.Username,
		CreatedAt: &u.CreatedAt,
	}
}
