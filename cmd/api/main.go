package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
	"training-plan/internal/api/application"
)

var app application.App

func init() {
	var err error
	app, err = application.Bootstrap()

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

	//go func() {
	//	c := make(chan os.Signal, 1)
	//	signal.Notify(c, os.Interrupt, os.Kill)
	//	<-c
	//	cancel()
	//}()

	runHTTP()

	err := app.EventRouter.Close()

	app.Logger.Fatal(err)
}

func runHTTP() {
	app.Logger.Printf("starting %s server", "4001")

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", app.Config.Port),
		Handler:      Routes(&app),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}
