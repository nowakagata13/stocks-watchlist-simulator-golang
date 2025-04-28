package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Stock struct {
	Symbol        string  `json:"symbol"`
	CurrentPrice  float64 `json:"current_price"`
	ChangePercent float64 `json:"change_percent"`
}

func GetStockInfo(symbol string) (Stock, error) {
	apiKey := os.Getenv("FINNHUB_API_KEY")
	if apiKey == "" {
		return Stock{}, fmt.Errorf("missing FINNHUB_API_KEY environment variable")
	}

	url := fmt.Sprintf("https://finnhub.io/api/v1/quote?symbol=%s&token=%s", symbol, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return Stock{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Stock{}, fmt.Errorf("API request failed with status: %s", resp.Status)
	}

	var data struct {
		C  float64 `json:"c"`  // current price
		Pc float64 `json:"pc"` // previous close
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return Stock{}, err
	}

	// Calculate change percent safely
	var changePercent float64
	if data.Pc != 0 {
		changePercent = ((data.C - data.Pc) / data.Pc) * 100
	} else {
		changePercent = 0
	}

	return Stock{
		Symbol:        symbol,
		CurrentPrice:  data.C,
		ChangePercent: changePercent,
	}, nil
}
