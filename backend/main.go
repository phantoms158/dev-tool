package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phantoms158/gin-bookstore/controllers"
	"github.com/phantoms158/gin-bookstore/databases"
	"github.com/phantoms158/gin-bookstore/middlewares"
)

func main() {
	r := gin.Default()
	r.Use(CORSMiddleware())
	databases.ConnectDatabase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Hello World"})
	})

	r.POST("/register", controllers.Register)

	r.POST("/login", controllers.Login)


	auth := r.Group("/auth")
	
	auth.Use(middlewares.AuthorizeJWT()) 
	{
		// auth.GET("/books", controllers.FindBooks)
		// auth.POST("/books", controllers.CreateBook)
		// auth.GET("/books/:id", controllers.FindBook)
		// auth.PATCH("/books/:id", controllers.UpdateBook)
		// auth.DELETE("/books/:id", controllers.DeleteBook)
	}

	
	r.Run()
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