package mongodb

import (
	"context"

	"github.com/MalikSaddique/chat_application_go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (u *MessageInterfaceImpl) FetchMessages(senderID, receiverID int64, skip, limit int) ([]models.Message, error) {
	filter := bson.M{}
	filter["sender_id"] = senderID
	filter["receiver_id"] = receiverID

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

	return messages, nil
}
