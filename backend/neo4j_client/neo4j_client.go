package neo4jclient

import (
	"context"
	"log"
	"os"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

var DB neo4j.DriverWithContext

func InitNeo4j() {
	createDriver(os.Getenv("NEO4J_URI"), os.Getenv("NEO4J_USER"), os.Getenv("NEO4J_PASS"))
}

func createDriver(uri, username, password string) {
	db, err := neo4j.NewDriverWithContext(uri, neo4j.BasicAuth(username, password, ""))

	if err != nil {
		panic(err.Error())
	} else {
		DB = db
	}

	ctx := context.Background()

	if err = db.VerifyConnectivity(ctx); err != nil {
		panic(err.Error())
	}
}

func CloseDriver() {
	log.Println("Closing Neo4j Database Driver")

	ctx := context.Background()

	err := DB.Close(ctx)
	if err != nil {
		panic(err.Error())
	} else {
		log.Println("Database Driver Closed")
	}
}

func CreateSession() neo4j.SessionWithContext {
	ctx := context.Background()
	return DB.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
}

func KillSession(session neo4j.SessionWithContext) {
	ctx := context.Background()

	err := session.Close(ctx)
	if err != nil {
		panic(err.Error())
	}
}
