package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rhyzzor/fast-share/internal/config"
	"github.com/rhyzzor/fast-share/internal/database"
	"github.com/rhyzzor/fast-share/internal/handlers"
)

func main() {
	r := gin.Default()
	cfg := config.LoadConfig()

	db, err := database.Connect(cfg.MongoURI)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer database.Disconnect(db.Client())

	r.POST("/upload", handlers.UploadFile())

	r.Run(":" + cfg.Port)
}
