package messageserviceimpl

import (
	"strconv"

	"github.com/MalikSaddique/chat_application_go/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (m *MessageServiceImpl) SendMessage(senderID string, receiverID string, msg models.Message) error {
	err := m.MessageAuth.SaveMessage(senderID, receiverID, msg)
	if err != nil {
		return err
	}
	return nil
}

func (m *MessageServiceImpl) GetMessages(senderID string, receiverID string, pageStr string, limitStr string) ([]models.Message, error) {
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}
	skip := (page - 1) * limit

	return m.MessageAuth.FetchMessages(senderID, receiverID, skip, limit)
}

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

func (u *MessageServiceImpl) DeleteMessage(c *gin.Context, idStr string) error {
	objID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return err
	}

	err = u.MessageAuth.DeleteMessageDB(c, objID)
	return err
}
