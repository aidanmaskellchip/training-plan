package application

import (
	"net/http"
	"training-plan/internal/api/application/action"
	"training-plan/internal/api/transport"
	"training-plan/internal/api/transport/request"
	"training-plan/internal/api/transport/response"
)

func (app *App) CreatePlanHandler(w http.ResponseWriter, r *http.Request) {
	input := request.CreatePlanRequest{}

	err := transport.ReadJSON(w, r, &input)
	if err != nil {
		response.BadRequestResponse(w, r, err)
		return
	}

	plan, err := action.CreatePlanAction(&input, app.Repos)
	if err != nil {
		response.BadRequestResponse(w, r, err)
		return
	}

	err = transport.WriteJSON(w, http.StatusOK, transport.Envelope{"msg": "success", "plan": plan}, nil)
	if err != nil {
		app.Logger.Println(err)
		response.ServerErrorResponse(w, r)
	}
}
