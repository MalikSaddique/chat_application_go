package main

import (
	"log"

	messageserviceimpl "github.com/MalikSaddique/chat_application_go/controllers/message_service/message_service_impl"
	mongodb "github.com/MalikSaddique/chat_application_go/db/mongoDB"
	"github.com/MalikSaddique/socket/router"
	"github.com/MalikSaddique/socket/websockets/websocketsimpl"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	connMongo, err := mongodb.MongoDbConn()
	if err != nil {
		log.Fatalf("MongoDB connection error: %s", err)
	}
	messagedb := mongodb.NewMongoDb(connMongo)

	websockets := websocketsimpl.NewWebSockets(messagedb)
	messageService := messageserviceimpl.NewMessageService(messagedb, websockets)

	webSocketRouter := router.NewRouter(messageService, websockets)
	err = webSocketRouter.Engine.Run(":8004")
	if err != nil {
		log.Fatalf("Websocket server failed to start: %s", err)
	}

}
