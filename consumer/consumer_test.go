package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendEmailHandler(t *testing.T) {
	// Create a sample request payload
	emailReq := EmailRequest{
		Email:   "test@example.com",
		Message: "Hello, this is a test message.",
	}
	payload, err := json.Marshal(emailReq)
	if err != nil {
		t.Fatalf("Failed to marshal request payload: %v", err)
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", "/send-email", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(sendEmailHandler)

	// Serve the HTTP request
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expected := "Email sent successfully"
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
