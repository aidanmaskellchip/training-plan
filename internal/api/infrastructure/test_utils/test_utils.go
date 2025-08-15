package testutils

import (
	"log"
	"os"
	"testing"
	"training-plan/internal/api/application"
	"training-plan/internal/api/infrastructure/config"
	"training-plan/internal/api/infrastructure/repository"
)

func NewTestApplication(t *testing.T) (app *application.App) {
	t.Helper()

	return &application.App{
		Config: NewMockConfig(t),
		Logger: log.New(os.Stdout, "", log.Ldate|log.Ltime),
		Repos:  repository.NewMockRepos(),
	}
}

func NewMockConfig(t *testing.T) (conf *config.Config) {
	t.Helper()

	return &config.Config{
		Env:    "testing",
		AppEnv: "local",
	}
}
