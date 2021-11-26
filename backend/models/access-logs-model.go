package models

import (
	"devtools-backend/databases"
	"time"

	"github.com/gin-gonic/gin"
)

type AccessLog struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	ClientIP  string    `json:"client_ip"`
	Username  string    `json:"username"`
	UserAgent string    `json:"user_agent"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//  handles access log
func WriteAccessLog(ctx *gin.Context, username string) error {
	accessLog := AccessLog{
		ClientIP:  ctx.ClientIP(),
		Username:  username,
		UserAgent: ctx.Request.UserAgent(),
	}
	err := databases.DB.Create(&accessLog).Error
	return err

}
