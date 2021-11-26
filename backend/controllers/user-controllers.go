package controllers

import (
	"devtools-backend/forms"
	"devtools-backend/helpers"
	"devtools-backend/models"
	"devtools-backend/services"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// POST /login
// login
func Login(c *gin.Context) {

	var input forms.LoginCommand

	err := c.ShouldBindJSON(&input)
	fmt.Print(input)
	if err != nil {
		helpers.ResponseWithError(c, http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	user := models.GetUserByEmail(input.Email)

    if user.Email == "" {
		helpers.ResponseWithError(c, http.StatusNotFound, gin.H{"message": "User account was not found"})
    }
	
	// Get the hashed password from the saved document
    hashedPassword := []byte(user.Password)
    // Get the password provided in the request.body
    password := []byte(input.Password)

    err = helpers.PasswordCompare(password, hashedPassword)

    if err != nil {
		helpers.ResponseWithError(c, http.StatusForbidden, gin.H{"message": "Invalid user credentials"})
    }

    token, err2 := services.GenerateToken(input.Email)

    // If we fail to generate token for access
    if err2 != nil {
		helpers.ResponseWithError(c, http.StatusForbidden, gin.H{"message": "There was a problem logging you in, try again later"})
    }
	
	if token != "" {
		models.WriteAccessLog(c, input.Email)

		helpers.ResponseWithSuccess(c, http.StatusOK, gin.H{
			"message": "Log in success",
			"token": token,
		})
	} else {
		helpers.ResponseWithError(c, http.StatusUnauthorized, nil)
	}
}

// POST /register
// register
func Register(c *gin.Context) {
	var input forms.RegisterCommand
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	result := models.GetUserByEmail(input.Email)

	if result.Email != "" {
		helpers.ResponseWithError(c, http.StatusForbidden, gin.H{"message": "Email is already in use"})
	}

	err = models.Signup(input)

    if err != nil {
		helpers.ResponseWithError(c, http.StatusBadRequest, gin.H{"message": "Problem creating an account"})
    }

    helpers.ResponseWithSuccess(c, http.StatusCreated, gin.H{"message": "New user account registered"})
}

// POST /me
// me
func Me(c *gin.Context) {
	// session := sessions.Default(c)
	result, isExists := c.Get("user_data")
	user := result.(jwt.MapClaims)
	if(isExists) {
		fmt.Printf("@@@@@@@@@@@@@@@@@@ %s\n",user["email"])
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

// POST /status
// status
// func status(c *gin.Context) {
	// c.JSON(http.StatusOK, gin.H{"status": "You are logged in"})
// }