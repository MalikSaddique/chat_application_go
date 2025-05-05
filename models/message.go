package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	SenderID   int64              `json:"sender_id,omitempty" bson:"sender_id,omitempty"`
	ReceiverID int64              `json:"receiver_id,omitempty" bson:"receiver_id,omitempty"`
	Message    string             `json:"message" bson:"message"`
	Timestamp  time.Time          `json:"timestamp,omitempty" bson:"timestamp,omitempty"`
}

type SendMessageRequest struct {
	ReceiverID string `bson:"receiver_id"`
	Message    string `bson:"message"`
}

type Conversation struct {
	ChatId       int64        `bson:"chat_id,omitempty" json:"chat_id,omitempty"`
	Participants Participants `bson:"participants"`
	LastMessage  string       `bson:"last_message"`
	LastUpdated  time.Time    `bson:"last_updated"`
}

type Participants struct {
	SenderID   int64 `bson:"sender_id"`
	ReceiverID int64 `bson:"receiver_id"`
}
