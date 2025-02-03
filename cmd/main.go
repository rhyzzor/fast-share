package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rhyzzor/fast-share/internal/config"
	"github.com/rhyzzor/fast-share/internal/database"
)

func main() {
	r := gin.Default()
	cfg := config.LoadConfig()

	db, err := database.Connect(cfg.MongoURI)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer database.Disconnect(db.Client())

	r.Run(":" + cfg.Port)
}

func setupRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	return r
}
