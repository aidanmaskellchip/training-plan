package application

import (
	"net/http"
	action2 "training-plan/internal/api/application/action"
	"training-plan/internal/api/transport"
	request2 "training-plan/internal/api/transport/request"
	"training-plan/internal/api/transport/response"
)

func (app *App) UploadUserActivityHandler(w http.ResponseWriter, r *http.Request) {
	req := &request2.UploadUserActivityRequest{}

	err := transport.ReadJSON(w, r, &req)
	if err != nil {
		response.BadRequestResponse(w, r, err)
		return
	}

	if err = action2.UploadUserActivityAction(req, app.Repos); err != nil {
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
	req := &request2.EditUserActivityRequest{}

	err := transport.ReadJSON(w, r, &req)
	if err != nil {
		response.BadRequestResponse(w, r, err)
		return
	}

	actId := transport.ReadParam(r, "id")

	if err = action2.EditUserActivityAction(actId, req, app.Repos); err != nil {
		response.BadRequestResponse(w, r, err)
		return
	}

	err = transport.WriteJSON(w, http.StatusOK, transport.Envelope{"msg": "success"}, nil)
	if err != nil {
		app.Logger.Println(err)
		response.ServerErrorResponse(w, r)
	}
}
