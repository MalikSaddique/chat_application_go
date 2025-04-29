package models

import (
	"time"
)

type Message struct {
	ID        int       `json:"id" db:"id"`
	Sender    string    `json:"sender" db:"sender"`
	Receiver  string    `json:"receiver" db:"receiver"`
	Message   string    `json:"message" db:"content"`
	Timestamp time.Time `json:"timestamp" db:"timestamp"`
}

type SendMessageRequest struct {
	ReceiverID string `json:"receiver_id" binding:"required"`
	Message    string `json:"message" binding:"required"`
}
