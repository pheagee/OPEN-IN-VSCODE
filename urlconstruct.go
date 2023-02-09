
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