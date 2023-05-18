package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Define a route for the GET request to /api/hello
	router.GET("/api/hello", helloHandler)

	// Run the server on port 8080
	router.Run(":8080")
}

func helloHandler(c *gin.Context) {
	// Handle the GET request to /api/hello
	message := "Hello, World!"

	// Send the response as JSON
	c.JSON(http.StatusOK, gin.H{"message": message})
}

