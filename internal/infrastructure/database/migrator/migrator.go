package migrator

import (
	"fmt"
	"gorm.io/gorm"
	model2 "training-plan/internal/domain/model"
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

	if err := create(&model2.User{}); err != nil {
		return fmt.Errorf("failed to migrate users: %w", err)
	}

	if err := create(&model2.RunningProfile{}); err != nil {
		return fmt.Errorf("failed to migrate running profiles: %w", err)
	}

	return
}
