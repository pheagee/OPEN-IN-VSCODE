
package main

import (
	"bytes"
	"testing"
)

var writeOneParameterQueryTestInputOutput = []struct {
	adaptString    string
	numOfGPUs      string
	expectedOutput string
}{
	{"adapt_q_280x=", "0", "adapt_q_280x=0&"},
	{"adapt_q_280x=", "1", "adapt_q_280x=1&adapt_280x=true&"},
}
