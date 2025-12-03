package utils

import (
	"fmt"
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

func silly() {
	x := 1

	// x := 2 // Not allowed - the compiler can spot that this is silly

	if true {
		x := 2
		fmt.Println(x) // Prints 2
	}

	fmt.Println(x) // Prints 1

	if true {
		x = 3
	}

	fmt.Println(x) // Prints 3
}
