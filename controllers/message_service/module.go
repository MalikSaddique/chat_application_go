package messageservice

import "github.com/MalikSaddique/chat_application_go/models"

type MessageService interface {
	SendMessage(senderID string, receiverID string, msg models.Message) error
	GetMessages(senderID string, receiverID string, pageStr string, limitStr string) ([]models.Message, error)
}
