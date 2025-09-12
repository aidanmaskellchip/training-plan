package factory

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"training-plan/internal/api/domain/plan/entities"
)

func TestNewWeek(t *testing.T) {
	t.Parallel()

	week := NewWeek()

	assert.Equal(t, entities.ActivityWeek{
		Mon: entities.Activity{},
		Tue: entities.Activity{},
		Wed: entities.Activity{},
		Thu: entities.Activity{},
		Fri: entities.Activity{},
		Sat: entities.Activity{},
		Sun: entities.Activity{},
	}, week)
}
