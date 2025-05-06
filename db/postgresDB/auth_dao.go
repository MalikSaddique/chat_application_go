package db

import (
	"database/sql"

	"github.com/MalikSaddique/chat_application_go/models"
	"github.com/MalikSaddique/chat_application_go/pkg/logger"
	"github.com/gin-gonic/gin"
)

var log = logger.Logger("Postgres-DB")

type Storage interface {
	FindUserByEmail(email string) (*models.UserLogin, error)
	SignUp(c *gin.Context, req *models.UserSignUp) *models.UserSignUp
}

type StorageImpl struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) Storage {
	return &StorageImpl{
		db: db,
	}

}
