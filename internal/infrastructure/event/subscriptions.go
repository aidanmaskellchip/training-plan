package event

import eventtypes "training-plan/internal/infrastructure/event/event_types"

// LoadSubscriptions Here we can list which channels should be listening to what events
func (eb *EventBus) LoadSubscriptions() {
	ch := make(chan Event)
	eb.Subscribe(eventtypes.ActivityUploadedEventTypeName, ch)
}
