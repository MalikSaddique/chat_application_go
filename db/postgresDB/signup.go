package db

import (
	"database/sql"
	"net/http"

	"github.com/MalikSaddique/chat_application_go/models"
	"github.com/gin-gonic/gin"
)

func (u *StorageImpl) SignUp(c *gin.Context, req *models.UserSignUp) *models.UserSignUp {

	err := u.db.QueryRow("SELECT email FROM users WHERE email = $1", &req.Email).Scan(&req.Email)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists with this email"})
		return nil
	} else if err != sql.ErrNoRows {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking existing user"})
		return nil
	}

	_, err = u.db.Exec("INSERT INTO users (email, password) VALUES ($1, $2)", &req.Email, &req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return nil
	}

	return req

}
