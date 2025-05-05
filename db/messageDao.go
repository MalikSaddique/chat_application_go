package db

import (
	"context"
	"strconv"
	"time"

	"github.com/MalikSaddique/chat_application_go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MessageInterface interface {
	SaveMessage(senderID string, receiverID string, msg models.Message) error
	FetchMessages(senderID string, receiverID string, offset int, limit int) ([]models.Message, error)
}

type MessageInterfaceImpl struct {
	mongoClient *mongo.Client
}

func NewMongoDb(client *mongo.Client) MessageInterface {
	return &MessageInterfaceImpl{
		mongoClient: client,
	}

}

var _ MessageInterface = &MessageInterfaceImpl{}

func (u *MessageInterfaceImpl) SaveMessage(senderID string, receiverID string, msg models.Message) error {
	sid, err := strconv.ParseInt(senderID, 10, 64)
	if err != nil {
		return err
	}

	rid, err := strconv.ParseInt(receiverID, 10, 64)
	if err != nil {
		return err
	}

	c := context.TODO()
	db := u.mongoClient.Database("chat_app_go")
	convoCollection := db.Collection("conversations")
	messageCollection := db.Collection("messages")

	var conversationID int64

	if msg.ReceiverID == 0 {

		participants := models.Participants{
			SenderID:   sid,
			ReceiverID: rid,
		}

		var conversation models.Conversation
		err = convoCollection.FindOne(c, bson.M{"participants": participants}).Decode(&conversation)

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
		SenderID:   sid,
		ReceiverID: rid,
		Message:    msg.Message,
		Timestamp:  time.Now(),
	}

	_, err = messageCollection.InsertOne(c, message)
	return err
}

func (u *MessageInterfaceImpl) FetchMessages(senderIDStr string, receiverIDStr string, skip int, limit int) ([]models.Message, error) {
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
	findOptions.SetSort(bson.D{{"timestamp", -1}})

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
