package config

import (
	"github.com/joho/godotenv"
	"log"
	"net/url"
	"os"
	"strconv"
)

type Config struct {
	BackendURL *url.URL
	ServerPort int
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env files found")
	}

	rawURL := os.Getenv("BACKEND_URL")
	if rawURL == "" {
		log.Fatal("BACKEND_URL is not defined in the .env file")
	}

	backendURL, err := url.Parse(rawURL)
	if err != nil {
		log.Fatalf("BACKEND_URL parsing error: %v", err)
	}

	rawPort := os.Getenv("PORT")
	if rawPort == "" {
		log.Fatal("PORT is not defined in the .env file")
	}

	port, err := strconv.Atoi(rawPort)
	if err != nil {
		log.Fatalf("PORT parsing error: %v", err)
	}

	return &Config{
		BackendURL: backendURL,
		ServerPort: port,
	}, nil
}
