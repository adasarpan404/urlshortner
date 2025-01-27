package main

import (
	"log"

	"github.com/adasarpan404/urlshortner/database"
	"github.com/adasarpan404/urlshortner/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	database.InitDatabase()
	db := database.DB

	err := db.AutoMigrate(&models.URL{}, &models.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	r.POST("/shorten", func(ctx *gin.Context) {
		var req struct {
			URL string `json:"url"`
		}
	})
}
