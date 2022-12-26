package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	neo4jclient "lazydependency/neo4j_client"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func CreateProject(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Error getting file: %s", err),
		})
		return
	}

	if file.Filename != "package.json" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "File must be named package.json",
		})
		return
	}

	ext := filepath.Ext(file.Filename)
	dst := "./uploads/" + uuid.New().String() + ext

	if err = c.SaveUploadedFile(file, dst); err != nil {
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

	byteValue, _ := io.ReadAll(jsonFile)

	var result gin.H
	json.Unmarshal([]byte(byteValue), &result)

	if err = os.Remove(dst); err != nil {
		log.Printf("Error removing file: %s", err)
	}

	dependencyList := result["dependencies"]

	if isNilFixed(dependencyList) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "No dependencies found in package.json",
		})
		return
	}

	uuid := uuid.New().String()

	query := "CREATE (s:Project {uuid: $pid}) "

	for k, v := range dependencyList.(map[string]interface{}) {
		id := randSeq(10)
		query += fmt.Sprintf("MERGE (:Dependency {name: '%s'})-[:HAS]->(%s:Version {name: '%s'}) ", k, id, v)
		query += fmt.Sprintf("MERGE (%s)-[:IN_PROJECT]->(s) ", id)
	}

	session := neo4jclient.DB.NewSession(c, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(c)

	_, err = session.ExecuteWrite(c, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		_, err := tx.Run(c, query, gin.H{
			"pid": uuid,
		})
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

func ListProjects(c *gin.Context) {
}

func GetProject(c *gin.Context) {
}

func UpdateProject(c *gin.Context) {
	pid := c.Param("id")

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Error getting file: %s", err),
		})
		return
	}

	if file.Filename != "package.json" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "File must be named package.json",
		})
		return
	}

	ext := filepath.Ext(file.Filename)
	dst := "./uploads/" + uuid.New().String() + ext

	if err = c.SaveUploadedFile(file, dst); err != nil {
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

	byteValue, _ := io.ReadAll(jsonFile)

	var result gin.H
	json.Unmarshal([]byte(byteValue), &result)

	if err = os.Remove(dst); err != nil {
		log.Printf("Error removing file: %s", err)
	}

	removeProjectQuery := "MATCH (s:Project {uuid: $pid}) DETACH DELETE s"

	dependencyList := result["dependencies"]

	if isNilFixed(dependencyList) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "No dependencies found in package.json",
		})
		return
	}

	createProjectQuery := "CREATE (s:Project {uuid: $pid}) "

	for k, v := range dependencyList.(map[string]interface{}) {
		id := randSeq(10)
		createProjectQuery += fmt.Sprintf("MERGE (:Dependency {name: '%s'})-[:HAS]->(%s:Version {name: '%s'}) ", k, id, v)
		createProjectQuery += fmt.Sprintf("MERGE (%s)-[:IN_PROJECT]->(s) ", id)
	}

	session := neo4jclient.DB.NewSession(c, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(c)

	_, err = session.ExecuteWrite(c, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		_, err := tx.Run(c, removeProjectQuery, gin.H{
			"pid": pid,
		})
		if err != nil {
			return nil, err
		}

		_, err = tx.Run(c, createProjectQuery, gin.H{
			"pid": pid,
		})

		if err != nil {
			return nil, err
		}

		return nil, nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Error updating project: %s", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": result["dependencies"],
	})
}

func DeleteProject(c *gin.Context) {
	pid := c.Param("id")

	query := "MATCH (s:Project {uuid: $pid}) DETACH DELETE s"

	session := neo4jclient.DB.NewSession(c, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(c)

	_, err := session.ExecuteWrite(c, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		_, err := tx.Run(c, query, gin.H{
			"pid": pid,
		})
		if err != nil {
			return nil, err
		}
		return nil, nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Error deleting project: %s", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Project deleted",
	})
}

func GetReccomendations(c *gin.Context) {
}
