package application

import (
	"net/http"
	"training-plan/internal/application/action"
	"training-plan/internal/application/query"
	"training-plan/internal/transport"
	"training-plan/internal/transport/request"
	"training-plan/internal/transport/response"
)

func (app *App) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	input := request.CreateUserRequest{}

	err := transport.ReadJSON(w, r, &input)
	if err != nil {
		response.BadRequestResponse(w, r, err)
		return
	}

	if err = action.CreateUserAction(&input, app.Repos); err != nil {
		response.BadRequestResponse(w, r, err)
		return
	}

	err = transport.WriteJSON(w, http.StatusOK, transport.Envelope{"msg": "success"}, nil)
	if err != nil {
		app.Logger.Println(err)
		response.ServerErrorResponse(w, r)
	}
}

func (app *App) FindUserHandler(w http.ResponseWriter, r *http.Request) {
	id := transport.ReadParam(r, "id")

	user, err := query.FindUserQuery(&id, app.Repos)
	if err != nil {
		app.Logger.Println(err)
		response.BadRequestResponse(w, r, err)
		return
	}

	err = transport.WriteJSON(w, http.StatusOK, transport.Envelope{
		"msg": "success",
		"data": transport.Envelope{
			"user": user,
		},
	}, nil)

	if err != nil {
		app.Logger.Println(err)
		response.ServerErrorResponse(w, r)
	}
}

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
