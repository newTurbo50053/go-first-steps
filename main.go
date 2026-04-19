package main

// "encoding/json"
// "fmt"
//"io"
//"net/http"
//"strings"
//"math"

// const r = 'a'

// type Stock struct {
// 	Symbol string  `json:"Symbol"`
// 	Price  float64 `json:"Price"`
// 	Volume int     `json:"Volume"`
// }

// func priceChange(stock Stock, newPrice float64, oldPrice float64) {
// 	change := (newPrice - oldPrice) / oldPrice * 100

// 	fmt.Printf("Price change in percentage for: %v is %.2f%% \n", stock.Symbol, change)
// }

// func mostValuableStock(stocks []Stock) {

// 	maxNazwa := stocks[0].Symbol
// 	maxPrice := stocks[0].Price

// 	for _, s := range stocks {
// 		if s.Price > maxPrice {
// 			maxPrice = s.Price
// 			maxNazwa = s.Symbol

// 		}
// 	}
// 	fmt.Println(maxNazwa, maxPrice)

// }

// func averagePrice(stocks []Stock) {
// 	var averagePrice float64
// 	for _, s := range stocks {
// 		averagePrice += (s.Price)
// 	}
// 	fmt.Println("Average price: ", averagePrice/float64(len(stocks)))
// }
func main() {

	api()

	//fmt.Println(r + 'a')

	// numbers := []int{10, 20, 30}

	// numbers = append(numbers[:1], append([]int{15}, numbers[1:]...)...)
	// fmt.Println(cap(numbers))

	// for i, v := range numbers {
	// 	fmt.Println(i, v)
	// }

	// const country = "Poland"
	// name := "Tomasz"
	// age := 21
	// score := []int{70, 80, 90}
	// score = append(score, 100)
	// fmt.Printf("Name: %v \nAge: %v \nCountry: %v \nScores: %v ", name, age, country, score)

	// slice1 := []int{100, 50, 50, 50, 100}
	// sum := 0
	// for i, k := range slice1 {

	// 	fmt.Println(i, k)
	// 	sum = sum + k

	// }
	// fmt.Println("Suma:", sum)

	// prices := []int{110, 20, 30, 40}
	// prices[2] += 50
	// fmt.Println(prices[0])
	// fmt.Println(prices[2])
	// fmt.Println(prices)

	// for i := 0; i < len(prices); i++ {
	// 	for j := i + 1; j < len(prices); j++ {
	// 		if prices[i] < prices[j] {
	// 			prices[i], prices[j] = prices[j], prices[i]
	// 		}
	// 	}
	// }
	// fmt.Println(prices)
	// url := "https://query1.finance.yahoo.com/v7/finance/quote?symbols=AAPL"

	// client := &http.Client{}

	// req, err := http.NewRequest("GET", url, nil)
	// if err != nil {
	// 	panic(err)
	// }

	// req.Header.Set("User-Agent", "Mozilla/5.0")

	// resp, err := client.Do(req)
	// if err != nil {
	// 	panic(err)
	// }

	// defer resp.Body.Close()

	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(body))
	//srendnia
	//tablica := [...]int{1, 23, 4, 5, 6, 67}

	// var licznik float64 = 0
	// for i := 0; i < len(tablica); i++ {
	// 	licznik += float64(tablica[i])
	// }

	// fmt.Println(licznik / float64(len(tablica)))
	// max := tablica[0]

	// for i := 0; i < len(tablica); i++ {

	// 	if tablica[i] > max {
	// 		max = tablica[i]
	// 	}
	// }
	// fmt.Println(max)

	// tekst := "67 to w tych czasach moggerskich jest ostre 67  no i 69 yo 69 67 po co ci kapusta ci cic ci"
	// slowa := strings.Fields(tekst)

	// licznik := make(map[string]int)
	// for _, slowo := range slowa {
	// 	licznik[slowo]++
	// }
	// fmt.Println(licznik)

	// 	liczby := []int{11, 2, 3, 41, 23, 6, 7, 8, 9, 10, 311, 12, 13, 134, 15, 16, 172, 184, 19, 20}
	// 	//max := liczby[0]
	// 	for i := 0; i < len(liczby); i++ {
	// 		for j := i; j < len(liczby)-i-1; j++ {
	// 			if liczby[j] > liczby[j+1] {

	// 				liczby[j], liczby[j+1] = liczby[j+1], liczby[j]
	// 			}

	// 		}
	// 	}
	// 	fmt.Println(liczby)

	// stock := []Stock{
	// 	{Symbol: "AAPL", Price: 150.0, Volume: 1000},
	// 	{Symbol: "GOOGL", Price: 2800.0, Volume: 500},
	// }
	// priceChange(stock[0], 160.0, 150.0)
	// mostValuableStock(stock)

	// jsonData := []byte(`[
	// 	{"Symbol":"AAPL","Price":175.5,"Volume":1000},
	// 	{"Symbol":"MSFT","Price":320.1,"Volume":500},
	// 	{"Symbol":"GOOGL","Price":2800.0,"Volume":300},
	// 	{"Symbol":"AMZN","Price":140.3,"Volume":800}
	// 	]`)
	// var Stock []Stock
	// err := json.Unmarshal(jsonData, &Stock)

	// if err != nil {
	// 	fmt.Println("Error parsing JSON:", err)
	// 	return
	// }

	// for i := 0; i < len(Stock); i++ {
	// 	fmt.Println(Stock[i].Symbol, Stock[i].Price, Stock[i].Volume)
	// }

	// fmt.Println("Most valuable stock:")
	// mostValuableStock(Stock)
	// fmt.Println("Average")
	// averagePrice(Stock)
}
