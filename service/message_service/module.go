package messageservice

import "github.com/MalikSaddique/chat_application_go/models"

type MessageService interface {
	SendMessage(senderID string, msg models.SendMessageRequest) error
	GetMessages(senderID, receiverID string) ([]models.Message, error)
}
