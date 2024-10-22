package eventhandlers

import (
	"fmt"
	"training-plan/internal/infrastructure/event"
	eventtypes "training-plan/internal/infrastructure/event/event_types"
)

func ActivityUploadedHandler(eChan <-chan event.Event) {
	for e := range eChan {
		actUploadedE, ok := e.Data.(eventtypes.ActivityUploadedEvent)
		if !ok {
			fmt.Println("Invalid event data")
			continue
		}

		fmt.Println("Received activity uploaded event for user: " + actUploadedE.UserID)
	}
}
