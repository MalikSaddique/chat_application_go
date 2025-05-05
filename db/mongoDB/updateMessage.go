package mongodb

import (
	"context"

	"github.com/MalikSaddique/chat_application_go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (u *MessageInterfaceImpl) UpdateMessageDB(c context.Context, id primitive.ObjectID, updatedMsg *models.Message) error {
	collection := u.mongoClient.Database("chat_app_go").Collection("messages")

	update := bson.M{
		"$set": bson.M{
			"message": updatedMsg.Message,
		},
	}

	_, err := collection.UpdateByID(c, id, update)
	return err
}
