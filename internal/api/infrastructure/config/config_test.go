package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) { //nolint: tparallel // cannot run suite using t.Setenv in parallel
	config, err := Get()

	t.Run("should not return an error", func(t *testing.T) {
		t.Parallel()

		assert.Nil(t, err)
	})

	t.Run("should correctly return unset default values", func(t *testing.T) {
		t.Parallel()

		assert.Equal(t, "local", config.Env)
	})

	t.Run("should return values set from the ENV", func(t *testing.T) {
		cfg = nil

		t.Setenv("ENV", "mock-env")

		c, _ := Get()
		assert.Equal(t, "mock-env", c.Env)
	})

	t.Run("should return existing config if set", func(t *testing.T) {
		cfg = nil

		c, _ := Get()
		assert.Equal(t, "local", c.Env)

		// Updating an env variable now should not impact our existing config
		t.Setenv("ENV", "mock-env")

		// This config should not have the updated value as it should return the previous instance
		newConfig, _ := Get()

		assert.Equal(t, "local", newConfig.Env)
	})
}
