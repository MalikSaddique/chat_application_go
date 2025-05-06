package messageserviceimpl

import "github.com/MalikSaddique/chat_application_go/models"

func (m *MessageServiceImpl) SendMessage(senderID, receiverID int64, msg models.Message) error {
	err := m.MessageAuth.SaveMessage(senderID, receiverID, msg)
	if err != nil {
		return err
	}
	return nil
}
