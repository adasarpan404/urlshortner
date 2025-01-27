package main

import (
	"log"
	"net/http"

	"github.com/adasarpan404/urlshortner/database"
	"github.com/adasarpan404/urlshortner/models"
	"github.com/adasarpan404/urlshortner/utils"
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

		if err := ctx.ShouldBindJSON(&req); err != nil || req.URL == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		shortCode := utils.GenerateShortCode()

		url := &models.URL{ShortCode: shortCode, LongURL: req.URL}

		if err := db.Create(url).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save URL"})
			return
		}
		resp := map[string]string{
			"short_url": "http://localhost:8080/" + shortCode,
		}
		ctx.JSON(http.StatusOK, resp)
	})

	r.GET("/:shortCode")
}
