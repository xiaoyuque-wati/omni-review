package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	pubsub "cloud.google.com/go/pubsub"
	"github.com/xiaoyuque-wati/omni-review/common"
)

func main() {
	pubsubClient, err := pubsub.NewClient(context.Background(), "wati-gke")
	if err != nil {
		log.Fatalf("failed to create pubsub client: %v", err)
	}
	topic := pubsubClient.Topic("commerce-dev-omni-review")

	ticker := time.NewTicker(10 * time.Second) // Adjust the interval as needed
	defer ticker.Stop()

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Service is healthy")
	})

	go http.ListenAndServe(":8080", nil)
	fmt.Println("Server is running on port 8080")
	for {
		select {
		case <-ticker.C:
			message := common.NewMessage()

			// Customize fields if needed
			message.Recipient.Fields["department"] = "engineering"

			jsonData, err := json.Marshal(message)
			if err != nil {
				log.Printf("failed to marshal message: %v", err)
				continue
			}

			result := topic.Publish(context.Background(), &pubsub.Message{
				Data: jsonData,
			})
			_, err = result.Get(context.Background())
			if err != nil {
				log.Printf("failed to publish message: %v", err)
			} else {
				log.Printf("message published successfully")
			}
		}
	}
}
