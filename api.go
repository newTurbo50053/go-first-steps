package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func fetchGlobalQuote(symbol string) (StockQuote, error) {
	var response GlobalQuoteResponse

	apiKey := os.Getenv("ALPHA_API_KEY")

	if apiKey == "" {
		return StockQuote{}, fmt.Errorf("API KEY not foud")
	}

	url := "https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=" + symbol + "&apikey=" + apiKey
	resp, err := http.Get(url)

	if err != nil {
		return StockQuote{}, fmt.Errorf("1")

	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		return StockQuote{}, fmt.Errorf("2")
	}

	err = json.Unmarshal(bodyBytes, &response)

	if err != nil {
		return StockQuote{}, fmt.Errorf("3")
	}

	price, err := strconv.ParseFloat(response.GlobalQuote.Price, 64)

	if err != nil {
		return StockQuote{}, fmt.Errorf(err.Error())
	}
	volume, err := strconv.Atoi(response.GlobalQuote.Volume)

	if err != nil {
		return StockQuote{}, fmt.Errorf("5")
	}

	return StockQuote{
		Symbol: response.GlobalQuote.Symbol,
		Price:  price,
		Volume: volume,
	}, nil

}

func fetchMultipleQuotes(symbols []string) ([]StockQuote, error) {
	quotes := []StockQuote{}

	for _, symb := range symbols {

		quote, err := fetchGlobalQuote(symb)
		if err != nil {
			return nil, fmt.Errorf(err.Error())
		}
		quotes = append(quotes, quote)

	}
	return quotes, nil
}
func api() {

	// ibmQuote, err := fetchGlobalQuote("IBM")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(ibmQuote)

	symbols := []string{"IBM", "AAPL", "MSFT"}

	quotes, err := fetchMultipleQuotes(symbols)
	if err != nil {
		fmt.Println(err)
		return
	}
	// for _,symb := range quotes {
	// 	fmt.Println(Symbol : response.GlobalQuoteResponse.)
	// }
	fmt.Println(quotes)

}
