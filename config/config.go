package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var OpenWeatherMapAPIKey string

// LoadConfig loads configuration from .env file
func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	OpenWeatherMapAPIKey = os.Getenv("OPENWEATHERMAP_API_KEY")
	if OpenWeatherMapAPIKey == "" {
		log.Fatalf("OPENWEATHERMAP_API_KEY is not set in the .env file")
	}
}
