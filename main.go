package main

import (
	"log"

	"github.com/adasarpan404/urlshortner/database"
	"github.com/adasarpan404/urlshortner/routes"
)

func main() {
	database.InitDatabase()

	r := routes.SetupRouter()

	log.Println("Server is running on http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
