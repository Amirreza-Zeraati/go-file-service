package main

import (
	"file-service/initializers"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func init() {
	initializers.LoadEnvFile()
}

func main() {
	r := gin.Default()
	r.Static("/app", "./templates")

	r.POST("/upload-chunk", func(c *gin.Context) {
		fileId := c.PostForm("fileId")
		fileName := c.PostForm("fileName")
		chunkIndexStr := c.PostForm("chunkIndex")
		totalChunksStr := c.PostForm("totalChunks")

		chunkIndex, err := strconv.Atoi(chunkIndexStr)
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid chunk index")
			return
		}

		totalChunks, err := strconv.Atoi(totalChunksStr)
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid total chunks")
			return
		}

		tempDir := filepath.Join("uploads", fileId)
		err = os.MkdirAll(tempDir, os.ModePerm)
		if err != nil {
			return
		}

		file, err := c.FormFile("chunk")
		if err != nil {
			c.String(http.StatusBadRequest, "Chunk upload error")
			return
		}

		chunkPath := filepath.Join(tempDir, fmt.Sprintf("chunk_%d", chunkIndex))
		err = c.SaveUploadedFile(file, chunkPath)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to save chunk")
			return
		}

		files, _ := os.ReadDir(tempDir)
		if len(files) == totalChunks {
			baseName := filepath.Base(fileName)
			ext := filepath.Ext(baseName)
			nameOnly := baseName[:len(baseName)-len(ext)]

			finalPath := filepath.Join("uploads", baseName)
			i := 1
			for {
				if _, err := os.Stat(finalPath); os.IsNotExist(err) {
					break
				}
				finalPath = filepath.Join("uploads", fmt.Sprintf("%s(%d)%s", nameOnly, i, ext))
				i++
			}

			destFile, err := os.Create(finalPath)
			if err != nil {
				c.String(http.StatusInternalServerError, "Failed to create final file")
				return
			}
			defer func(destFile *os.File) {
				err := destFile.Close()
				if err != nil {

				}
			}(destFile)

			for i := 0; i < totalChunks; i++ {
				chunkPath := filepath.Join(tempDir, fmt.Sprintf("chunk_%d", i))
				chunkFile, err := os.Open(chunkPath)
				if err != nil {
					c.String(http.StatusInternalServerError, "Failed to open chunk")
					return
				}
				_, err = io.Copy(destFile, chunkFile)
				if err != nil {
					return
				}
				err = chunkFile.Close()
				if err != nil {
					return
				}
				err = os.Remove(chunkPath)
				if err != nil {
					return
				}
			}
			err = os.Remove(tempDir)
			if err != nil {
				return
			}
		}
		c.String(http.StatusOK, "Chunk uploaded")
	})
	err := r.Run()
	if err != nil {
		return
	}
}
