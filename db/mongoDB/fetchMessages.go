package mongodb

import (
	"context"
	"fmt"
	"strconv"

	"github.com/MalikSaddique/chat_application_go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (u *MessageInterfaceImpl) FetchMessages(senderIDStr, receiverIDStr string, skip, limit int) ([]models.Message, error) {
	filter := bson.M{}

	if senderIDStr != "" {
		senderID, err := strconv.Atoi(senderIDStr)
		if err == nil {
			filter["sender_id"] = senderID
		}
	}

	if receiverIDStr != "" {
		receiverID, err := strconv.Atoi(receiverIDStr)
		if err == nil {
			filter["receiver_id"] = receiverID
		}
	}

	messageCollection := u.mongoClient.Database("chat_app_go").Collection("messages")
	findOptions := options.Find()
	findOptions.SetLimit(int64(limit))
	findOptions.SetSkip(int64(skip))

	cursor, err := messageCollection.Find(context.TODO(), filter, findOptions)
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
	fmt.Println(messages)

	return messages, nil
}
