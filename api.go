package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type GlobalQuoteResponse struct {
	GlobalQuote Quote `json:"Global Quote"`
}

type Quote struct {
	Symbol string `json:"01. symbol"`
	Price  string `json:"05. price"`
	Volume string `json:"06. volume"`
}

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

func api() {

	var response GlobalQuoteResponse

	apiKey := os.Getenv("ALPHA_API_KEY")

	if apiKey == "" {
		fmt.Println("API KEY not found")
		return
	}

	symbol := "IBM"
	url := "https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=" + symbol + "&apikey=" + apiKey
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("PROBLRM: ", err)
		return
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)

	err = json.Unmarshal(bodyBytes, &response)

	if err != nil {
		fmt.Println("PROBLEM:", err)
		return
	}
	fmt.Println(string(bodyBytes))
	fmt.Printf("Symbol: %s\nCena : %v", response.GlobalQuote.Symbol, response.GlobalQuote.Price)

}
