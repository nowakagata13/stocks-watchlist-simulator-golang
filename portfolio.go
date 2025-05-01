package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// Portfolio holds the stocks and transactions
var portfolio = Portfolio{
	Holdings:     make(map[string]int),
	Transactions: []Transaction{},
}

// Transaction represents a buy/sell action
type Transaction struct {
	Type     string  // Buy or Sell
	Ticker   string  // Stock symbol
	Quantity int     // Number of stocks
	Price    float64 // Price per stock
}

// StockPrice represents the stock price data returned from the API
type StockPrice struct {
	CurrentPrice float64 `json:"c"` // 'c' is the field for current price in the Finnhub response
}

// Portfolio represents the user's stock holdings
type Portfolio struct {
	Holdings     map[string]int // Stock symbol -> quantity
	Transactions []Transaction  // List of transactions
}

// buyStock buys a certain quantity of stock
func buyStock(ticker string, qty int, price float64) {
	if portfolio.Holdings == nil {
		portfolio.Holdings = make(map[string]int)
	}
	portfolio.Holdings[ticker] += qty
	portfolio.Transactions = append(portfolio.Transactions, Transaction{"Buy", ticker, qty, price})
}

// sellStock sells a certain quantity of stock
func sellStock(ticker string, qty int, price float64) {
	if portfolio.Holdings[ticker] < qty {
		fmt.Println("Not enough shares to sell.")
		return
	}
	portfolio.Holdings[ticker] -= qty
	portfolio.Transactions = append(portfolio.Transactions, Transaction{"Sell", ticker, qty, price})
}

// viewPortfolio displays the user's stock holdings
func viewPortfolio() {
	fmt.Println("\n=== Portfolio ===")
	for ticker, qty := range portfolio.Holdings {
		fmt.Printf("%s: %d shares\n", ticker, qty)
	}
}

// viewTransactions displays all buy/sell transactions
func viewTransactions() {
	fmt.Println("\n=== Transaction History ===")
	for _, t := range portfolio.Transactions {
		fmt.Printf("%s %d of %s at $%.2f\n", t.Type, t.Quantity, t.Ticker, t.Price)
	}
}

// fetchPrice retrieves the current price of a stock
func fetchPrice(symbol string) float64 {
	// Load API key from the environment file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("FINNHUB_API_KEY")
	if apiKey == "" {
		log.Fatal("API key is not set.")
	}

	// Construct the API request URL
	url := fmt.Sprintf("https://finnhub.io/api/v1/quote?symbol=%s&token=%s", symbol, apiKey)

	// Make the HTTP GET request to fetch the stock price
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error fetching stock price: %v", err)
	}
	defer resp.Body.Close()

	// Parse the response JSON
	var stockPrice StockPrice
	if err := json.NewDecoder(resp.Body).Decode(&stockPrice); err != nil {
		log.Fatalf("Error decoding stock price data: %v", err)
	}

	// Return the current price
	return stockPrice.CurrentPrice
}
