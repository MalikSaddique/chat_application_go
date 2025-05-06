package authserviceimpl

import (
	"fmt"

	"github.com/MalikSaddique/chat_application_go/models"
	"github.com/gin-gonic/gin"
)

func (u *AuthServiceImpl) SignUp(c *gin.Context, req *models.UserSignUp) *models.UserSignUp {

	createdUser := u.userAuth.SignUp(c, req)

	if createdUser == nil {
		fmt.Println("signup error")
		return nil
	}

	response := models.UserSignUp{
		// Id:       createdUser.Id,
		Email:    createdUser.Email,
		Password: createdUser.Password,
		Message:  "User created successfully",
	}

	return &response
}
