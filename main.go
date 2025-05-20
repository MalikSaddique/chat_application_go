package main

// @title Go Chat App
// @version 1.0
// @description This is a chat application API built with Go and Gin.
// @host localhost:8003
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
import (
	"os"

	authserviceimpl "github.com/MalikSaddique/chat_application_go/controllers/auth_service/auth_service_impl"
	messageserviceimpl "github.com/MalikSaddique/chat_application_go/controllers/message_service/message_service_impl"
	mongodb "github.com/MalikSaddique/chat_application_go/db/mongoDB"
	db "github.com/MalikSaddique/chat_application_go/db/postgresDB"
	"github.com/MalikSaddique/chat_application_go/pkg/logger"
	"github.com/MalikSaddique/chat_application_go/router"
	websocketclient "github.com/MalikSaddique/chat_application_go/web_socket_client"

	websocketsimpl "github.com/MalikSaddique/socket/websockets/websocketsimpl"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var log = logger.Logger("ChatApp")

func main() {
	log.Infof("App started")

	err := godotenv.Load(".env")
	key := os.Getenv("BACKEND_WS_KEY")
	if err != nil {
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

	webSockets := websocketsimpl.NewWebSockets(messagedb)
	messageService := messageserviceimpl.NewMessageService(messagedb, webSockets)

	go websocketclient.ConnectToWebSocketServer("ws://localhost:8004/backend/ws", key)

	httpRouter := router.NewRouter(authService, messageService)
	if err := httpRouter.Engine.Run(":8003"); err != nil {
		log.Fatalf("HTTP server failed to start: %s", err)
	}

}
