package utils

import (
	"strings"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func SplitLines(input []byte) []string {
	myString := string(input[:])
	return strings.Split(myString, "\r\n")
}
