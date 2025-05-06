package mongodb

import (
	"context"
	"time"

	"github.com/MalikSaddique/chat_application_go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (u *MessageInterfaceImpl) SaveMessage(senderID, receiverID int64, msg models.Message) error {
	c := context.TODO()
	db := u.mongoClient.Database("chat_app_go")
	convoCollection := db.Collection("conversations")
	messageCollection := db.Collection("messages")

	var conversationID int64

	if msg.ReceiverID == 0 {

		participants := models.Participants{
			SenderID:   senderID,
			ReceiverID: receiverID,
		}

		var conversation models.Conversation
		err := convoCollection.FindOne(c, bson.M{"participants": participants}).Decode(&conversation)

		if err == mongo.ErrNoDocuments {
			newConvo := models.Conversation{
				Participants: participants,
				LastMessage:  msg.Message,
				LastUpdated:  time.Now(),
			}
			res, err := convoCollection.InsertOne(c, newConvo)
			if err != nil {
				return err
			}
			if oid, ok := res.InsertedID.(primitive.ObjectID); ok {

				conversationID = oid.Timestamp().Unix()
			}
		} else if err != nil {
			return err
		} else {
			conversationID = conversation.ChatId
			_, err := convoCollection.UpdateOne(c, bson.M{"chat_id": conversationID}, bson.M{
				"$set": bson.M{"last_message": msg.Message, "last_updated": time.Now()},
			})
			if err != nil {
				return err
			}
		}
	} else {
		conversationID = msg.ReceiverID
		_, err := convoCollection.UpdateOne(c, bson.M{"chat_id": conversationID}, bson.M{
			"$set": bson.M{"last_message": msg.Message, "last_updated": time.Now()},
		})
		if err != nil {
			return err
		}
	}

	message := models.Message{
		SenderID:   senderID,
		ReceiverID: receiverID,
		Message:    msg.Message,
		Timestamp:  time.Now(),
	}

	_, err := messageCollection.InsertOne(c, message)
	return err
}
