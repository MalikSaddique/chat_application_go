package messageserviceimpl

import (
	messageservice "github.com/MalikSaddique/chat_application_go/controllers/message_service"
	mongodb "github.com/MalikSaddique/chat_application_go/db/mongoDB"
	db "github.com/MalikSaddique/chat_application_go/db/postgresDB"
	"github.com/MalikSaddique/chat_application_go/pkg"
	websockets "github.com/MalikSaddique/socket/websockets"
)

type MessageServiceImpl struct {
	MessageAuth mongodb.MessageInterface
	WebSocket   pkg.MessagePusher
}

func NewMessageService(input mongodb.MessageInterface, web pkg.MessagePusher) messageservice.MessageService {
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
