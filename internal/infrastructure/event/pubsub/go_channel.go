package pubsub

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	eventhandler "training-plan/internal/infrastructure/event/event_handler"
)

func NewGoChannelEventRouter(logger *watermill.LoggerAdapter, sub *gochannel.GoChannel) (*message.Router, error) {
	router, err := message.NewRouter(message.RouterConfig{}, *logger)
	if err != nil {
		return nil, err
	}

	// This is for gracefully shutting down router
	router.AddPlugin(plugin.SignalsHandler)

	// Add middleware for all routes -- needs arguments
	router.AddMiddleware()

	// Add a handler for the topic -- also a method for using handler which then publishes to another topic
	handler := router.AddNoPublisherHandler(
		"user_created_event_handler",
		"user_created_topic",
		sub,
		eventhandler.UserCreatedEventHandler{}.Handle,
	)

	// Add handler specific middleware -- needs arguments
	handler.AddMiddleware()

	return router, nil
}

func NewGoChannelPubSub(logger *watermill.LoggerAdapter) *gochannel.GoChannel {
	return gochannel.NewGoChannel(
		gochannel.Config{},
		*logger,
	)
}
