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

func UploadConnections(c *gin.Context) {
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
		query += "MERGE (:Dependency {name: '" + k + "'})-[:HAS]->(" + id + ":Version {name: '" + v.(string) + "'}) "
		query += "MERGE (" + id + ")-[:IN_STACK]->(s) "
	}

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
