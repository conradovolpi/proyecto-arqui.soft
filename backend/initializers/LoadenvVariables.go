package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadEnvVariables loads environment variables from a .env file
func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	} else {
		log.Println("Environment variables loaded successfully")
	}
}
