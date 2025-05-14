package modelsWeb

import "time"

type MessageRes struct {
	ReceiverID int64     `json:"receiver_id,omitempty" bson:"receiver_id,omitempty"`
	SenderID   int64     `json:"sender_id,omitempty" bson:"sender_id,omitempty"`
	Message    string    `json:"message" bson:"message"`
	Timestamp  time.Time `json:"timestamp,omitempty" bson:"timestamp,omitempty"`
}
