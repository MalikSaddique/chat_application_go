package messageserviceimpl

import (
	"github.com/MalikSaddique/chat_application_go/database"
	messageservice "github.com/MalikSaddique/chat_application_go/service/message_service"
)

type MessageServiceImpl struct {
	UserAuth database.Storage
}

func NewMessageService(input database.Storage) messageservice.MessageService {
	return &MessageServiceImpl{
		UserAuth: input,
	}
}

type NewMessageServiceImpl struct {
	UserAuth database.Storage
}
