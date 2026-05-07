package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type Stock struct {
	Symbol string  `json:"Symbol"`
	Price  float64 `json:"Price"`
	Volume int     `json:"Volume"`
}

func priceChange(stock Stock, newPrice float64, oldPrice float64) {
	change := (newPrice - oldPrice) / oldPrice * 100

	fmt.Printf("Price change in percentage for: %v is %.2f%% \n", stock.Symbol, change)
}

func mostValuableStock(stocks []Stock) (string, float64) {

	maxNazwa := stocks[0].Symbol
	maxPrice := stocks[0].Price

	for _, s := range stocks {
		if s.Price > maxPrice {
			maxPrice = s.Price
			maxNazwa = s.Symbol

		}
	}
	return maxNazwa, maxPrice

}

func averagePrice(stocks []Stock) float64 {
	var price float64
	for _, s := range stocks {
		price += (s.Price)
	}
	var avgPrice float64 = price / float64(len(stocks))
	return avgPrice
}

func api() {
	apiKey := os.Getenv("ALPHA_API_KEY")

	if apiKey == "" {
		fmt.Println("API KEY not found bitch ass ")
		return
	}

	symbol := "IBM"

	resp, err := http.Get("https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=" + symbol + "&apikey=" + apiKey)

	if err != nil {
		fmt.Println("PROBLRM")
	}

	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	fmt.Println(string(bodyBytes))

}
