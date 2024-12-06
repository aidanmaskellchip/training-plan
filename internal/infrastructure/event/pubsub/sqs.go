package pubsub

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-aws/sqs"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
	"github.com/aws/aws-sdk-go-v2/aws"
	amazonsqs "github.com/aws/aws-sdk-go-v2/service/sqs"
	transport "github.com/aws/smithy-go/endpoints"
	"github.com/samber/lo"
	"net/url"
	eventhandler "training-plan/internal/infrastructure/event/event_handler"
)

func NewSQSSubscriber(logger *watermill.LoggerAdapter) (*sqs.Subscriber, error) {
	// Could store this in config
	sqsOpts := []func(*amazonsqs.Options){
		amazonsqs.WithEndpointResolverV2(sqs.OverrideEndpointResolver{
			Endpoint: transport.Endpoint{
				URI: *lo.Must(url.Parse("http://localstack:4566")),
			},
		}),
	}

	subscriberConfig := sqs.SubscriberConfig{
		AWSConfig: aws.Config{
			Credentials: aws.AnonymousCredentials{},
		},
		OptFns: sqsOpts,
	}

	return sqs.NewSubscriber(subscriberConfig, *logger)
}

func NewSQSPublisher(logger *watermill.LoggerAdapter) (*sqs.Publisher, error) {
	// Could store this in config
	sqsOpts := []func(*amazonsqs.Options){
		amazonsqs.WithEndpointResolverV2(sqs.OverrideEndpointResolver{
			Endpoint: transport.Endpoint{
				URI: *lo.Must(url.Parse("http://localstack:4566")),
			},
		}),
	}

	publisherConfig := sqs.PublisherConfig{
		AWSConfig: aws.Config{
			Credentials: aws.AnonymousCredentials{},
		},
		OptFns: sqsOpts,
	}

	return sqs.NewPublisher(publisherConfig, *logger)
}

func NewSQSEventRouter(logger *watermill.LoggerAdapter, sub *sqs.Subscriber) (*message.Router, error) {
	router, err := message.NewRouter(message.RouterConfig{}, *logger)
	if err != nil {
		return nil, err
	}

	// This is for gracefully shutting down router, can also do router.Close()
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
