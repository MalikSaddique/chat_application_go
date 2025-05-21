package authserviceimpl

import (
	"log"

	"github.com/MalikSaddique/chat_application_go/models"
	"github.com/gin-gonic/gin"
)

func (m *AuthServiceImpl) SearchUser(c *gin.Context, email string) ([]models.UserResponse, error) {
	users, err := m.userAuth.SearchUser(email)
	if err != nil {
		log.Println(err)
	}
	return users, nil
}
