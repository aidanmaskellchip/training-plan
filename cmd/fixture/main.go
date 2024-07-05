package main

import (
	"fmt"
	"training-plan/internal/application"
	"training-plan/internal/database/migrator"
)

var app application.App

func init() {
	var err error
	app, err = application.Setup()

	if err != nil {
		panic(fmt.Sprintf("Unable to init the application or fixture: %v", err))
	}
}

func main() {
	app.Logger.Println("running migrations")

	err := migrator.Migrate(app.DB)
	if err != nil {
		app.Logger.Panic(fmt.Sprintf("failed database migration: %s", err))
	}

	app.Logger.Println("successfully migrated database")
}
