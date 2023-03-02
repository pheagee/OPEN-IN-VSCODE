
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
		adapt_true := parts[0] + "_" + parts[2] + trueString
		buffer.WriteString(adapt_true)
	}
}

// Build the url with the query parameters for use in www.whattomine.com
func constructUrlQuery(config ConfigFileJson, totalGPUsCharacteristics GPU) string {
	var buffer bytes.Buffer
	buffer.WriteString(urlBase)
	writeOneParameterQuery(&buffer, "adapt_q_280x=", strconv.FormatUint(config.GPU.GPU280x, 10))
	writeOneParameterQuery(&buffer, "adapt_q_380=", strconv.FormatUint(config.GPU.GPU380, 10))
	writeOneParameterQuery(&buffer, "adapt_q_fury=", strconv.FormatUint(config.GPU.GPUFury, 10))
	writeOneParameterQuery(&buffer, "adapt_q_470=", strconv.FormatUint(config.GPU.GPU470, 10))