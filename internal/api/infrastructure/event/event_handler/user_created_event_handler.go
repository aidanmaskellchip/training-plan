package eventhandler

import (
	"fmt"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
)

type UserCreatedEventHandler struct{}

func (uch UserCreatedEventHandler) Handle(msg *message.Message) error {
	fmt.Printf("received message: %s, payload: %s\n", msg.UUID, string(msg.Payload))

	// now do some things
	fmt.Println("now doing some other stuff...")

	fmt.Println("handler actions complete")
	return nil
}

func (uch UserCreatedEventHandler) HandleAndPublish(msg *message.Message) ([]*message.Message, error) {
	fmt.Printf("received message: %s, payload: %s\n", msg.UUID, string(msg.Payload))

	// now do some things
	fmt.Println("now doing some other stuff...")

	// prepare messages to be published to another topic
	msg = message.NewMessage(watermill.NewUUID(), []byte("success message from user created event handler"))
	return message.Messages{msg}, nil
}
