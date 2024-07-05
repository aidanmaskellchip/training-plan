package application

import (
	"net/http"
	"training-plan/internal/transport/response"
)

func (app *App) HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	env := response.Envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.Config.Env,
			"version":     app.Config.Version,
		},
	}

	err := response.WriteJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.Logger.Println(err)
		response.ServerErrorResponse(w, r)
	}
}
