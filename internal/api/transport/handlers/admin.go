package handlers

import (
	"net/http"
	"training-plan/internal/api/application"
	"training-plan/internal/api/transport"
	"training-plan/internal/api/transport/response"
)

func HealthcheckHandler(app *application.App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		env := transport.Envelope{
			"status": "available",
			"system_info": map[string]string{
				"environment": app.Config.Env,
				"version":     app.Config.Version,
			},
		}

		err := transport.WriteJSON(w, http.StatusOK, env, nil)
		if err != nil {
			app.Logger.Println(err)
			response.ServerErrorResponse(w, r)
		}
	}
}

func PingDBHandler(app *application.App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := app.DB.Exec("SELECT 1").Error; err != nil {
			response.ServerErrorResponse(w, r)
			return
		}

		err := transport.WriteJSON(w, http.StatusOK, transport.Envelope{"msg": "db says hello"}, nil)
		if err != nil {
			app.Logger.Println(err)
			response.ServerErrorResponse(w, r)
		}
	}
}
