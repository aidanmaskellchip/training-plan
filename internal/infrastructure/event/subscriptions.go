package event

import eventtypes "training-plan/internal/infrastructure/event/event_types"

func (eb *EventBus) LoadSubscriptions() {
	eb.Subscribe(eventtypes.ActivityUploadedEventTypeName, ActivityUploadedCh)
}
