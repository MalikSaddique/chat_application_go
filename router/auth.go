package router

import (
	"net/http"

	"github.com/MalikSaddique/chat_application_go/models"
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

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
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
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	signup := models.UserSignUp{
		// Id:       req.Id,
		Email:    req.Email,
		Password: req.Password,
	}
	response := r.AuthService.SignUp(c, &signup)
	if response == nil {
		return
	}
	c.JSON(http.StatusOK, response)

}

// RefreshKey godoc
// @Summary      Refresh Access Token
// @Description  Validates refresh token and generates a new access token
// @Tags         auth
// @Security     BearerAuth
// @Produce      json
// @Success      200  "Returns new access token"
// @Failure      401  "Unauthorized"
// @Failure      500 "Internal Server Error"
// @Router       /refresh [get]
func (r *Router) RefreshKey(c *gin.Context) {
	newToken, err := r.AuthService.RefreshAccessToken(c)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"access_token": newToken,
	})
}
