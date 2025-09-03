package factory

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"training-plan/internal/api/domain/activity/entities"
	"training-plan/internal/api/domain/model"
)

func TestNewWeek(t *testing.T) {
	t.Parallel()

	week := NewWeek()

	assert.Equal(t, model.ActivityWeek{
		Mon: entities.Activity{},
		Tue: entities.Activity{},
		Wed: entities.Activity{},
		Thu: entities.Activity{},
		Fri: entities.Activity{},
		Sat: entities.Activity{},
		Sun: entities.Activity{},
	}, week)
}
