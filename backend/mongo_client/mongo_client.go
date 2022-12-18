package mongoclient

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func InitMongoDB() {
	createDriver(os.Getenv("MONGODB_URI"))
}

func createDriver(uri string) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err.Error())
	}

	if err = client.Ping(context.Background(), nil); err != nil {
		panic(err.Error())
	}

	DB = client
}

func CloseDriver() {
	log.Println("Closing MongoDB Database Driver")

	ctx := context.Background()

	if err := DB.Disconnect(ctx); err != nil {
		panic(err.Error())
	} else {
		log.Println("Database Driver Closed")
	}
}
