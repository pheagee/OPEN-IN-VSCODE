
package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"time"
)

const (
	bitcoinUrl string = "https://api.coinmarketcap.com/v1/ticker/bitcoin/"
)

// Struct type that represents one coin from www.whattomine.com
type Coin struct {
	Id                  uint64  `json:"id"`
	Tag                 string  `json:"tag"`
	Algorithm           string  `json:"algorithm"`
	Block_time          float64 `json:"block_time"`
	Block_reward        float64 `json:"block_reward"`
	Block_reward24      float64 `json:"block_reward24"`
	Last_block          uint64  `json:"last_block"`
	Difficulty          float64 `json:"difficulty"`
	Difficulty24        float64 `json:"difficulty24"`
	Nethash             float64 `json:"nethash"`
	Exchange_rate       float64 `json:"exchange_rate"`
	Exchange_rate24     float64 `json:"exchange_rate24"`
	Exchange_rage_vol   float64 `json:"exchange_rage_vol"`
	Exchange_rage_curr  string  `json:"exchange_rage_curr"`
	Market_cap          string  `json:"market_cap"`
	Estimated_rewards   string  `json:"estimated_rewards"`
	Estimated_rewards24 string  `json:"estimated_rewards24"`
	Btc_revenue         string  `json:"btc_revenue"`
	Btc_revenue24       string  `json:"btc_revenue24"`
	Profitability       uint64  `json:"profitability"`
	Profitability24     uint64  `json:"profitability24"`
	Lagging             bool    `json:"lagging"`
	Timestamp           uint64  `json:"timestamp"`
}

// Struct type that represents one coin from www.coinmarketcap.com/api
type CoinMarketCapCoin struct {
	Id                 string `json:"id"`                 // "bitcoin"
	Name               string `json:"name"`               // "Bitcoin"
	Symbol             string `json:"symbol"`             //: "BTC",
	Rank               uint64 `json:"rank"`               // "1",
	Price_USD          string `json:"price_usd"`          //: "13887.8",
	Print_BTC          string `json:"price_btc"`          //: "1.0",
	Day_Volume_USD     string `json:"24h_volume_usd"`     //: "10438500000.0",
	Market_Cap_USD     string `json:"market_cap_usd"`     //: "232799274884",
	Available_Supply   string `json:"available_supply"`   //: "16762862.0",
	Total_Supply       string `json:"total_supply"`       //: "16762862.0",
	Max_Supply         string `json:"max_supply"`         //: "21000000.0",
	Percent_Change_1h  string `json:"percent_change_1h"`  //: "-0.55",