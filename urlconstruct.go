
package main

import (
	"bytes"
	"strconv"
	"strings"
)

const (
	urlBase    string = "https://whattomine.com/coins.json?utf8=%E2%9C%93&"
	trueString string = "true&"
)

// appends the strings adaptString and numberOfGPUs
// additionally appends
func writeOneParameterQuery(buffer *bytes.Buffer, adaptString, numberOfGPUs string) {
	buffer.WriteString(adaptString)
	buffer.WriteString(numberOfGPUs)
	buffer.WriteString("&")
	// Add adapt_MODEL=true& whenever there's > 0 cards for that model
	if numberOfGPUs != "0" {
		parts := strings.Split(adaptString, "_")