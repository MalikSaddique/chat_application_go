package router

import (
	"github.com/gin-gonic/gin"
)

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
