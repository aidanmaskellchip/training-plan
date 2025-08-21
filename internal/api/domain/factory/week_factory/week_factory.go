package weekfactory

import (
	"training-plan/internal/api/domain/model"
	"training-plan/internal/api/domain/value_objects"
)

func NewWeek() model.ActivityWeek {
	return model.ActivityWeek{
		Mon: valueobjects.Activity{},
		Tue: valueobjects.Activity{},
		Wed: valueobjects.Activity{},
		Thu: valueobjects.Activity{},
		Fri: valueobjects.Activity{},
		Sat: valueobjects.Activity{},
		Sun: valueobjects.Activity{},
	}
}
