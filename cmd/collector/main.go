package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type (
	Ticker struct {
		ProductCode     string  `json:"product_code"`
		State           string  `json:"state"`
		Timestamp       string  `json:"timestamp"`
		TickId          uint64  `json:"tick_id"`
		BestBid         float64 `json:"best_bid"`
		BestAsk         float64 `json:"best_ask"`
		BestBidSize     float64 `json:"best_bid_size"`
		BestAskSize     float64 `json:"best_ask_size"`
		TotalBidDepth   float64 `json:"total_bid_depth"`
		TotalAskDepth   float64 `json:"total_ask_depth"`
		MarketBidSize   float64 `json:"market_bid_size"`
		MarketAskSize   float64 `json:"market_ask_size"`
		LTP             float64 `json:"ltp"`
		Volume          float64 `json:"volume"`
		VolumeByProduct float64 `json:"volume_by_product"`
	}
)

func apiRequest() ([]byte, error) {
	endpoint := "https://api.bitflyer.com/v1/ticker"
	request, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	client := new(http.Client)
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode >= 400 {
		return nil, errors.New(string(body))
	}

	return body, nil
}

func main() {
	content, err := apiRequest()
	if err != nil {
		panic(err)
	}

	var ticker Ticker
	err = json.Unmarshal(content, &ticker)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", ticker)

	fmt.Println(ticker.ProductCode)
	fmt.Println(ticker.State)
	fmt.Println(ticker.Timestamp)
	fmt.Println(ticker.TickId)
	fmt.Println(ticker.BestBid)
	fmt.Println(ticker.LTP)
}
