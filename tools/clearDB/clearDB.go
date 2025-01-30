package main

import (
	"log"

	"github.com/adasarpan404/urlshortner/database"
	"github.com/adasarpan404/urlshortner/models"
)

func main() {
	database.InitDatabase()
	db := database.DB

	if err := db.Where("1 = 1").Delete(&models.URL{}).Error; err != nil {
		log.Fatalf("Failed to clear URL table: %v", err)
	}

	if err := db.Where("1 = 1").Delete(&models.User{}).Error; err != nil {
		log.Fatalf("Failed to clear User table: %v", err)
	}

	log.Println("All data cleared from tables successfully!")
}
