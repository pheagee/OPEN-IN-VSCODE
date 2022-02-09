package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
)

// Json representation of config file
type ConfigFileJson struct {
	GPU            GPUConfig `json:"gpu"`
	CostPerKw      float64   `json:"cost_per_kw"`
	MinerDirectory string    `