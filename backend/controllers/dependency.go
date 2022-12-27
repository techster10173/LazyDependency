package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	neo4jclient "lazydependency/neo4j_client"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func GetDependency(c *gin.Context) {
}

func LikeDependency(c *gin.Context) {
}

func UnlikeDependency(c *gin.Context) {
}

// This adds a new version node to a given dependency node (if it doesnt already exist)
func AddVersionNode(c *gin.Context) {
	// TODO Implement
}

// this creates a new edge between two dependency/version combos
func CreateEdge(c *gin.Context) {
	// TODO Implement
}

// this gets the edge id between two dependency/versions and then gets the comments for this edge
func GetEdge(c *gin.Context) {
	// TODO Implement
}

// this adds a comment for a given edge (the comments consists of a comment, account/project id, and a 0 or 1 for a check or x)
func AddComment(c *gin.Context) {
	// TODO Implement
}

func UploadConnections(c *gin.Context) {
	//fmt.Printf("Starting server at port 8080")
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Error getting file: %s", err),
		})
		return
	}

	ext := filepath.Ext(file.Filename)
	dst := "./uploads/" + uuid.New().String() + ext

	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Error saving file: %s", err),
		})
		return
	}

	jsonFile, err := os.Open(dst)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Error opening file: %s", err),
		})
		return
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result gin.H
	json.Unmarshal([]byte(byteValue), &result)

	if err = os.Remove(dst); err != nil {
		log.Printf("Error removing file: %s", err)
	}

	uuid := uuid.New().String()

	query := "CREATE (s:STACK {uuid: '" + uuid + "'}) "

	for k, v := range result["dependencies"].(map[string]interface{}) {
		id := randSeq(10)
		newVrsn := newVersion(k, v.(string))
		log.Printf(newVrsn)
		log.Printf(k)
		query += "MERGE (:Dependency {name: '" + k + "'})-[:HAS]->(" + id + ":Version {name: '" + newVrsn + "'}) "
		query += "MERGE (" + id + ")-[:IN_STACK]->(s) "
	}

	//i made some code

	session := neo4jclient.DB.NewSession(c, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(c)

	_, err = session.ExecuteWrite(c, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		_, err := tx.Run(c, query, gin.H{})
		if err != nil {
			return nil, err
		}
		return nil, nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Error creating dependencies: %s", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": result["dependencies"],
	})
}
