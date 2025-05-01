package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// DB instance
var db *sql.DB

// Initialize the database
func initDB() {
	var err error
	// Open SQLite database (creates the file if it doesn't exist)
	db, err = sql.Open("sqlite3", "./stocks.db")
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}

	// Create tables if they don't exist
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS stocks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		symbol TEXT UNIQUE
	);
	CREATE TABLE IF NOT EXISTS transactions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		type TEXT,
		ticker TEXT,
		quantity INTEGER,
		price REAL
	);
	`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal("Error creating tables: ", err)
	}

	fmt.Println("Database initialized.")
}

// Add stock to the watchlist in the database
func addStockToDB(symbol string) {
	_, err := db.Exec("INSERT OR IGNORE INTO stocks (symbol) VALUES (?)", symbol)
	if err != nil {
		log.Fatalf("Error adding stock to database: %v", err)
	}
	fmt.Printf("Stock %s added to the watchlist.\n", symbol)
}

// Remove stock from the watchlist in the database
func removeStockFromDB(symbol string) {
	_, err := db.Exec("DELETE FROM stocks WHERE symbol = ?", symbol)
	if err != nil {
		log.Fatalf("Error removing stock from database: %v", err)
	}
	fmt.Printf("Stock %s removed from the watchlist.\n", symbol)
}

// Get all stocks from the watchlist
func getStocksFromDB() []string {
	rows, err := db.Query("SELECT symbol FROM stocks")
	if err != nil {
		log.Fatalf("Error fetching stocks from database: %v", err)
	}
	defer rows.Close()

	var stocks []string
	for rows.Next() {
		var symbol string
		if err := rows.Scan(&symbol); err != nil {
			log.Fatalf("Error scanning stock: %v", err)
		}
		stocks = append(stocks, symbol)
	}

	return stocks
}

// Add a transaction to the database
func addTransactionToDB(tType, ticker string, qty int, price float64) {
	_, err := db.Exec("INSERT INTO transactions (type, ticker, quantity, price) VALUES (?, ?, ?, ?)", tType, ticker, qty, price)
	if err != nil {
		log.Fatalf("Error adding transaction to database: %v", err)
	}
	fmt.Printf("%s of %d %s at $%.2f recorded.\n", tType, qty, ticker, price)
}
