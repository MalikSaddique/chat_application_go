package messageserviceimpl

import (
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (u *MessageServiceImpl) DeleteMessage(c *gin.Context, idStr string) error {
	objID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return err
	}
	userID := c.Query("userID")
	if userID == "" {
		log.Println("User not authenticated")
	}

	err = u.MessageAuth.DeleteMessageDB(c, objID)
	return err
}
