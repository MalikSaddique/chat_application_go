package database

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/MalikSaddique/chat_application_go/models"
	"github.com/gin-gonic/gin"
)

type Storage interface {
	FindUserByEmail(email string) (*models.UserLogin, error)
	SignUp(c *gin.Context, req *models.UserSignUp) *models.UserSignUp
	SaveMessage(senderID string, msg models.SendMessageRequest) error
}

type StorageImpl struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) Storage {
	return &StorageImpl{
		db: db,
	}
}

func (u *StorageImpl) FindUserByEmail(email string) (*models.UserLogin, error) {
	fmt.Println(59)

	var user models.UserLogin

	err := u.db.QueryRow("SELECT id,  email, password FROM users WHERE email=$1", email).Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil

}

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
