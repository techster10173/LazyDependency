package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	mongoClient "lazydependency/mongo_client"
	neo4jClient "lazydependency/neo4j_client"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	mongoClient.InitMongoDB()
	neo4jClient.InitNeo4j()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-quit
		log.Println("Receive Interrupt Signal")
		mongoClient.CloseDriver()
		neo4jClient.CloseDriver()

		os.Exit(0)
	}()
}
