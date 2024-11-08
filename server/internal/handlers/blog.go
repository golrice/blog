// handlers/user.go
package handlers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
)

type ArticleMetaData struct {
	Title       string `json:"title"`
	Date        string `json:"date"`
	Description string `json:"description"`
}

func GetAllBlogs(c *gin.Context) {
	// 获取所有博客文章的处理逻辑
	c.JSON(http.StatusOK, gin.H{
		"message": "Getting all blogs",
	})
}

func GetBlogByTitle(c *gin.Context) {
	// 获取指定博客文章的处理逻辑
	files, err := os.ReadDir("articles")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No articles found",
		})
		return
	}

	// 获取需要的博客文章的标题
	title := c.Param("title")
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if file.Name() == title {
			f, err := os.Open("articles/" + file.Name())
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{
					"message": "Blog not found",
				})
				return
			}
			defer f.Close()

			content, err := io.ReadAll(f)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{
					"message": "Blog not found",
				})
				return
			}
			if strings.Contains(title, ".md") {
				content = blackfriday.Run(content)
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "Getting blog by title",
				"content": string(content),
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "Blog not found",
	})
}

func CreateBlog(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": "Blog created",
	})
}

func GetAllBlogMetaData(c *gin.Context) {
	files, err := os.ReadDir("articles")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No articles found",
		})
		return
	}

	var articles []ArticleMetaData
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		fileinfo, err := file.Info()
		if err != nil {
			continue
		}
		f, err := os.Open("articles/" + fileinfo.Name())
		if err != nil {
			continue
		}
		defer f.Close()
		content := make([]byte, 30)
		_, err = f.Read(content)
		if err != nil {
			continue
		}

		article := ArticleMetaData{
			Title:       fileinfo.Name(),
			Date:        fileinfo.ModTime().Format("2006-01-02"),
			Description: string(content),
		}
		articles = append(articles, article)
	}
	c.JSON(http.StatusOK, articles)
}
