package main

import (
	"context"
	"log"

	cloudevents "github.com/cloudevents/sdk-go"
	"github.com/kelseyhightower/envconfig"
	"knative.dev/eventing/pkg/kncloudevents"
)

type EnvConfig struct {
	Sink        string `envconfig:"K_SINK"`
	CEOverrides string `envconfig:"K_CE_OVERRIDES"`
}

var sink string

func receiveAndForward(event cloudevents.Event) {
	client, err := kncloudevents.NewDefaultClient(sink)
	if err != nil {
		log.Fatalf("could not create forwarder client: %v\n", err)
	}

	ctx := context.Background()

	_, _, err = client.Send(ctx, event)
	if err != nil {
		log.Fatalf("could not send event: %v\n", err)
	}
}

func main() {
	var envConfig EnvConfig
	err := envconfig.Process("", &envConfig)
	if err != nil {
		log.Fatalf("could not load environment variables: %v\n", err)
	}

	sink = envConfig.Sink

	client, err := kncloudevents.NewDefaultClient()
	if err != nil {
		log.Fatalf("could not create client: %v\n", err)
	}

	ctx := context.Background()

	err = client.StartReceiver(ctx, receiveAndForward)
	if err != nil {
		log.Fatalf("failed to start receiver, %v", err)
	}

	log.Printf("listening on port %d\n", 8080)
}
