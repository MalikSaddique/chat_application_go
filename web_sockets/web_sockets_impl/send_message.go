package websocketsimpl

import (
	"log"

	"github.com/MalikSaddique/chat_application_go/models"
	"github.com/gin-gonic/gin"
)

func (w *WebSocketsImpl) SendMessage(c *gin.Context, msg models.Message) {
	err := w.MessageAuth.SaveMessage(c, msg)
	if err != nil {
		log.Println("Failed to save message to DB:", err)
	}
}
