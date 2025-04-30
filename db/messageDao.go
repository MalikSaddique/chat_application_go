package db

import (
	"context"
	"strconv"
	"time"

	"github.com/MalikSaddique/chat_application_go/models"
	"go.mongodb.org/mongo-driver/bson"
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
	sid, err := strconv.ParseInt(senderID, 10, 64)
	if err != nil {
		return nil, err
	}

	rid, err := strconv.ParseInt(receiverID, 10, 64)
	if err != nil {
		return nil, err
	}
	filter := bson.M{
		"$or": []bson.M{
			{"sender_id": sid, "receiver_id": rid},
			{"sender_id": rid, "receiver_id": sid},
		},
	}

	collection := u.mongoClient.Database("chat_app_go").Collection("messages")

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var messages []models.Message
	for cursor.Next(context.TODO()) {
		var msg models.Message
		if err := cursor.Decode(&msg); err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}

	return messages, nil
}
