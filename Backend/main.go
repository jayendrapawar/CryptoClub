package main

import (
	"CryptoClub/config"
	"CryptoClub/fetch"
	"CryptoClub/models"
	"CryptoClub/store"
	"fmt"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

const pollInterval = 50 * time.Second // polling every 10 seconds

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file", err)
		return
	}

	cfg := config.LoadConfig()

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
