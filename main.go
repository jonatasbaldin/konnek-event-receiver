package main

import (
	"context"
	"log"

	cloudevents "github.com/cloudevents/sdk-go"
	"github.com/kelseyhightower/envconfig"
)

type EnvConfig struct {
	Port        int    `envconfig:"PORT" default:"8080"`
	Sink        string `envconfig:"K_SINK"`
	CEOverrides string `envconfig:"K_CE_OVERRIDES"`
}

var sink string

func receiveAndForward(event cloudevents.Event) {
	transport, err := cloudevents.NewHTTPTransport(cloudevents.WithTarget(sink))
	if err != nil {
		log.Fatalf("could not create transport: %v", err)
	}

	client, err := cloudevents.NewClient(transport)
	if err != nil {
		log.Fatalf("could not create client: %v", err)
	}

	ctx := context.Background()

	_, _, err = client.Send(ctx, event)
	if err != nil {
		log.Fatalf("could not send event: %v", err)
	}
}

func main() {
	var env EnvConfig
	err := envconfig.Process("", &env)
	if err != nil {
		log.Fatalf("could not load environment variables: %v", err)
	}

	sink = env.Sink

	client, err := cloudevents.NewDefaultClient()
	if err != nil {
		log.Fatalf("could not create client: %v", err)
	}

	ctx := context.Background()

	err = client.StartReceiver(ctx, receiveAndForward)
	if err != nil {
		log.Fatalf("failed to start receiver, %v", err)
	}

	log.Printf("listening on port %d", env.Port)
}
