package main

import (
	"log"
	"net/http"

	"github.com/adasarpan404/urlshortner/database"
	"github.com/adasarpan404/urlshortner/models"
	"github.com/adasarpan404/urlshortner/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

	r.GET("/:shortCode", func(ctx *gin.Context) {
		shortCode := ctx.Param("shortCode")
		if shortCode == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Short code not provided"})
			return
		}

		var url models.URL
		if err := db.Where("short_code = ?", shortCode).First(&url).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "Short URL not found"})
			} else {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			}
			return
		}
		ctx.Redirect(http.StatusFound, url.LongURL)
	})

	log.Println("Server is running on http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
