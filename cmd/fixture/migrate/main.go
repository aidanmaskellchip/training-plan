package main

import (
	"fmt"
	"training-plan/internal/api/application"
	"training-plan/internal/api/infrastructure/database/migrator"
)

var app application.App

func init() {
	var err error
	app, err = application.Bootstrap()

	if err != nil {
		panic(fmt.Sprintf("Unable to init the application or fixture: %v", err))
	}
}

func main() {
	app.Logger.Println("running migrations")

	err := migrator.Migrate(app.DB)
	if err != nil {
		app.Logger.Panicf("failed database migration: %s", err)
	}

	app.Logger.Println("successfully migrated database")
}
