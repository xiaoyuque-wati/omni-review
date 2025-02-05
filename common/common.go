package common

import (
	"context"
	"log"

	pubsub "cloud.google.com/go/pubsub"
)

func CreatePubSubClient(projectID string) (*pubsub.Client, error) {
	client, err := pubsub.NewClient(context.Background(), projectID)
	if err != nil {
		log.Fatalf("failed to create pubsub client: %v", err)
		return nil, err
	}
	return client, nil
}
