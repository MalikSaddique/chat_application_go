package messageserviceimpl

import (
	"github.com/MalikSaddique/chat_application_go/models"
)

func (m *MessageServiceImpl) SendMessage(senderID string, receiverID string, msg models.Message) error {
	err := m.MessageAuth.SaveMessage(senderID, receiverID, msg)
	if err != nil {
		return err
	}
	return nil
}

func (m *MessageServiceImpl) GetMessages(chatID string) ([]models.Message, error) {
	return m.MessageAuth.FetchMessages(chatID)
}
