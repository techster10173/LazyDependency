package controllers

import "github.com/gin-gonic/gin"

// This gets a list of the versions that exist in the graph for this dependency
func GetDependencyNode(c *gin.Context) {
	// TODO Implement
}

// this gets a list of what other dependencies we have compability data on for this version (just what edges exist)
func GetVersionNode(c *gin.Context) {
	// TODO Implement
}

// this adds a new dependency node (if it doesnt already exist)
func AddDependencyNode(c *gin.Context) {
	// TODO Implement
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
