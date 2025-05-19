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

	stoken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImhlbGxvMTFAZ21haWwuY29tIiwiZXhwIjoxNzQ3NjU5NDMyLCJ1c2VyX2lkIjoxfQ.6cv2pq8jc0szBymhIY1EuMBeo7DmcptooRnMw5-egvA"
	rtoken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImhlbGxvQGdtYWlsLmNvbSIsImV4cCI6MTc0NzY1OTU0MiwidXNlcl9pZCI6Mn0.ZDqR3UrdoTz8iHKuaNbxSpUVQ-pzmt6kMGJL6qBBHEA"
	go websocketclient.ConnectToWebSocketServer("ws://localhost:8004/protected/ws", stoken)
	go websocketclient.ConnectToWebSocketServer("ws://localhost:8004/protected/ws", rtoken)

	httpRouter := router.NewRouter(authService, messageService)
	if err := httpRouter.Engine.Run(":8003"); err != nil {
		log.Fatalf("HTTP server failed to start: %s", err)
	}

}
