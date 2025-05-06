package websockets

import "github.com/gorilla/websocket"

type WebSockets interface {
	AddConn(userID string, wsConn *websocket.Conn)
}
