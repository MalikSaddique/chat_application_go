package mongodb

import (
	"context"

	"github.com/MalikSaddique/chat_application_go/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MessageInterface interface {
	SaveMessage(c *gin.Context, msg models.Message) error
	FetchMessages(senderID, receiverID int64, skip, limit int) ([]models.Message, error)
	UpdateMessageDB(c context.Context, id primitive.ObjectID, updatedMsg *models.Message) error
	DeleteMessageDB(c context.Context, id primitive.ObjectID) error
	FetchUndeliveredMessages(receiverID int64) ([]models.Message, error)
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
