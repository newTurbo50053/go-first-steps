package main

type GlobalQuoteResponse struct {
	GlobalQuote Quote `json:"Global Quote"`
}

type Quote struct {
	Symbol string `json:"01. symbol"`
	Price  string `json:"05. price"`
	Volume string `json:"06. volume"`
}

type StockQuote struct {
	Symbol string
	Price  float64
	Volume int
}
