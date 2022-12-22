package main

import (
	"flag"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"lazydependency/controllers"
	mongoClient "lazydependency/mongo_client"
	neo4jClient "lazydependency/neo4j_client"
	router "lazydependency/router"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	prod := flag.Bool("prod", false, "production-release mode")

	if !(*prod) {
		if err := godotenv.Load(); err != nil {
			log.Println("No .env file found")
		}
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	mongoClient.InitMongoDB()
	neo4jClient.InitNeo4j()
	controllers.Init()

	rand.Seed(time.Now().UnixNano())

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
