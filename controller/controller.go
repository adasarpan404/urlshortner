package controller

import (
	"net/http"

	"github.com/adasarpan404/urlshortner/database"
	"github.com/adasarpan404/urlshortner/models"
	"github.com/adasarpan404/urlshortner/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ShortenUrl(c *gin.Context) {
	var req struct {
		URL string `json:"url"`
	}

	if err := c.ShouldBindJSON(&req); err != nil || req.URL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request"})
		return
	}

	shortCode := utils.GenerateShortCode()

	url := &models.URL{ShortCode: shortCode, LongURL: req.URL}

	if err := database.DB.Create(url).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed To save URL"})
		return
	}

	resp := map[string]string{
		"short_url": "http://localhost:8080/" + shortCode,
	}

	c.JSON(http.StatusOK, resp)
}

func GET(c *gin.Context) {
	shortCode := c.Param("shortCode")

	if shortCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Short Code not Provided"})
		return
	}

	var url models.URL

	if err := database.DB.Where("short_code = ?", shortCode).First(&url).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Short URL not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}
	c.Redirect(http.StatusFound, url.LongURL)
}
