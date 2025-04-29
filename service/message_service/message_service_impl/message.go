package messageserviceimpl

import (
	"github.com/MalikSaddique/chat_application_go/models"
)

func (m *MessageServiceImpl) SendMessage(senderID string, msg models.SendMessageRequest) error {
	err := m.UserAuth.SaveMessage(senderID, msg)
	if err != nil {
		return err
	}
	return nil
}

func (m *MessageServiceImpl) GetMessages(senderID, receiverID string) ([]models.Message, error) {
	return m.UserAuth.FetchMessages(senderID, receiverID)
}
