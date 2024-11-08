// server/routes.go
package server

import (
	"github.com/gin-gonic/gin"

	"github.com/golrice/blog/internal/handlers"
)

func SetupRoutes(r *gin.Engine) {
	// 用户路由组
	userGroup := r.Group("/user")
	{
		userGroup.GET("/", handlers.GetUserList)
		userGroup.POST("/", handlers.CreateUser)
	}

	// 博客文章路由组
	blogGroup := r.Group("/blog")
	{
		blogGroup.GET("/", handlers.GetAllBlogs)
		blogGroup.POST("/", handlers.CreateBlog)
		blogGroup.GET("/files", handlers.GetAllBlogMetaData)
		blogGroup.GET("/files/:title", handlers.GetBlogByTitle)
	}
}
