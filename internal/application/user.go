package application

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"net/http"
	"training-plan/internal/application/action"
	"training-plan/internal/application/query"
	"training-plan/internal/infrastructure/event/events"
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

func (app *App) FindUserRunningProfilesHandler(w http.ResponseWriter, r *http.Request) {
	id := transport.ReadParam(r, "id")

	rps, err := query.FindUserRunningProfilesQuery(&id, app.Repos)
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

	_, err := query.FindUserQuery(&id, app.Repos)
	if err != nil {
		app.Logger.Println(err)
		response.BadRequestResponse(w, r, err)
		return
	}

	stats, err := query.GetUserStatsQuery(&id, app.Repos)
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

func (app *App) GetUserProfileHandler(w http.ResponseWriter, r *http.Request) {
	id := transport.ReadParam(r, "id")

	_, err := query.FindUserQuery(&id, app.Repos)
	if err != nil {
		app.Logger.Println(err)
		response.BadRequestResponse(w, r, err)
		return
	}

	profile, err := query.GetUserProfileQuery(&id, app.Repos)
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
