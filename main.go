package main

import (
	"log"

	"github.com/MalikSaddique/chat_application_go/database"
	"github.com/MalikSaddique/chat_application_go/routes"
	authserviceimpl "github.com/MalikSaddique/chat_application_go/service/auth_service/auth_service_impl"
	messageserviceimpl "github.com/MalikSaddique/chat_application_go/service/message_service/message_service_impl"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	conn, err := database.DbConnection()
	if err != nil {
		log.Fatalf("Database connection error: %s", err)
	}

	userdb := database.NewStorage(conn)

	authService := authserviceimpl.NewAuthService(authserviceimpl.NewAuthServiceImpl{
		UserAuth: userdb,
	})
	messageService := messageserviceimpl.NewMessageService(userdb)

	router := routes.NewRouter(authService, messageService)

	router.Engine.Run(":8002")

}
