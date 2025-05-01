# Stocks Watchlist and Simulator (Go)

A simple command-line tool to montior stock prices and simulate buying and selling them, built in **Golang**.

The project uses free API endpoint from Finnhub that allows fetching live prices of stocks.

-  Add, remove, and list stocks.
-  Buy, sell and monitor your portfolio
-  Automatic background refreshing every minute.
-  Stock data fetched live using **Finnhub API**.
-  Good project for practicing Go file operations, APIs, concurrency (goroutines), and modular code design.

---

## Installation

3. **Set up API key:**

To be able to access live prices of stock, you need a free API key from [Finnhub.io](https://finnhub.io/).

Create a `.env` file in your project root and add the key there as FINNHUB_API_KEY.


2. **Install dependencies:**

Clone the repo locally, open the directory and run

```bash
go mod tidy
```

##  **Run**

```bash
go build .
go run main.go
```

Follow the simple text menu:

```
Choose an option:
1. List Watchlist
2. Add Stock
3. Remove Stock
4. Buy Stock
5. Sell Stock
6. View Portfolio
7. View Transactions
8. Quit
> 
```

Stock prices are refreshed automatically in the background every 60 seconds.

---

## Project Structure

```
.
├── .github/workflows
    ├── go.yml          # yml file to run build workflow in Github Actions
├── .env                # API Key storage
├── go.mod              # Go modules
├── go.sum
├── main.go             # Program entry point (menu & user input)
├── portfolio.go        # Buy, sell and view your portfolio of stocks
├── stock.go            # Fetch stock info from API
├── watchlist.go        # Manage watchlist (load, save, add, remove)
├── watchlist.json      # Your saved watchlist
```

---

## Tech 

- Golang— Programming language
- SQLite - simple SQL database
- Github Actions

---

# Preview

Here’s how the CLI looks:

```
+--------+---------------+-----------+
| SYMBOL | CURRENT PRICE  | CHANGE %  |
+--------+---------------+-----------+
| AAPL   | 190.34         | +0.71%    |
| TSLA   | 210.21         | -1.45%    |
+--------+---------------+-----------+
```

