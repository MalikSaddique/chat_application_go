package router

import (
	"net/http"

	"github.com/MalikSaddique/chat_application_go/models"
	"github.com/MalikSaddique/chat_application_go/utils"
	"github.com/gin-gonic/gin"
)

// Login godoc
// @Summary      Login a user
// @Description  Authenticate user and return JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body  models.UserLoginReq  true  "User Credentials"
// @Success      200
// @Failure      401
// @Router       /login [post]
func (r *Router) Login(c *gin.Context) {

	var req models.UserLoginReq
	var login models.UserLogin
	help := utils.DecryptErrors

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": help(err)})
		return
	}

	login = models.UserLogin{
		Email:    req.Email,
		Password: req.Password,
	}

	tokenPair, err := r.AuthService.Login(c, &login)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tokenPair)
}
