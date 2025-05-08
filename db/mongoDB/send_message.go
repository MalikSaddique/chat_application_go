package mongodb

import (
	"time"

	"github.com/MalikSaddique/chat_application_go/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (u *MessageInterfaceImpl) SaveMessage(c *gin.Context, msg models.Message) error {
	db := u.mongoClient.Database("chat_app_go")
	convoCollection := db.Collection("conversations")
	messageCollection := db.Collection("messages")

	var conversationID int64

	if msg.ReceiverID == 0 {

		participants := models.Participants{
			SenderID:   msg.SenderID,
			ReceiverID: msg.ReceiverID,
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
	// msg.Delivered = false
	message := models.Message{
		SenderID:   msg.SenderID,
		ReceiverID: msg.ReceiverID,
		Message:    msg.Message,
		Timestamp:  time.Now(),
	}

	_, err := messageCollection.InsertOne(c, message)
	return err
}
