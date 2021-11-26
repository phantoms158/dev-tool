package helpers

import "github.com/gin-gonic/gin"

// This will help in handling error response
func ResponseWithError(c *gin.Context, code int, message interface{}) {
    c.AbortWithStatusJSON(code, message)
	return
}

// This will help in handling error response
func ResponseWithSuccess(c *gin.Context, code int, message interface{}) {
    c.JSON(code, message)
	return
}