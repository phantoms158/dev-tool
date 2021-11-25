package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/phantoms158/gin-bookstore/helpers"
	"github.com/phantoms158/gin-bookstore/models"
)

const (
	// userkey = "user"
)

// POST /login
// login
func Login(c *gin.Context) {
	var input models.LoginCommand
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// // Login Information
	// user := models.User{Username: input.Username, Password: input.Password}
	// databases.DB.Create(&user)

	// c.JSON(http.StatusOK, gin.H{"data": user})
	token := models.GenerateToken(input.Username, true)
	if token != "" {
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	} else {
		c.JSON(http.StatusUnauthorized, nil)
	}
}

// POST /register
// register
func Register(c *gin.Context) {
	var input models.RegisterCommand
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// // Login Information
	user := models.User{Email: input.Email, Username: input.Username, Password: helpers.GeneratePasswordHash([]byte(input.Password)), CreatedAt:  time.Now(), UpdatedAt:  time.Now(),}
	// databases.DB.(&user)
	log.Printf("User ", user)
	c.JSON(http.StatusOK, gin.H{"data": user})
	// token := models.GenerateToken(input.Username, true)
	// if token != "" {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"token": token,
	// 	})
	// } else {
	// 	c.JSON(http.StatusUnauthorized, nil)
	// }
}

// POST /me
// me
// func me(c *gin.Context) {
	// session := sessions.Default(c)
	// user := session.Get(userkey)
	// c.JSON(http.StatusOK, gin.H{"user": user})
// }

// POST /status
// status
// func status(c *gin.Context) {
	// c.JSON(http.StatusOK, gin.H{"status": "You are logged in"})
// }