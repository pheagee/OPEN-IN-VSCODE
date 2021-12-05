
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