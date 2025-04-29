package messageservice

import "github.com/gin-gonic/gin"

type MessageService interface {
	SendMessage(c *gin.Context) error
}
