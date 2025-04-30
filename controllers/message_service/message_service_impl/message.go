package messageserviceimpl

import (
	"github.com/MalikSaddique/chat_application_go/models"
)

func (m *MessageServiceImpl) SendMessage(senderID string, msg models.Message) error {
	err := m.MessageAuth.SaveMessage(senderID, msg)
	if err != nil {
		return err
	}
	return nil
}

func (m *MessageServiceImpl) GetMessages(senderID, receiverID string) ([]models.Message, error) {
	return m.MessageAuth.FetchMessages(senderID, receiverID)
}
