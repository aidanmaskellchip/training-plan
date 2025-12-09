package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"training-plan/internal/api/application"
	"training-plan/internal/api/transport/handlers"
	"training-plan/internal/api/transport/response"
)

func Routes(app *application.App) *httprouter.Router {
	router := httprouter.New()

	//error handling
	router.NotFound = http.HandlerFunc(response.NotFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(response.MethodNotAllowedResponse)

	//admin
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", handlers.HealthcheckHandler(app))
	router.HandlerFunc(http.MethodGet, "/v1/ping", handlers.PingDBHandler(app))

	//user
	router.HandlerFunc(http.MethodPost, "/v1/users/create", handlers.CreateUserHandler(app))
	router.HandlerFunc(http.MethodGet, "/v1/users/:id", handlers.FindUserHandler(app))
	router.HandlerFunc(http.MethodGet, "/v1/users/:id/running-profiles", handlers.FindUserRunningProfilesHandler(app))
	router.HandlerFunc(http.MethodGet, "/v1/users/:id/stats/overview", handlers.GetUserStatsHandler(app))
	router.HandlerFunc(http.MethodGet, "/v1/users/:id/profile", handlers.GetUserProfileHandler(app))
	router.HandlerFunc(http.MethodGet, "/v1/users/:id/accolades", handlers.GetUserAccoladesHandler(app))

	//running profile
	router.HandlerFunc(http.MethodGet, "/v1/running-profiles/:id", handlers.FindRunningProfileHandler(app))
	router.HandlerFunc(http.MethodPost, "/v1/running-profiles/create", handlers.CreateRunningProfileHandler(app))

	// user activity
	router.HandlerFunc(http.MethodPost, "/v1/users-activities/upload", handlers.UploadUserActivityHandler(app))
	router.HandlerFunc(http.MethodPatch, "/v1/users-activities/:id/edit", handlers.EditUserActivityHandler(app))
	// edit a user's uploaded activity

	//plan
	// TODO: IN PROGRESS...
	//router.HandlerFunc(http.MethodPost, "/v1/plans/create", handlers.CreatePlanHandler(app)

	// TODO:
	// Accolades ? e.g. 5 uploaded activities
	// Event driven

	return router
}
