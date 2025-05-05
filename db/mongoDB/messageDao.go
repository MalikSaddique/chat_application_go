package mongodb

import (
	"context"

	"github.com/MalikSaddique/chat_application_go/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MessageInterface interface {
	SaveMessage(senderID string, receiverID string, msg models.Message) error
	FetchMessages(senderID string, receiverID string, offset int, limit int) ([]models.Message, error)
	UpdateMessageDB(c context.Context, id primitive.ObjectID, updatedMsg *models.Message) error
	DeleteMessageDB(c context.Context, id primitive.ObjectID) error
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
