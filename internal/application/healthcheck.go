package application

import (
	"net/http"
	"training-plan/internal/transport"
	"training-plan/internal/transport/response"
)

func (app *App) HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
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

func (app *App) PingDBHandler(w http.ResponseWriter, r *http.Request) {
	if err := app.DB.Exec("SELECT 1").Error; err != nil {
		response.ServerErrorResponse(w, r)
	}

	err := transport.WriteJSON(w, http.StatusOK, transport.Envelope{"msg": "db says hello"}, nil)
	if err != nil {
		app.Logger.Println(err)
		response.ServerErrorResponse(w, r)
	}
}
