package main

import (
	"devtools-backend/controllers"
	"devtools-backend/databases"
	"devtools-backend/middlewares"
	"devtools-backend/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// init gets called before the main function
func init() {
	// Log error if .env file does not exist
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}
}

func main() {
	// Init gin router
	r := gin.Default()
	r.Use(CORSMiddleware())
	r.Use(middlewares.Logger())
	databases.ConnectDataBase()
	databases.DB.AutoMigrate(&models.User{}, &models.AccessLog{}, &models.ActivityLog{})
	// Its great to version your API's
	v1 := r.Group("/api/v1")
	{
		// Define a GET request to call the Default
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"data": "Hello World"})
		})

		v1.POST("/register", controllers.Register)

		v1.POST("/login", controllers.Login)

		auth := v1.Group("/auth")

		auth.Use(middlewares.AuthorizeJWT())
		{
			auth.GET("/me", controllers.Me)
			// auth.POST("/books", controllers.CreateBook)
			// auth.GET("/books/:id", controllers.FindBook)
			// auth.PATCH("/books/:id", controllers.UpdateBook)
			// auth.DELETE("/books/:id", controllers.DeleteBook)
		}
	}

	// Handle error response when a route is not defined
	r.NoRoute(func(c *gin.Context) {
		// In gin this is how you return a JSON response
		c.JSON(404, gin.H{"message": "Not found"})
	})

	r.Run(":8080")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
