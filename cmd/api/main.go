package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
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
	go func() {
		err := app.EventRouter.Run(context.Background())
		if err != nil {
			panic(fmt.Sprintf("Unable to run the event router: %v", err))
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, os.Kill)
		<-c
		cancel()
	}()

	runHTTP(ctx)

	err := app.EventRouter.Close()

	app.Logger.Fatal(err)
}

func runHTTP(ctx context.Context) {
	app.Logger.Printf("starting %s server", "4001")

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", app.Config.Port),
		Handler:      app.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	go func() {
		<-ctx.Done()
		_ = srv.Close()
	}()

	if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}
