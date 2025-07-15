package main

import (
	"file-service/handlers"
	"file-service/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvFile()
}

func main() {
	r := gin.Default()
	r.Static("/app", "./templates")
	r.POST("/upload-chunk", handlers.UploadChunkHandler)
	err := r.Run()
	if err != nil {
		return
	}
}
