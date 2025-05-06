package router

import (
	"net/http"

	"github.com/MalikSaddique/chat_application_go/models"
	"github.com/MalikSaddique/chat_application_go/utils"
	"github.com/gin-gonic/gin"
)

// SignUp godoc
// @Summary      Register a new user
// @Description  Create a new user account
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body  models.UserLoginReq  true  "User Info"
// @Success      201
// @Failure      400
// @Router       /signup [post]
func (r *Router) SignUp(c *gin.Context) {
	var req *models.UserSignUp
	help := utils.DecryptErrors
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": help(err)})
		return
	}

	signup := models.UserSignUp{
		Email:    req.Email,
		Password: req.Password,
	}
	response := r.AuthService.SignUp(c, &signup)
	if response == nil {
		return
	}
	c.JSON(http.StatusOK, response)
}
