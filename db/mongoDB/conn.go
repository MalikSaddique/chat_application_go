package mongodb

import (
	"context"
	"fmt"
	"os"

	"github.com/MalikSaddique/chat_application_go/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	log = logger.Logger("Mongo-DB")
)

func MongoDbConn() (*mongo.Client, error) {
	url := os.Getenv("MONGODB_URI")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Ping(context.Background(), nil); err != nil {
		return nil, fmt.Errorf("database ping failed: %v", err)
	}

	log.Info("Connected to MongoDB with database/mongoDB!")

	return client, err
}
