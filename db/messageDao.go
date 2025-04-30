package db

import (
	"context"
	"strconv"
	"time"

	"github.com/MalikSaddique/chat_application_go/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type MessageInterface interface {
	SaveMessage(senderID string, msg models.Message) error
	FetchMessages(senderID, receiverID string) ([]models.Message, error)
}

type MessageInterfaceImpl struct {
	mongoClient *mongo.Client
}

func NewMongoDb(client *mongo.Client) MessageInterface {
	return &MessageInterfaceImpl{
		mongoClient: client,
	}

}

func (u *MessageInterfaceImpl) SaveMessage(senderID string, msg models.Message) error {
	sid, err := strconv.ParseInt(senderID, 10, 64)
	if err != nil {
		return err
	}

	message := models.Message{
		SenderID:   sid,
		ReceiverID: msg.ReceiverID,
		Message:    msg.Message,
		Timestamp:  time.Now(),
	}

	collection := u.mongoClient.Database("chat_app_go").Collection("messages")
	_, err = collection.InsertOne(context.TODO(), message)

	return err
}

func (u *MessageInterfaceImpl) FetchMessages(senderID, receiverID string) ([]models.Message, error) {
	query := `
		SELECT sender_id, receiver_id, message, timestamp
		FROM messages
		WHERE (sender_id = $1 AND receiver_id = $2) OR (sender_id = $2 AND receiver_id = $1)
		ORDER BY timestamp ASC
	`
	rows, err := u.db.Query(query, senderID, receiverID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var msg models.Message
		err := rows.Scan(&msg.SenderID, &msg.ReceiverID, &msg.Message, &msg.Timestamp)
		if err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}
	return messages, nil
}
