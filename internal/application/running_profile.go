package application

import (
	"net/http"
	"training-plan/internal/application/action"
	"training-plan/internal/transport"
	"training-plan/internal/transport/request"
	"training-plan/internal/transport/response"
)

func (app *App) CreateRunningProfileHandler(w http.ResponseWriter, r *http.Request) {
	input := request.CreateRunningProfileRequest{}

	err := transport.ReadJSON(w, r, &input)
	if err != nil {
		response.BadRequestResponse(w, r, err)
		return
	}

	if err = action.CreateRunningProfileAction(&input, app.Repos); err != nil {
		response.BadRequestResponse(w, r, err)
	}

	err = transport.WriteJSON(w, http.StatusOK, transport.Envelope{"msg": "success"}, nil)
	if err != nil {
		app.Logger.Println(err)
		response.ServerErrorResponse(w, r)
	}
}