package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"math/rand"
	"time"

	"github.com/adasarpan404/urlshortner/database"
	"github.com/adasarpan404/urlshortner/models"
	"gorm.io/gorm"
)

func main() {
	fmt.Printf("Testing Arpan Das")

	database.InitDatabase()
	db := database.DB

	err := db.AutoMigrate(&models.URL{}, &models.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/shorten", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		var req struct {
			URL string `json:"url"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.URL == "" {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		shortCode := generateShortCode()

		// Save to database
		url := &models.URL{ShortCode: shortCode, LongURL: req.URL}
		if err := db.Create(url).Error; err != nil {
			http.Error(w, "Failed to save URL", http.StatusInternalServerError)
			return
		}

		resp := map[string]string{
			"short_url": "http://localhost:8080/" + shortCode,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})

	// Endpoint: Redirect
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		shortCode := r.URL.Path[1:]
		if shortCode == "" {
			http.Error(w, "Short code not provided", http.StatusBadRequest)
			return
		}

		// Fetch from database
		var url models.URL
		if err := db.Where("short_code = ?", shortCode).First(&url).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				http.Error(w, "Short URL not found", http.StatusNotFound)
			} else {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
			return
		}

		http.Redirect(w, r, url.LongURL, http.StatusFound)
	})

	log.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func generateShortCode() string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	code := make([]byte, 6)
	for i := range code {
		code[i] = chars[rand.Intn(len(chars))]
	}
	return string(code)
}
