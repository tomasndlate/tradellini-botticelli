package main

import (
	"github.com/gin-gonic/gin"

	"api/handlers"
)

func main() {
	r := gin.Default()

	r.POST("/start", func(c *gin.Context) {
		handlers.StartPrint("John Doe")
		c.JSON(200, gin.H{"message": "start"})
	})

	r.POST("/stop", func(c *gin.Context) {
		handlers.StopPrint()
		c.JSON(200, gin.H{"message": "stop"})
	})

	r.Run(":8080") // Starts server on http://localhost:8080
}
