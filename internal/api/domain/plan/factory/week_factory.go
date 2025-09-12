package factory

import (
	"training-plan/internal/api/domain/plan/entities"
)

func NewWeek() entities.ActivityWeek {
	return entities.ActivityWeek{
		Mon: entities.Activity{},
		Tue: entities.Activity{},
		Wed: entities.Activity{},
		Thu: entities.Activity{},
		Fri: entities.Activity{},
		Sat: entities.Activity{},
		Sun: entities.Activity{},
	}
}
