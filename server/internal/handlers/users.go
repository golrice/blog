// handlers/user.go
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserList(c *gin.Context) {
	// 获取用户列表的处理逻辑
	c.JSON(http.StatusOK, gin.H{
		"message": "Getting user list",
	})
}

func CreateUser(c *gin.Context) {
	// 创建用户的处理逻辑
	c.JSON(http.StatusCreated, gin.H{
		"message": "User created",
	})
}
