package websocketsimpl

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	modelsWeb "github.com/MalikSaddique/socket/models"
	"github.com/gorilla/websocket"
)

func (w *WebSocketsImpl) SendMessage(msg *modelsWeb.MessageRes) {
	msgResponse := modelsWeb.MessageRes{
		SenderID:   msg.SenderID,
		ReceiverID: msg.ReceiverID,
		Message:    msg.Message,
		Timestamp:  msg.Timestamp,
	}
	receiverID := strconv.FormatInt(msg.ReceiverID, 10)
	// log.Println("Id", receiverID)
	// conn, ok := ConnMap[receiverID]
	// fmt.Println("222", ConnMap[receiverID], "awdkh")
	// log.Println("conn rec", conn)
	// if ok {
	// 	msgBytes, err := json.Marshal(msgResponse)
	// 	if err == nil {
	// 		err = conn.WriteMessage(websocket.TextMessage, msgBytes)
	// 		if err != nil {
	// 			log.Println("error", "Failed to send message via WebSocket", err)
	// 		}
	// 	}
	// 	log.Println("Message sent to user:", receiverID)
	// 	log.Println("Message", msgResponse)
	// } else if !ok {
	// 	log.Println("No connection found for receiverID", receiverID)
	// }
	ConnLock.Lock()
	defer ConnLock.Unlock()

	fmt.Println("ConnMap")
	fmt.Println(ConnMap)

	conn, ok := ConnMap[receiverID]
	if !ok {
		log.Println("No connection found for user:", receiverID)
	}
	if ok {
		msgBytes, err := json.Marshal(msgResponse)
		if err == nil {
			err = conn.WriteMessage(websocket.TextMessage, msgBytes)
			if err != nil {
				log.Println("error", "Failed to send message via WebSocket", err)
			}
		}

	}
}
