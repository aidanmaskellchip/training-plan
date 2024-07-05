package main

import (
	"fmt"
	"net/http"
	"time"
	"training-plan/internal/application"
)

var app application.App

func init() {
	var err error
	app, err = application.Setup()

	if err != nil {
		panic(fmt.Sprintf("Unable to init the application on startup: %v", err))
	}
}

func main() {
	app.Logger.Println("Starting")

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", app.Config.Port),
		Handler:      app.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.Logger.Printf("starting %s server on %s", app.Config.Env, srv.Addr)
	err := srv.ListenAndServe()

	app.Logger.Fatal(err)
}
