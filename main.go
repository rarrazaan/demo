package main

import (
	"net/http"

	"github.com/gin-contrib/cors" // Import the gin-contrib/cors middleware
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Create a Gin router
	// gin.Default() includes a logger and recovery middleware
	r := gin.Default()

	// 2. Set up CORS middleware
	// We'll use the gin-contrib/cors package.
	// For this demo, we allow all origins ("*") to make testing simple.
	config := cors.DefaultConfig()
	// config.AllowAllOrigins = true
	// You could also be specific like your original example:
	config.AllowOrigins = []string{"https://demo-frontend.vercel.app", "http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Content-Type", "Authorization"}
	config.AllowCredentials = true
	
	r.Use(cors.New(config))

	// 3. Define your API routes
	
	// /api/hello endpoint
	r.GET("/api/hello", func(c *gin.Context) {
		// c.JSON simplifies sending JSON responses
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello! Your Gin backend is running on Render.com!",
		})
	})

	// / (root) endpoint for Render's health checks
	r.GET("/", func(c *gin.Context) {
		// c.String simplifies sending a plain text response
		c.String(http.StatusOK, "Gin backend is healthy and running...")
	})

	// 4. Start the server
	// Gin's r.Run() is smart! It will automatically listen on
	// the PORT environment variable if Render sets it.
	// If not, it defaults to 8080.
	r.Run()
}