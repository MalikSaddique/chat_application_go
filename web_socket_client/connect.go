package websocketclient

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var Conn *websocket.Conn

func ConnectToWebSocketServer(url string, token string) {
	var err error
	header := http.Header{}
	header.Add("Authorization", "Bearer "+token)
	Conn, _, err = websocket.DefaultDialer.Dial(url, header)
	if err != nil {
		log.Fatalf("Failed to connect to WebSocket server: %v", err)
	}
	log.Println("Connected to WebSocket server")

}
