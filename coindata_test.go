
package main

import (
	"testing"
)

var testInputOutput = []struct {
	input          map[string]float64
	expectedOutput PairList
}{
	{
		map[string]float64{
			"Bitcoin":  15000,
			"Ripple":   1,
			"Ethereum": 1000,
		}, PairList{
			Pair{"Bitcoin", 15000},