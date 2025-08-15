package application

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"net/http"
	"training-plan/internal/api/application/action"
	query2 "training-plan/internal/api/application/query"
	"training-plan/internal/api/infrastructure/event/events"
	"training-plan/internal/api/transport"
	"training-plan/internal/api/transport/request"
	"training-plan/internal/api/transport/response"
)

func (app *App) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	input := request.CreateUserRequest{}

	err := transport.ReadJSON(w, r, &input)
	if err != nil {
		response.BadRequestResponse(w, r, err)
		return
	}

	user, err := action.CreateUserAction(&input, app.Repos)
	if err != nil {
		response.BadRequestResponse(w, r, err)
		return
	}

	evData := events.UserCreatedEvent{
		UserID:   user.ID.String(),
		Username: user.Username,
	}

	msg := message.NewMessage(watermill.NewUUID(), evData.ToBytes())

	err = app.SqsPub.Publish("user_created_topic", msg)
	if err != nil {
		response.BadRequestResponse(w, r, err)
	}

	err = transport.WriteJSON(w, http.StatusOK, transport.Envelope{"msg": "success"}, nil)
	if err != nil {
		app.Logger.Println(err)
		response.ServerErrorResponse(w, r)
	}
}

func (app *App) FindUserHandler(w http.ResponseWriter, r *http.Request) {
	id := transport.ReadParam(r, "id")

	user, err := query2.FindUserQuery(&id, app.Repos)
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

func (app *App) FindUserRunningProfilesHandler(w http.ResponseWriter, r *http.Request) {
	id := transport.ReadParam(r, "id")

	rps, err := query2.FindUserRunningProfilesQuery(&id, app.Repos)
	if err != nil {
		app.Logger.Println(err)
		response.BadRequestResponse(w, r, err)
		return
	}

	err = transport.WriteJSON(w, http.StatusOK, transport.Envelope{
		"msg": "success",
		"data": transport.Envelope{
			"profiles": rps,
		},
	}, nil)

	if err != nil {
		app.Logger.Println(err)
		response.ServerErrorResponse(w, r)
	}
}

func (app *App) GetUserStatsHandler(w http.ResponseWriter, r *http.Request) {
	id := transport.ReadParam(r, "id")

	_, err := query2.FindUserQuery(&id, app.Repos)
	if err != nil {
		app.Logger.Println(err)
		response.BadRequestResponse(w, r, err)
		return
	}

	stats, err := query2.GetUserStatsQuery(&id, app.Repos)
	if err != nil {
		app.Logger.Println(err)
		response.BadRequestResponse(w, r, err)
		return
	}

	err = transport.WriteJSON(w, http.StatusOK, transport.Envelope{
		"msg": "success",
		"data": transport.Envelope{
			"stats": stats,
		},
	}, nil)

	if err != nil {
		app.Logger.Println(err)
		response.ServerErrorResponse(w, r)
	}
}

func (app *App) GetUserAccoladesHandler(w http.ResponseWriter, r *http.Request) {
	id := transport.ReadParam(r, "id")

	accolades, err := app.AccoladeService.GetUserAccolades(id)
	if err != nil {
		app.Logger.Println(err)
		response.ServerErrorResponse(w, r)
		return
	}

	err = transport.WriteJSON(w, http.StatusOK, transport.Envelope{
		"msg": "success",
		"data": transport.Envelope{
			"accolades": accolades,
		},
	}, nil)

	if err != nil {
		app.Logger.Println(err)
		response.ServerErrorResponse(w, r)
	}
}

func (app *App) GetUserProfileHandler(w http.ResponseWriter, r *http.Request) {
	id := transport.ReadParam(r, "id")

	_, err := query2.FindUserQuery(&id, app.Repos)
	if err != nil {
		app.Logger.Println(err)
		response.BadRequestResponse(w, r, err)
		return
	}

	profile, err := query2.GetUserProfileQuery(&id, app.Repos)
	if err != nil {
		app.Logger.Println(err)
		response.BadRequestResponse(w, r, err)
		return
	}

	err = transport.WriteJSON(w, http.StatusOK, transport.Envelope{
		"msg": "success",
		"data": transport.Envelope{
			"profile": profile,
		},
	}, nil)

	if err != nil {
		app.Logger.Println(err)
		response.ServerErrorResponse(w, r)
	}
}
