package messageserviceimpl

import (
	"github.com/MalikSaddique/chat_application_go/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (u *MessageServiceImpl) UpdateMessage(c *gin.Context, idStr string, updatedMsg *models.Message) (*models.Message, error) {
	objID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return nil, err
	}

	err = u.MessageAuth.UpdateMessageDB(c, objID, updatedMsg)
	if err != nil {
		return nil, err
	}

	return updatedMsg, nil
}
