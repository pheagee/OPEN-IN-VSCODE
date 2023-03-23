
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
	writeOneParameterQuery(&buffer, "adapt_q_480=", strconv.FormatUint(config.GPU.GPU480, 10))
	writeOneParameterQuery(&buffer, "adapt_q_570=", strconv.FormatUint(config.GPU.GPU570, 10))
	writeOneParameterQuery(&buffer, "adapt_q_580=", strconv.FormatUint(config.GPU.GPU580, 10))
	writeOneParameterQuery(&buffer, "adapt_q_vega56=", strconv.FormatUint(config.GPU.GPUVega56, 10))
	writeOneParameterQuery(&buffer, "adapt_q_vega64=", strconv.FormatUint(config.GPU.GPUVega64, 10))
	writeOneParameterQuery(&buffer, "adapt_q_750Ti=", strconv.FormatUint(config.GPU.GPU750Ti, 10))
	writeOneParameterQuery(&buffer, "adapt_q_1050Ti=", strconv.FormatUint(config.GPU.GPU1050Ti, 10))
	// this "10606" seems like a typo but that's the way the parameter is written
	writeOneParameterQuery(&buffer, "adapt_q_10606=", strconv.FormatUint(config.GPU.GPU1060, 10))
	writeOneParameterQuery(&buffer, "adapt_q_1070=", strconv.FormatUint(config.GPU.GPU1070, 10))
	writeOneParameterQuery(&buffer, "adapt_q_1080=", strconv.FormatUint(config.GPU.GPU1080, 10))
	writeOneParameterQuery(&buffer, "adapt_q_1080Ti=", strconv.FormatUint(config.GPU.GPU1080Ti, 10))

	// Ethash //

	// "eth=true&"
	buffer.WriteString("eth=true&")
	// factor%5Beth_hr%5D=120.0&
	buffer.WriteString("factor%5Beth_hr%5D=" + strconv.FormatFloat(totalGPUsCharacteristics.Ethash.HashRate, 'f', -1, 64) + "&")
	// factor%5Beth_p%5D=590.0&
	buffer.WriteString("factor%5Beth_p%5D=" + strconv.FormatFloat(totalGPUsCharacteristics.Ethash.Power, 'f', -1, 64) + "&")

	// Groestl //

	// grof=true&
	buffer.WriteString("grof=true&")
	// factor%5Bgro_hr%5D=142.0&
	buffer.WriteString("factor%5Bgro_hr%5D=" + strconv.FormatFloat(totalGPUsCharacteristics.Groestl.HashRate, 'f', -1, 64) + "&")
	// factor%5Bgro_p%5D=590.0&
	buffer.WriteString("factor%5Bgro_p%5D=" + strconv.FormatFloat(totalGPUsCharacteristics.Groestl.Power, 'f', -1, 64) + "&")

	// X11Gost //

	// x11gf=true&
	buffer.WriteString("x11gf=true&")
	// factor%5Bx11g_hr%5D=46.0&
	buffer.WriteString("factor%5Bx11g_hr%5D=" + strconv.FormatFloat(totalGPUsCharacteristics.X11Gost.HashRate, 'f', -1, 64) + "&")
	// factor%5Bx11g_p%5D=590.0&
	buffer.WriteString("factor%5Bx11g_p%5D=" + strconv.FormatFloat(totalGPUsCharacteristics.X11Gost.Power, 'f', -1, 64) + "&")

	// CryptoNight //
