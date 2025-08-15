package weekfactory

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"training-plan/internal/api/domain/model"
	"training-plan/internal/api/domain/value_objects"
)

func TestNewWeek(t *testing.T) {
	t.Parallel()

	week := NewWeek()

	assert.Equal(t, model.ActivityWeek{
		Mon: valueobjects.Activity{},
		Tue: valueobjects.Activity{},
		Wed: valueobjects.Activity{},
		Thu: valueobjects.Activity{},
		Fri: valueobjects.Activity{},
		Sat: valueobjects.Activity{},
		Sun: valueobjects.Activity{},
	}, week)
}
