package handlers

import (
	"net/http"
	"training-plan/internal/api/application"
	"training-plan/internal/api/transport"
	"training-plan/internal/api/transport/request"
	"training-plan/internal/api/transport/response"
)

func CreatePlanHandler(app *application.App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		input := request.CreatePlanRequest{}

		err := transport.ReadJSON(w, r, &input)
		if err != nil {
			response.BadRequestResponse(w, r, err)
			return
		}

		// create plan action to implement

		err = transport.WriteJSON(w, http.StatusOK, transport.Envelope{"msg": "success"}, nil)
		if err != nil {
			app.Logger.Println(err)
			response.ServerErrorResponse(w, r)
		}
	}
}
