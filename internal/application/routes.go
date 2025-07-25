package application

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"training-plan/internal/transport/response"
)

func (app *App) Routes() *httprouter.Router {
	router := httprouter.New()

	//error handling
	router.NotFound = http.HandlerFunc(response.NotFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(response.MethodNotAllowedResponse)

	//admin
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.HealthcheckHandler)
	router.HandlerFunc(http.MethodGet, "/v1/ping", app.PingDBHandler)

	//user
	router.HandlerFunc(http.MethodPost, "/v1/users/create", app.CreateUserHandler)
	router.HandlerFunc(http.MethodGet, "/v1/users/:id", app.FindUserHandler)
	router.HandlerFunc(http.MethodGet, "/v1/users/:id/running-profiles", app.FindUserRunningProfilesHandler)
	router.HandlerFunc(http.MethodGet, "/v1/users/:id/stats/overview", app.GetUserStatsHandler)
	router.HandlerFunc(http.MethodGet, "/v1/users/:id/profile", app.GetUserProfileHandler)
	router.HandlerFunc(http.MethodGet, "/v1/users/:id/accolades", app.GetUserAccoladesHandler)

	//running profile
	router.HandlerFunc(http.MethodGet, "/v1/running-profiles/:id", app.FindRunningProfileHandler)
	router.HandlerFunc(http.MethodPost, "/v1/running-profiles/create", app.CreateRunningProfileHandler)

	// user activity
	router.HandlerFunc(http.MethodPost, "/v1/users-activities/upload", app.UploadUserActivityHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/users-activities/:id/edit", app.EditUserActivityHandler)
	// edit a user's uploaded activity

	//plan
	// TODO: IN PROGRESS...
	//router.HandlerFunc(http.MethodPost, "/v1/plans/create", app.CreatePlanHandler)

	// TODO:
	// Accolades ? e.g. 5 uploaded activities
	// Event driven

	return router
}
