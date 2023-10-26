package simple_rest

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func RunGolangSnippet() {
	fmt.Println("I am inside the simple rest package")

	// Create a new Gin router
	router := gin.Default()

	// Use custom logger middleware
	router.Use(LoggerMiddleware())

	// Use our custom authentication middleware for a specific group of routes
	authGroup := router.Group("/api")
	authGroup.Use(AuthMiddleware())
	{
		// Define a route for the root URL
		router.GET("/", func(c *gin.Context) {
			c.String(200, "Hello, World!")
		})

		// Define a route for the bye URL
		router.GET("/bye", func(c *gin.Context) {
			c.String(200, "Goodbye, World!")
		})

		// Route with URL parameters
		userController := &UserController{}
		router.GET("users/:id", userController.GetUserInfo)

		// Route with query parameters
		router.GET("/search", func(c *gin.Context) {
			query := c.DefaultQuery("q", "default-value")
			c.String(200, "Search query: "+query)
		})
	}

	// Run the server on port 8080
	router.Run(":8080")
}

/*

	Middleware functions in Gin are essential components that intercept
	HTTP requests and responses. They can perform pre-processing tasks before a
	request reaches the designated route handler or post-processing task before the response
	is sent to the client.

	Gin provides built-in middleware functions for common functionalities, such as logging,
	CORS handling, and recovery from panics. Additionally, developers can create custom
	middleware to extend Gin's capabilities according to their specific project requirements.

*/

// Using built-in middleware
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		log.Printf("Request - Method: %s | Status %d | Duration: %v", c.Request.Method, c.Writer.Status(), duration)
	}
}

/*

	Developers often need to implement custom middleware for project-specific requirements. Custom
	middleware can handle tasks like authentication, data validation, rate limiting, and more.

*/

// Using custom middleware
func AuthMiddleware() gin.HandlerFunc {
	// In a real-wrold application, you would perform proper authentication here.
	// For the sake of this example, we'll just check if an API key is present.
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-Key")
		if apiKey == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}
		c.Next()
	}
}

// Controllers
type UserController struct{}

// GetUserInfo is a controller method to get user information
func (uc *UserController) GetUserInfo(c *gin.Context) {
	userID := c.Param("id")
	// Fetch user information from the database or other data source
	// For simplicity, we'll just return a JSON response.
	c.JSON(200, gin.H{"id": userID, "name": "John Doe", "email": "john@example.com"})
}
