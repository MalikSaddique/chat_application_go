package router

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) SearchUser(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
		return
	}

	users, err := r.AuthService.SearchUser(c, email)
	if err != nil {
		log.Println("Error: ", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
