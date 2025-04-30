package messageserviceimpl

import (
	messageservice "github.com/MalikSaddique/chat_application_go/controllers/message_service"
	"github.com/MalikSaddique/chat_application_go/db"
)

type MessageServiceImpl struct {
	MessageAuth db.MessageInterface
}

func NewMessageService(input db.MessageInterface) messageservice.MessageService {
	return &MessageServiceImpl{
		MessageAuth: input,
	}
}

type NewMessageServiceImpl struct {
	UserAuth db.Storage
}
