package messageserviceimpl

import (
	"strconv"

	"github.com/MalikSaddique/chat_application_go/models"
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
