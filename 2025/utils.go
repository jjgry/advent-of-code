package main

import "strings"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func splitLines(input []byte) []string {
	myString := string(input[:])
	return strings.Split(myString, "\r\n")
}
