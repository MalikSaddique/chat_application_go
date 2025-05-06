package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (u *MessageInterfaceImpl) DeleteMessageDB(c context.Context, id primitive.ObjectID) error {
	collection := u.mongoClient.Database("chat_app_go").Collection("messages")

	_, err := collection.DeleteOne(c, bson.M{"_id": id})
	return err
}
