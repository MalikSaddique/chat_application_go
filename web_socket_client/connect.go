package websocketclient

import (
	"log"

	"github.com/gorilla/websocket"
)

var Conn *websocket.Conn

func ConnectToWebSocketServer(baseURL string, key string) {
	wsURL := baseURL + "?key=" + key

	var err error
	Conn, _, err = websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		log.Fatalf("Failed to connect to WebSocket server: %v", err)
	}
	log.Println("Connected to WebSocket server")
}
