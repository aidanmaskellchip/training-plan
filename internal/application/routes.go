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

	//running profile
	router.HandlerFunc(http.MethodPost, "/v1/running-profiles/create", app.CreateRunningProfileHandler)

	//plan
	router.HandlerFunc(http.MethodPost, "/v1/plans/create", app.CreatePlanHandler)

	return router
}
