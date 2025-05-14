package main

import (
	authserviceimpl "github.com/MalikSaddique/chat_application_go/controllers/auth_service/auth_service_impl"
	messageserviceimpl "github.com/MalikSaddique/chat_application_go/controllers/message_service/message_service_impl"
	mongodb "github.com/MalikSaddique/chat_application_go/db/mongoDB"
	db "github.com/MalikSaddique/chat_application_go/db/postgresDB"
	"github.com/MalikSaddique/chat_application_go/pkg"
	"github.com/MalikSaddique/chat_application_go/pkg/logger"
	"github.com/MalikSaddique/chat_application_go/router"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var log = logger.Logger("ChatApp")

func main() {
	log.Infof("App started")

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	connMongo, err := mongodb.MongoDbConn()
	if err != nil {
		log.Fatalf("MongoDB connection error: %s", err)
	}

	conn, err := db.DbConnection()
	if err != nil {
		log.Fatalf("PostgreSQL connection error: %s", err)
	}

	userdb := db.NewStorage(conn)
	messagedb := mongodb.NewMongoDb(connMongo)

	authService := authserviceimpl.NewAuthService(authserviceimpl.NewAuthServiceImpl{
		UserAuth: userdb,
	})

	wsClient := &pkg.WebSocketHTTPClient{
		ServerURL: "http://localhost:8004",
	}

	// messageService := messageserviceimpl.NewMessageService(&messageserviceimpl.MessageServiceImpl{
	// 	MessageAuth: messagedb,
	// 	WebSocket:   wsClient,
	// })
	messageService := messageserviceimpl.NewMessageService(messagedb, wsClient)

	httpRouter := router.NewRouter(authService, messageService)

	if err := httpRouter.Engine.Run(":8003"); err != nil {
		log.Fatalf("HTTP server failed to start: %s", err)
	}
}
