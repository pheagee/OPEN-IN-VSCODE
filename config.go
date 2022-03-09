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
	MinerDirectory string    `json:"miner_directory"`
}

// Each GPU type possible
type GPUConfig struct {
	GPU280x   uint64 `json:"280x"`
	GPU380    uint64 `json:"380"`
	GPU