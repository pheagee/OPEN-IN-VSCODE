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
	GPUFury   uint64 `json:"Fury"`
	GPU470    uint64 `json:"470"`
	GPU480    uint64 `json:"480"`
	GPU570    uint64 `json:"570"`
	GPU580    uint64 `json:580"`
	GPUVega56 uint64 `json:"Vega56"`
	GPUVega64 uin