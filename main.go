package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	loadWatchlist()

	go refreshWatchlist()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. List Watchlist")
		fmt.Println("2. Add Stock")
		fmt.Println("3. Remove Stock")
		fmt.Println("4. Quit")
		fmt.Print("> ")

		if !scanner.Scan() {
			break
		}
		choice := scanner.Text()

		switch choice {
		case "1":
			listWatchlist()
		case "2":
			fmt.Print("Enter stock symbol to add: ")
			if scanner.Scan() {
				symbol := scanner.Text()
				addStock(symbol)
			}
		case "3":
			fmt.Print("Enter stock symbol to remove: ")
			if scanner.Scan() {
				symbol := scanner.Text()
				removeStock(symbol)
			}
		case "4":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}
