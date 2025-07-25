package model

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestPlanBeforeCreate(t *testing.T) {
	t.Parallel()

	plan := Plan{}
	err := plan.BeforeCreate(&gorm.DB{})

	assert.NoError(t, err)
	assert.NotNil(t, plan.ID)
}
