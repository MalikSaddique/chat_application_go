package websocketsimpl

import (
	mongodb "github.com/MalikSaddique/chat_application_go/db/mongoDB"
	db "github.com/MalikSaddique/chat_application_go/db/postgresDB"
	websockets "github.com/MalikSaddique/socket/websockets"
)

type WebSocketsImpl struct {
	MessageAuth mongodb.MessageInterface
}

func NewWebSockets(input mongodb.MessageInterface) websockets.WebSockets {
	return &WebSocketsImpl{
		MessageAuth: input,
	}
}

type NewMessageServiceImpl struct {
	UserAuth db.Storage
}

var _ websockets.WebSockets = &WebSocketsImpl{}
