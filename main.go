package main

import (
	"file-service/initializers"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

func init() {
	initializers.LoadEnvFile()
}

func main() {
	r := gin.Default()
	r.Static("/app", "./templates")
	r.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		filePath := filepath.Base(file.Filename)
		err := c.SaveUploadedFile(file, filePath)
		if err != nil {
			return
		}
		c.String(http.StatusOK, "File uploaded successfully")
	})
	err := r.Run()
	if err != nil {
		return
	}
}
