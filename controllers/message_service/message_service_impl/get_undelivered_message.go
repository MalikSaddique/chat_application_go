package messageserviceimpl

import "github.com/MalikSaddique/chat_application_go/models"

func (m *MessageServiceImpl) GetUndeliveredMessages(receiverID int64) ([]models.Message, error) {
	return m.MessageAuth.FetchUndeliveredMessages(receiverID)
}
