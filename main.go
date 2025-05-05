package main

import (
	authserviceimpl "github.com/MalikSaddique/chat_application_go/controllers/auth_service/auth_service_impl"
	messageserviceimpl "github.com/MalikSaddique/chat_application_go/controllers/message_service/message_service_impl"
	mongodb "github.com/MalikSaddique/chat_application_go/db/mongoDB"
	db "github.com/MalikSaddique/chat_application_go/db/postgresDB"
	"github.com/MalikSaddique/chat_application_go/pkg/logger"
	"github.com/MalikSaddique/chat_application_go/router"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	log = logger.Logger("ChatApp")
)

func main() {

	log.Info("App started")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	connMongo, err := mongodb.MongoDbConn()
	if err != nil {
		log.Fatalf("db connection error: %s", err)
	}
	conn, err := db.DbConnection()
	if err != nil {
		log.Fatalf("db connection error: %s", err)
	}

	userdb := db.NewStorage(conn)
	messagedb := mongodb.NewMongoDb(connMongo)

	authService := authserviceimpl.NewAuthService(authserviceimpl.NewAuthServiceImpl{
		UserAuth: userdb,
	})
	messageService := messageserviceimpl.NewMessageService(messagedb)

	router := router.NewRouter(authService, messageService)

	router.Engine.Run(":8002")

}
