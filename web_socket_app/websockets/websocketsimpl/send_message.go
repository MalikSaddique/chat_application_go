package websocketsimpl

// func (w *WebSocketsImpl) SendMessage(msg *models.Message) {
// 	msgResponse := models.MessageResponse{
// 		SenderID:   msg.SenderID,
// 		ReceiverID: msg.ReceiverID,
// 		Message:    msg.Message,
// 		Timestamp:  msg.Timestamp,
// 	}
// 	receiverID := strconv.FormatInt(msg.ReceiverID, 10)
// 	conn, ok := ConnMap[receiverID]
// 	log.Println("Conn rec", conn)
// 	if ok {
// 		msgBytes, err := json.Marshal(msgResponse)
// 		if err == nil {
// 			err = conn.WriteMessage(websocket.TextMessage, msgBytes)
// 			if err != nil {
// 				log.Println("error", "Failed to send message via WebSocket", err)
// 			}
// 		}
// 	}
// 	log.Println("Message sent to user:", receiverID)
// }
