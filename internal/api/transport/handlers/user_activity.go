package handlers

import (
	"net/http"
	"training-plan/internal/api/application"
	"training-plan/internal/api/application/action"
	"training-plan/internal/api/transport"
	"training-plan/internal/api/transport/request"
	"training-plan/internal/api/transport/response"
)

func UploadUserActivityHandler(app *application.App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
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
}

func EditUserActivityHandler(app *application.App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
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
}
