package messageserviceimpl

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (u *MessageServiceImpl) DeleteMessage(c *gin.Context, idStr string) error {
	objID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return err
	}

	err = u.MessageAuth.DeleteMessageDB(c, objID)
	return err
}
