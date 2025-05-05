package messageserviceimpl

import (
	messageservice "github.com/MalikSaddique/chat_application_go/controllers/message_service"
	mongodb "github.com/MalikSaddique/chat_application_go/db/mongoDB"
	db "github.com/MalikSaddique/chat_application_go/db/postgresDB"
)

type MessageServiceImpl struct {
	MessageAuth mongodb.MessageInterface
}

func NewMessageService(input mongodb.MessageInterface) messageservice.MessageService {
	return &MessageServiceImpl{
		MessageAuth: input,
	}
}

type NewMessageServiceImpl struct {
	UserAuth db.Storage
}

var _ messageservice.MessageService = &MessageServiceImpl{}
