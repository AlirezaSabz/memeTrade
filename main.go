package main

import (
	"log"

	"github.com/joho/godotenv"
	"go.mod/internal/infrastructure/telegram"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	telegram.Start()
}
