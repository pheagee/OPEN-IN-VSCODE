
package main

// Data scraped from www.whattomine.com  01.01.2018
const numOfGPUs = 15

// Map of names of GPUs to each GPU characteristic (HashRate and Power Consumption) per Hashing Algorithm
var GPUs map[string]GPU

// Any algorithm has a specific hash rate and a specific power consumption
type Algorithm struct {
	HashRate float64
	Power    float64
}

// Each GPU is a collection of distinct algorithms characteristics