package factory

import (
	"training-plan/internal/api/domain/activity/entities"
	"training-plan/internal/api/domain/model"
)

func NewWeek() model.ActivityWeek {
	return model.ActivityWeek{
		Mon: entities.Activity{},
		Tue: entities.Activity{},
		Wed: entities.Activity{},
		Thu: entities.Activity{},
		Fri: entities.Activity{},
		Sat: entities.Activity{},
		Sun: entities.Activity{},
	}
}
