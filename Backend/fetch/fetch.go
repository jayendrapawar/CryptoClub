package fetch

import (
	"CryptoClub/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func FetchCryptoData(apiURL, apiKey string) ([]models.CryptoData, error) {
	payload := strings.NewReader(`{
		"currency": "INR",
		"sort": "rank",
		"order": "ascending",
		"offset": 0,
		"limit": 2,
		"meta": false
	}`)

	client := &http.Client{}
	req, err := http.NewRequest("POST", apiURL, payload)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Add("content-type", "application/json")
	req.Header.Add("x-api-key", apiKey)

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var response []struct {
		Code string  `json:"code"`
		Rate float64 `json:"rate"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	var data []models.CryptoData
	for _, item := range response {
		data = append(data, models.CryptoData{
			Currency: item.Code,
			Price:    item.Rate,
		})
	}

	for _, d := range data {
		fmt.Printf("Fetched - Currency: %s, Price: %f\n", d.Currency, d.Price)
	}

	return data, nil
}
