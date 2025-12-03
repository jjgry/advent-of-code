package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	var dayToRun string
	if len(os.Args) == 2 {
		// The day to run is passed as the first argument
		dayToRun = os.Args[1]
	} else {
		// Else use the current day as default
		now := time.Now()
		if now.Month() != time.December || now.Day() < 1 || now.Day() > 12 {
			fmt.Println("Today isn't a valid date! Defaulting to December 1st.")
			dayToRun = "1"
		} else {
			dayToRun = strconv.Itoa(now.Day())
		}
	}

	//
	//
	// For testing, you can hardcode the day here
	// dayToRun = "4"
	//
	//
	//

	// Add leading zero if needed
	dayToRunWithLeadingZero := fmt.Sprintf("%02s", dayToRun)
	fmt.Println("Running day:", dayToRunWithLeadingZero)

	inputFileName := fmt.Sprintf("C:/dev/advent-of-code/2025/input/%s.txt", dayToRunWithLeadingZero)
	fmt.Println("Using input file:", inputFileName)

	data, err := os.ReadFile(inputFileName)
	check(err)

	splitInput := splitLines(data)

	switch dayToRunWithLeadingZero {
	case "01":
		day01(splitInput)
	case "02":
		day02(splitInput)
	case "03":
		day03(splitInput)
	case "04":
		day04(splitInput)
	default:
		fmt.Println("Day not implemented")
	}
}
