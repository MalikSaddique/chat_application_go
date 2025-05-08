package websocketsimpl

import (
	messageservice "github.com/MalikSaddique/chat_application_go/controllers/message_service"
	mongodb "github.com/MalikSaddique/chat_application_go/db/mongoDB"
	db "github.com/MalikSaddique/chat_application_go/db/postgresDB"
	websockets "github.com/MalikSaddique/chat_application_go/web_sockets"
)

type WebSocketsImpl struct {
	MessageAuth    mongodb.MessageInterface
	MessageService messageservice.MessageService
}

func NewWebSockets(input mongodb.MessageInterface, messageWeb messageservice.MessageService) websockets.WebSockets {
	return &WebSocketsImpl{
		MessageAuth:    input,
		MessageService: messageWeb,
	}
}

type NewMessageServiceImpl struct {
	UserAuth db.Storage
}

var _ websockets.WebSockets = &WebSocketsImpl{}
