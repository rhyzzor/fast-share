package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func UploadFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, handler, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(400, gin.H{
				"error": "Bad request",
			})
			return
		}
		defer file.Close()

		fmt.Print("File uploaded: ", file)

		c.JSON(200, gin.H{
			"message":  "File uploaded successfully",
			"filename": handler.Filename,
		})
	}
}
