package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

// func priceChange(stock Stock, newPrice float64, oldPrice float64) {
// 	change := (newPrice - oldPrice) / oldPrice * 100

// 	fmt.Printf("Price change in percentage for: %v is %.2f%% \n", stock.Symbol, change)
// }

// func mostValuableStock(stocks []Stock) (string, float64) {

// 	maxNazwa := stocks[0].Symbol
// 	maxPrice := stocks[0].Price

// 	for _, s := range stocks {
// 		if s.Price > maxPrice {
// 			maxPrice = s.Price
// 			maxNazwa = s.Symbol

// 		}
// 	}
// 	return maxNazwa, maxPrice

// }

// func averagePrice(stocks []Stock) float64 {
// 	var price float64
// 	for _, s := range stocks {
// 		price += (s.Price)
// 	}
// 	return price / float64(len(stocks))
// }

func fetchGlobalQuote(symbol string) (StockQuote, error) {
	var response GlobalQuoteResponse

	apiKey := os.Getenv("ALPHA_API_KEY")

	if apiKey == "" {
		return StockQuote{}, fmt.Errorf("api key not found")
	}

	url := "https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=" + symbol + "&apikey=" + apiKey

	resp, err := http.Get(url)

	if err != nil {
		return StockQuote{}, fmt.Errorf("request error: %v", err)
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		return StockQuote{}, fmt.Errorf("read body error: %v", err)
	}

	err = json.Unmarshal(bodyBytes, &response)

	if err != nil {
		return StockQuote{}, fmt.Errorf("json parse error: %v", err)
	}

	price, err := strconv.ParseFloat(response.GlobalQuote.Price, 64)

	if err != nil {
		return StockQuote{}, fmt.Errorf("price conversion error: %v", err)
	}

	volume, err := strconv.Atoi(response.GlobalQuote.Volume)

	if err != nil {
		return StockQuote{}, fmt.Errorf("volume conversion error: %v", err)
	}

	return StockQuote{
		Symbol: response.GlobalQuote.Symbol,
		Price:  price,
		Volume: volume,
	}, nil

}

func api() {

	quote, err := fetchGlobalQuote("AAP3L")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(quote.Symbol)
	fmt.Println(quote.Price)
	fmt.Println(quote.Volume)
}
