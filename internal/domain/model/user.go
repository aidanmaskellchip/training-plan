package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID              uuid.UUID        `json:"id" gorm:"type:uuid;primaryKey;"`
	Username        string           `json:"username" gorm:"type:varchar(255);not null"`
	RunningProfiles []RunningProfile `json:"runningProfiles" gorm:"foreignKey:UserID"`
	Plans           []Plan           `json:"plans" gorm:"foreignKey:UserID"`
	gorm.Model
}
