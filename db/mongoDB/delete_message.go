package mongodb

import (
	"context"
	"fmt"
	"strconv"

	"github.com/MalikSaddique/chat_application_go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (u *MessageInterfaceImpl) DeleteMessageDB(c context.Context, id primitive.ObjectID) error {
	userID := c.Value("userID")
	collection := u.mongoClient.Database("chat_app_go").Collection("messages")
	var msg models.Message
	err := collection.FindOne(c, bson.M{"_id": id}).Decode(&msg)
	if err != nil {
		return fmt.Errorf("message not found or DB error: %v", err)
	}
	sID := strconv.FormatInt(msg.SenderID, 10)
	rID := strconv.FormatInt(msg.ReceiverID, 10)

	if sID != userID && rID != userID {
		return fmt.Errorf("unauthorized: you are not Authorize to this message")
	}

	_, err = collection.DeleteOne(c, bson.M{"_id": id})
	if err != nil {
		return fmt.Errorf("message not deleted: %v", err)
	}
	return err
}
