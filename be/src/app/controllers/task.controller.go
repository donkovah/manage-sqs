package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTask(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Read one from Tasks"})
}

func GetTasks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Read all from Tasks"})
}

func CreateTask(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "POST to Tasks"})
}

func UpdateTask(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "PATCH one fromTasks"})
}

func DeleteTask(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "DELETE one from Tasks"})
}
