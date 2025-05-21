package authservice

import (
	"github.com/MalikSaddique/chat_application_go/models"
	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Login(c *gin.Context, login *models.UserLogin) (*models.TokenPair, error)
	SignUp(c *gin.Context, req *models.UserSignUp) *models.UserSignUp
	RefreshAccessToken(c *gin.Context) (string, error)
	SearchUser(c *gin.Context, email string) ([]models.UserResponse, error)
}
