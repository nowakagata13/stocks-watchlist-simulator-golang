# 📈 Stock Watchlist CLI (Go)

A simple command-line tool to manage and monitor your favorite stock symbols, built in **Golang**.

- 🛠 Add, remove, and list stocks.
- 🔄 Automatic background refreshing every minute.
- 📆 Watchlist persists between sessions (saved to `watchlist.json`).
- 🌎 Stock data fetched live using **Finnhub API**.
- 📚 Good project for practicing Go file operations, APIs, concurrency (goroutines), and modular code design.

---

## 🚀 Features

- Add a stock symbol (e.g., `AAPL`, `TSLA`) to your personal watchlist
- Remove a stock symbol
- Display current prices and daily % change in a table
- Automatically refresh stock data every 60 seconds
- Store your watchlist locally (`watchlist.json`)
- Clean modular code split across multiple files

---

## 💻 Installation

1. **Clone the repo:**

```bash
git clone https://github.com/your-username/stock-watchlist-go.git
cd stock-watchlist-go
```

2. **Install dependencies:**

```bash
go mod tidy
```

3. **Set up API key:**

You need a free API key from [Finnhub.io](https://finnhub.io/).

Create a `.env` file in your project root:

```bash
FINNHUB_API_KEY=your_api_key_here
```

(Or set the environment variable manually.)

---

## 📦 How to Run

```bash
go run main.go
```

Follow the simple text menu:

```
Choose an option:
1. List Watchlist
2. Add Stock
3. Remove Stock
4. Quit
```

✅ Stock prices are refreshed automatically in the background every 60 seconds.

---

## 📁 Project Structure

```
.
├── .env                # API Key storage
├── go.mod              # Go modules
├── main.go             # Program entry point (menu & user input)
├── stock.go            # Fetch stock info from API
├── watchlist.go        # Manage watchlist (load, save, add, remove)
├── refresh.go          # Background refresher
├── utils.go            # (Optional helpers if needed)
├── watchlist.json      # Your saved watchlist
```

---

## 🛠 Built With

- [Go](https://golang.org/) — Programming language
- [Finnhub API](https://finnhub.io/) — Stock price API
- [tablewriter](https://github.com/olekukonko/tablewriter) — ASCII table formatting
- [godotenv](https://github.com/joho/godotenv) — .env file loader for Go

---

## 📜 Future Improvements (Ideas)

- Add alerts (e.g., notify if stock moves +5% or -5%)
- Support cryptocurrencies (e.g., Bitcoin)
- Historical price charts (ASCII art!)
- Command-line arguments support
- Export watchlist to CSV
- Dockerize for easy running anywhere

---

## 📄 License

This project is licensed under the MIT License — feel free to use and modify!

---

# 🙌 Contributing

Pull requests are welcome.  
For major changes, please open an issue first to discuss what you would like to change.

---

# 📬 Contact

**Your Name**  
[Your GitHub Profile](https://github.com/your-username)

---

# ✨ Preview

Here’s how the CLI looks:

```
+--------+---------------+-----------+
| SYMBOL | CURRENT PRICE  | CHANGE %  |
+--------+---------------+-----------+
| AAPL   | 190.34         | +0.71%    |
| TSLA   | 210.21         | -1.45%    |
+--------+---------------+-----------+
```

