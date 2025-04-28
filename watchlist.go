package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"
)

var watchlist []Stock

func loadWatchlist() {
	file, err := os.Open("watchlist.json")
	if err != nil {
		if os.IsNotExist(err) {
			watchlist = []Stock{}
			return
		}
		fmt.Println("Error loading watchlist:", err)
		return
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&watchlist)
	if err != nil {
		fmt.Println("Error decoding watchlist:", err)
	}
}

func saveWatchlist() {
	file, err := os.Create("watchlist.json")
	if err != nil {
		fmt.Println("Error saving watchlist:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(watchlist)
	if err != nil {
		fmt.Println("Error encoding watchlist:", err)
	}
}

func listWatchlist() {
	if len(watchlist) == 0 {
		fmt.Println("Watchlist is empty.")
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Symbol", "Current Price", "Change %"})

	for _, stock := range watchlist {
		price := fmt.Sprintf("%.2f", stock.CurrentPrice)
		change := fmt.Sprintf("%.2f%%", stock.ChangePercent)
		table.Append([]string{stock.Symbol, price, change})
	}

	table.Render()
}

func addStock(symbol string) {
	symbol = strings.ToUpper(symbol)

	for _, s := range watchlist {
		if s.Symbol == symbol {
			fmt.Println("Stock already in watchlist.")
			return
		}
	}

	stock, err := GetStockInfo(symbol)
	if err != nil {
		fmt.Println("Error fetching stock:", err)
		return
	}

	watchlist = append(watchlist, stock)
	saveWatchlist()
	fmt.Println("Added:", symbol)
}

func removeStock(symbol string) {
	symbol = strings.ToUpper(symbol)

	for i, s := range watchlist {
		if s.Symbol == symbol {
			watchlist = append(watchlist[:i], watchlist[i+1:]...)
			saveWatchlist()
			fmt.Println("Removed:", symbol)
			return
		}
	}

	fmt.Println("Stock not found in watchlist.")
}

func refreshWatchlist() {
	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()

	for {
		<-ticker.C

		for i, stock := range watchlist {
			updatedStock, err := GetStockInfo(stock.Symbol)
			if err != nil {
				fmt.Printf("Error updating %s: %v\n", stock.Symbol, err)
				continue
			}
			watchlist[i] = updatedStock
		}

		saveWatchlist()
	}
}
