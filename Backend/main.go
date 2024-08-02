package main

import (
	"CryptoClub/config"
	"CryptoClub/fetch"
	"CryptoClub/models"
	"CryptoClub/store"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const pollInterval = 10 * time.Second // Polling every 10 seconds

var cfg *config.Config // cfg is now a pointer

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}

	// Load configuration
	cfg = config.LoadConfig()

	// Initialize MongoDB client
	err = store.InitializeMongoClient(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize MongoDB client: %v", err)
	}

	// Start HTTP server in a separate goroutine
	go startServer()

	// Polling loop for fetching and storing data
	ticker := time.NewTicker(pollInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			prices, err := fetch.FetchCryptoData(cfg.APIURL, cfg.APIKey)
			if err != nil {
				fmt.Println("Error fetching data:", err)
				continue
			}

			var wg sync.WaitGroup
			for _, price := range prices {
				wg.Add(1)
				go func(p models.CryptoData) {
					defer wg.Done()
					err := store.StoreDataInMongo(p, cfg)
					if err != nil {
						fmt.Println("Error storing data in MongoDB:", err)
					}
				}(price)
			}

			wg.Wait()
		}
	}
}

func startServer() {
	r := gin.Default()

	// Configure CORS
	r.Use(cors.Default())

	// Define the /api/recent route
	r.GET("/api/recent", func(c *gin.Context) {
		recentData, err := store.FetchRecentData(cfg)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch recent data"})
			return
		}
		c.JSON(http.StatusOK, recentData)
	})

	// Start the server on port 8080
	fmt.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
