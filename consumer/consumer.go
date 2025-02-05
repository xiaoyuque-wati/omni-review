package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	pubsub "cloud.google.com/go/pubsub"
	"github.com/gorilla/mux"
)

type EmailRequest struct {
	Email   string `json:"email"`
	Message string `json:"message"`
}

func main() {
	ctx := context.Background()
	pubsubClient, err := pubsub.NewClient(ctx, "wati-gke")
	if err != nil {
		log.Fatalf("failed to create pubsub client: %v", err)
	}
	sub := pubsubClient.Subscription("commerce-dev-omni-review-sub-1")
	sub.ReceiveSettings.MaxOutstandingMessages = 10
	sub.ReceiveSettings.MaxOutstandingBytes = 1e9

	go func() {
		err = sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
			log.Printf("Received message: %s", msg.Data)
			// Simulate email sending
			err := sendEmail(msg.Data)
			if err != nil {
				log.Printf("Failed to send email: %v", err)
				msg.Nack()
				return
			}
			msg.Ack()
		})
		if err != nil {
			log.Fatalf("failed to receive messages: %v", err)
		}
	}()

	// Set up REST API
	r := mux.NewRouter()
	r.HandleFunc("/send-email", sendEmailHandler).Methods("POST")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func sendEmailHandler(w http.ResponseWriter, r *http.Request) {
	var req EmailRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Simulate email sending
	err := sendEmail([]byte(fmt.Sprintf("Email: %s, Message: %s", req.Email, req.Message)))
	if err != nil {
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Email sent successfully"))
}

func sendEmail(data []byte) error {
	// Implement your email sending logic here
	log.Printf("Sending email with data: %s", data)
	return nil
}
