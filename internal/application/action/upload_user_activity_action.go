package action

import (
	"errors"
	"fmt"
	"time"
	"training-plan/internal/domain/factory/activity_type_factory"
	"training-plan/internal/domain/model"
	"training-plan/internal/infrastructure/event"
	eventhandlers "training-plan/internal/infrastructure/event/event_handlers"
	eventtypes "training-plan/internal/infrastructure/event/event_types"
	"training-plan/internal/infrastructure/repository"
	"training-plan/internal/transport/request"
)

func UploadUserActivityAction(data *request.UploadUserActivityRequest, repos *repository.Repositories, eb *event.EventBus, ec *event.Channels) (err error) {
	if err := data.Validate(); err != nil {
		return err
	}

	_, err = repos.UserRepository.FindByID(data.UserID)
	if err != nil {
		return errors.New("user not found")
	}

	at, err := activitytypefactory.NewActivityType(data.Type)
	if err != nil {
		return err
	}

	ua := &model.UserActivity{
		UserID:   data.UserID,
		Type:     at.Type,
		Distance: data.Distance,
		Pace:     data.Pace,
	}

	if err := repos.UserActivityRepository.Create(*ua); err != nil {
		return err
	}

	go eventhandlers.ActivityUploadedHandler(ec.ActivityUploadedCh)

	ev := event.Event{
		Type:      eventtypes.ActivityUploadedEventTypeName,
		Timestamp: time.Now(),
		Data:      eventtypes.ActivityUploadedEvent{UserID: data.UserID.String()},
	}
	fmt.Println("hello")

	eb.Publish(ev)

	return nil
}
