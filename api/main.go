package main

import "github.com/gin-gonic/gin"

func main() {
	
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// Route with parameters
	r.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{"greeting": "Hello " + name})
	})

	// POST route
	r.POST("/submit", func(c *gin.Context) {
		var json struct {
			Name string `json:"name"`
		}
		if err := c.ShouldBindJSON(&json); err == nil {
			c.JSON(200, gin.H{"status": "received", "name": json.Name})
		} else {
			c.JSON(400, gin.H{"error": err.Error()})
		}
	})

	r.Run(":8080") // Starts server on http://localhost:8080
}
