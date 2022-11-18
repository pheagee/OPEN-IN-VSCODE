
package main

import (
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	coinsRegexp = `([G|g]obyte|[E|e]thereum|[T|t]rezarcoin|[Z|z]cash|[Z|z]classic|[Z|z]encash)`
)

func init() {
	// array that contains names of gpus
	var GPU_Names [numOfGPUs]string
	// initialize array that contains names of gpus
	GPU_Names[0] = "GPU280x"
	GPU_Names[1] = "GPU380"
	GPU_Names[2] = "GPUFury"
	GPU_Names[3] = "GPU470"
	GPU_Names[4] = "GPU480"
	GPU_Names[5] = "GPU570"
	GPU_Names[6] = "GPU580"
	GPU_Names[7] = "GPUVega56"
	GPU_Names[8] = "GPUVega64"
	GPU_Names[9] = "GPU750Ti"
	GPU_Names[10] = "GPU1050Ti"
	GPU_Names[11] = "GPU1060"
	GPU_Names[12] = "GPU1070"
	GPU_Names[13] = "GPU1080"
	GPU_Names[14] = "GPU1080Ti"

	// initilize array that contains GPU characteristics in an order corresponding the name written in the GPU_Names array
	var GPU_HashRates [numOfGPUs]GPU
	GPU_HashRates[0] = GPU280x
	GPU_HashRates[1] = GPU380
	GPU_HashRates[2] = GPUFury
	GPU_HashRates[3] = GPU470
	GPU_HashRates[4] = GPU480
	GPU_HashRates[5] = GPU570
	GPU_HashRates[6] = GPU580
	GPU_HashRates[7] = GPUVega56
	GPU_HashRates[8] = GPUVega64
	GPU_HashRates[9] = GPU750Ti
	GPU_HashRates[10] = GPU1050Ti
	GPU_HashRates[11] = GPU1060
	GPU_HashRates[12] = GPU1070
	GPU_HashRates[13] = GPU1080
	GPU_HashRates[14] = GPU1080Ti

	// initialize map that contains GPUs characteristics scraped from whattomine
	GPUs = make(map[string]GPU)
	for i := range GPU_Names {
		GPUs[GPU_Names[i]] = GPU_HashRates[i]
	}

}