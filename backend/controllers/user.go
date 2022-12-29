package controllers

import (
	"lazydependency/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
