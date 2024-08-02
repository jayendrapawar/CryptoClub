package config

import (
	"log"
	"os"
)

// Config holds configuration values
type Config struct {
	APIURL         string
	APIKey         string
	MongoURI       string
	DBName         string
	CollectionName string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	cfg := &Config{
		APIURL:         getEnv("API_URL", "https://default-api-url.com"),
		APIKey:         getEnv("API_KEY", ""),
		MongoURI:       getEnv("MONGO_URI", "https://localhost-mongodb:57921"),
		DBName:         getEnv("DB_NAME", "defaultDB"),
		CollectionName: getEnv("COLLECTION_NAME", "defaultCollection"),
	}

	// Validate required fields
	if cfg.APIKey == "" {
		log.Fatal("API_KEY environment variable is required")
	}

	return cfg
}

// getEnv retrieves environment variables with optional default values
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
