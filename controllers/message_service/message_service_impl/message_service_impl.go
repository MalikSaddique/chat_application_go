package messageserviceimpl

import (
	messageservice "github.com/MalikSaddique/chat_application_go/controllers/message_service"
	mongodb "github.com/MalikSaddique/chat_application_go/db/mongoDB"
	db "github.com/MalikSaddique/chat_application_go/db/postgresDB"
	websockets "github.com/MalikSaddique/chat_application_go/web_sockets"
)

type MessageServiceImpl struct {
	MessageAuth mongodb.MessageInterface
	WebSocket   websockets.WebSockets
}

func NewMessageService(input mongodb.MessageInterface, web websockets.WebSockets) messageservice.MessageService {
	return &MessageServiceImpl{
		MessageAuth: input,
		WebSocket:   web,
	}
}

type NewMessageServiceImpl struct {
	UserAuth db.Storage
	websockets.WebSockets
}

var _ messageservice.MessageService = &MessageServiceImpl{}
