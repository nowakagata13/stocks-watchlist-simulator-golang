package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize the database
	initDB()

	// Load and display the initial watchlist from the database
	stocks := getStocksFromDB()
	fmt.Println("Current watchlist:")
	for _, stock := range stocks {
		fmt.Println(stock)
	}

	// Start the watchlist refresh in the background (could be used for price fetching)
	go refreshWatchlist()

	// Initialize scanner for user input
	scanner := bufio.NewScanner(os.Stdin)

	// Main program loop
	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. List Watchlist")
		fmt.Println("2. Add Stock")
		fmt.Println("3. Remove Stock")
		fmt.Println("4. Buy Stock")
		fmt.Println("5. Sell Stock")
		fmt.Println("6. View Portfolio")
		fmt.Println("7. View Transactions")
		fmt.Println("8. Quit")
		fmt.Print("> ")

		// Read user input
		if !scanner.Scan() {
			break
		}
		choice := scanner.Text()

		// Handle user choices
		switch choice {
		case "1":
			// Display the current watchlist
			listWatchlist()
		case "2":
			// Add stock to watchlist
			fmt.Print("Enter stock symbol to add: ")
			if scanner.Scan() {
				symbol := scanner.Text()
				addStock(symbol)
			}
		case "3":
			// Remove stock from watchlist
			fmt.Print("Enter stock symbol to remove: ")
			if scanner.Scan() {
				symbol := scanner.Text()
				removeStock(symbol)
			}
		case "4":
			// Buy stock
			fmt.Print("Enter stock symbol to buy: ")
			if scanner.Scan() {
				ticker := scanner.Text()
				fmt.Print("Enter quantity: ")
				if scanner.Scan() {
					qtyStr := scanner.Text()
					qty, err := strconv.Atoi(qtyStr)
					if err != nil {
						fmt.Println("Invalid quantity.")
						break
					}
					price := fetchPrice(ticker)
					buyStock(ticker, qty, price)
					addTransactionToDB("Buy", ticker, qty, price)
					fmt.Printf("Bought %d of %s at $%.2f\n", qty, ticker, price)
				}
			}
		case "5":
			// Sell stock
			fmt.Print("Enter stock symbol to sell: ")
			if scanner.Scan() {
				ticker := scanner.Text()
				fmt.Print("Enter quantity: ")
				if scanner.Scan() {
					qtyStr := scanner.Text()
					qty, err := strconv.Atoi(qtyStr)
					if err != nil {
						fmt.Println("Invalid quantity.")
						break
					}
					price := fetchPrice(ticker)
					sellStock(ticker, qty, price)
					addTransactionToDB("Sell", ticker, qty, price)
					fmt.Printf("Sold %d of %s at $%.2f\n", qty, ticker, price)
				}
			}
		case "6":
			// View portfolio
			viewPortfolio()
		case "7":
			// View transaction history
			viewTransactions()
		case "8":
			// Exit the program
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}
