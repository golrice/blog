// server/server.go
package server

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartServer(port string) {
	r := gin.Default()

	r.Use(cors.Default())
	SetupRoutes(r)

	if err := r.Run(":" + port); err != nil {
		fmt.Printf("服务器启动失败: %v\n", err)
	}
}
