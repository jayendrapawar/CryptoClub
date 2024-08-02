package config

import "os"

type Config struct {
	APIURL         string
	APIKey         string
	MongoURI       string
	DBName         string
	CollectionName string
}

func LoadConfig() Config {
	return Config{
		APIURL:         os.Getenv("API_URL"),
		APIKey:         os.Getenv("API_KEY"),
		MongoURI:       os.Getenv("MONGO_URI"),
		DBName:         os.Getenv("DB_NAME"),
		CollectionName: os.Getenv("COLLECTION_NAME"),
	}
}
