package application

import (
	"net/http"
	"training-plan/internal/application/action"
	"training-plan/internal/transport"
	"training-plan/internal/transport/request"
	"training-plan/internal/transport/response"
)

func (app *App) UploadUserActivityHandler(w http.ResponseWriter, r *http.Request) {
	req := &request.UploadUserActivityRequest{}

	err := transport.ReadJSON(w, r, &req)
	if err != nil {
		response.BadRequestResponse(w, r, err)
		return
	}

	if err = action.UploadUserActivityAction(req, app.Repos); err != nil {
		response.BadRequestResponse(w, r, err)
		return
	}

	err = transport.WriteJSON(w, http.StatusOK, transport.Envelope{"msg": "success"}, nil)
	if err != nil {
		app.Logger.Println(err)
		response.ServerErrorResponse(w, r)
	}
}

func (app *App) EditUserActivityHandler(w http.ResponseWriter, r *http.Request) {
	req := &request.EditUserActivityRequest{}

	err := transport.ReadJSON(w, r, &req)
	if err != nil {
		response.BadRequestResponse(w, r, err)
		return
	}

	actId := transport.ReadParam(r, "id")

	if err = action.EditUserActivityAction(actId, req, app.Repos); err != nil {
		response.BadRequestResponse(w, r, err)
		return
	}

	err = transport.WriteJSON(w, http.StatusOK, transport.Envelope{"msg": "success"}, nil)
	if err != nil {
		app.Logger.Println(err)
		response.ServerErrorResponse(w, r)
	}
}
