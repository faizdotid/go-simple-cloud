package controllers

import (
	"github.com/gin-gonic/gin"
)

func IndexController(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to Go-simple-cloud",
	})
}
