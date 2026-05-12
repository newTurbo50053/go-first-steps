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
		return StockQuote{}, fmt.Errorf("4")
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

func api() {

	ibmQuote, err := fetchGlobalQuote("IBM")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(ibmQuote, err)
	}
}
