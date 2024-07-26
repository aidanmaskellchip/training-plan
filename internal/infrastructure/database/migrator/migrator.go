package migrator

import (
	"fmt"
	"gorm.io/gorm"
	"training-plan/internal/domain/model"
)

var database *gorm.DB

func create(i interface{}) error {
	if err := database.AutoMigrate(&i); err != nil {
		return fmt.Errorf("error: %w", err)
	}

	return nil
}

func Migrate(con *gorm.DB) (err error) {
	database = con

	if err := create(&model.User{}); err != nil {
		return fmt.Errorf("failed to migrate users: %w", err)
	}

	if err := create(&model.RunningProfile{}); err != nil {
		return fmt.Errorf("failed to migrate running profiles: %w", err)
	}

	if err := create(&model.UserActivity{}); err != nil {
		return fmt.Errorf("failed to migrate user activities: %w", err)
	}

	return
}
