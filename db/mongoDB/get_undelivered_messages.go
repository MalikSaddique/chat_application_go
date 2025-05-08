package mongodb

import (
	"context"

	"github.com/MalikSaddique/chat_application_go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (u *MessageInterfaceImpl) FetchUndeliveredMessages(receiverID int64) ([]models.Message, error) {
	filter := bson.M{
		"receiver_id": receiverID,
		"delivered":   false,
	}

	messageCollection := u.mongoClient.Database("chat_app_go").Collection("messages")

	cursor, err := messageCollection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var messages []models.Message
	var ID []primitive.ObjectID
	for cursor.Next(context.TODO()) {
		var msg models.Message
		if err := cursor.Decode(&msg); err != nil {
			return nil, err
		}
		messages = append(messages, msg)

		if msg.ID != primitive.NilObjectID {
			ID = append(ID, msg.ID)
		}
	}

	if len(ID) > 0 {
		updateFilter := bson.M{"_id": bson.M{"$in": ID}}
		update := bson.M{"$set": bson.M{"delivered": true}}

		_, updateErr := messageCollection.UpdateMany(context.TODO(), updateFilter, update)
		if updateErr != nil {
			return nil, updateErr
		}
	}

	return messages, nil
}
