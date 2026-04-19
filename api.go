package main

import (
	//"encoding/json"
	"fmt"
	"os"
	//"io"
	//"net/http"
	//"strings"
	//"math"
)

func api() {
	apiKey := os.Getenv("ALPHA_API_KEY")
	fmt.Println(apiKey)
}
