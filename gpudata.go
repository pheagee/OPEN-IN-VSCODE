
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
type GPU struct {
	Ethash      Algorithm
	Groestl     Algorithm
	X11Gost     Algorithm
	CryptoNight Algorithm
	Equihash    Algorithm
	Lyra2REv2   Algorithm
	NeoScrypt   Algorithm
	LBRY        Algorithm
	Blake2b     Algorithm
	Blake14r    Algorithm
	Pascal      Algorithm
	Skunkhash   Algorithm
}

// Data scraped from www.whattomine.com  01.01.2018
var GPU280x GPU = GPU{Algorithm{11, 220}, Algorithm{23.8, 250}, Algorithm{2.9, 200}, Algorithm{490, 220}, Algorithm{290, 230}, Algorithm{14050, 220}, Algorithm{490, 250}, Algorithm{60, 200}, Algorithm{960, 250}, Algorithm{1450, 220}, Algorithm{580, 250}, Algorithm{0.0, 0.0}}
var GPU380 GPU = GPU{Algorithm{20.2, 145}, Algorithm{15.5, 130}, Algorithm{2.5, 120}, Algorithm{530, 120}, Algorithm{205, 130}, Algorithm{6400, 125}, Algorithm{350, 145}, Algorithm{44, 135}, Algorithm{760, 150}, Algorithm{1140, 155}, Algorithm{480, 145}, Algorithm{9, 120}}
var GPUFury GPU = GPU{Algorithm{28.2, 180}, Algorithm{17.4, 180}, Algorithm{4.5, 140}, Algorithm{800, 120}, Algorithm{455, 200}, Algorithm{14200, 190}, Algorithm{500, 160}, Algorithm{83, 200}, Algorithm{1400, 260}, Algorithm{1900, 270}, Algorithm{950, 270}, Algorithm{0.0, 0.0}}
var GPU470 GPU = GPU{Algorithm{26, 120}, Algorithm{14.5, 120}, Algorithm{5.3, 125}, Algorithm{660, 100}, Algorithm{260, 110}, Algorithm{4400, 120}, Algorithm{600, 140}, Algorithm{80, 120}, Algorithm{800, 120}, Algorithm{1100, 120}, Algorithm{510, 120}, Algorithm{15.0, 105}}