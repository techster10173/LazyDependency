package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"lazydependency/controllers"
	mongoClient "lazydependency/mongo_client"
	neo4jClient "lazydependency/neo4j_client"
	router "lazydependency/router"

	"github.com/joho/godotenv"
)

func main() {
	prod := flag.Bool("prod", false, "production-release mode")

	if !(*prod) {
		if err := godotenv.Load(); err != nil {
			log.Println("No .env file found")
		}
	}

	mongoClient.InitMongoDB()
	neo4jClient.InitNeo4j()
	controllers.Init()

	port, exists := os.LookupEnv("PORT")
	if !exists || port == "" {
		port = "80"
	}

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router.InitServer(prod),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}

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
