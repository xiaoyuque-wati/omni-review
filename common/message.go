package common

import "time"

// Message represents the main message structure
type Message struct {
	ID          string    `json:"id"`
	TenantID    string    `json:"tenantId"`
	Recipient   Recipient `json:"recipient"`
	ScheduledAt int64     `json:"scheduledAt"`
}

// Recipient represents the recipient information
type Recipient struct {
	ID     string            `json:"id"`
	Type   string            `json:"type"`
	Fields map[string]string `json:"fields"`
	Value  string            `json:"value"`
}

// NewMessage creates a new Message with default values
func NewMessage() *Message {
	return &Message{
		TenantID: "123456",
		Recipient: Recipient{
			Type: "email",
			Fields: map[string]string{
				"name": "xiaoyu",
			},
			Value: "test@clare.ai",
		},
		ScheduledAt: time.Now().Unix(),
	}
}
