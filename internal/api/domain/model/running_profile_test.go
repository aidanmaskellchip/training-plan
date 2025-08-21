package model

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestRunningProfileBeforeCreate(t *testing.T) {
	t.Parallel()

	rp := RunningProfile{}
	err := rp.BeforeCreate(&gorm.DB{})

	assert.NoError(t, err)
	assert.NotNil(t, rp.ID)
}
