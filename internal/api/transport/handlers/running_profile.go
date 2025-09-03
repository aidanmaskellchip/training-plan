package handlers

import (
	"net/http"
	"training-plan/internal/api/application"
	"training-plan/internal/api/application/action"
	"training-plan/internal/api/application/query"
	"training-plan/internal/api/transport"
	"training-plan/internal/api/transport/request"
	"training-plan/internal/api/transport/response"
)

func CreateRunningProfileHandler(app *application.App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		input := request.CreateRunningProfileRequest{}

		err := transport.ReadJSON(w, r, &input)
		if err != nil {
			response.BadRequestResponse(w, r, err)
			return
		}

		if err = action.CreateRunningProfileAction(&input, app.Repos); err != nil {
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

func FindRunningProfileHandler(app *application.App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := transport.ReadParam(r, "id")

		rp, err := query.FindRunningProfileQuery(&id, app.Repos)
		if err != nil {
			app.Logger.Println(err)
			response.BadRequestResponse(w, r, err)
			return
		}

		err = transport.WriteJSON(w, http.StatusOK, transport.Envelope{
			"msg": "success",
			"data": transport.Envelope{
				"running_profile": rp,
			},
		}, nil)

		if err != nil {
			app.Logger.Println(err)
			response.ServerErrorResponse(w, r)
		}
	}
}
