package models

import (
	"time"
)

type Message struct {
	// ID         primitive.ObjectID `bson:"_id,omitempty"`
	SenderID   int64     `bson:"sender_id"`
	ReceiverID int64     `bson:"receiver_id"`
	Message    string    `bson:"message"`
	Timestamp  time.Time `bson:"timestamp"`
}

type SendMessageRequest struct {
	ReceiverID string `bson:"receiver_id"`
	Message    string `bson:"message"`
}
