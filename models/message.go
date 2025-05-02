package models

import (
	"time"
)

type Message struct {
	SenderID   int64     `json:"sender_id,omitempty" bson:"sender_id,omitempty"`
	ReceiverID int64     `json:"receiver_id,omitempty" bson:"receiver_id,omitempty"`
	Message    string    `json:"message" bson:"message"`
	Timestamp  time.Time `json:"timestamp,omitempty" bson:"timestamp,omitempty"`
}

type SendMessageRequest struct {
	ReceiverID string `bson:"receiver_id"`
	Message    string `bson:"message"`
}
