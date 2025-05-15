package websocketsimpl

import (
	mongodb "github.com/MalikSaddique/chat_application_go/db/mongoDB"
	db "github.com/MalikSaddique/chat_application_go/db/postgresDB"
	websockets "github.com/MalikSaddique/socket/websockets"
	"github.com/gorilla/websocket"
)

type WebSocketsImpl struct {
	MessageAuth mongodb.MessageInterface
	Clients     map[int]*websocket.Conn
}

func NewWebSockets(input mongodb.MessageInterface) websockets.WebSockets {
	return &WebSocketsImpl{
		MessageAuth: input,
		Clients:     make(map[int]*websocket.Conn),
	}
}

type NewMessageServiceImpl struct {
	UserAuth db.Storage
}

var _ websockets.WebSockets = &WebSocketsImpl{}
