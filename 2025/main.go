package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/jjgry/advent-of-code/2025/day01"
	"github.com/jjgry/advent-of-code/2025/day02"
	"github.com/jjgry/advent-of-code/2025/day03"
	"github.com/jjgry/advent-of-code/2025/day04"
	"github.com/jjgry/advent-of-code/2025/utils"
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
	// dayToRun = "1"
	//
	//

	// Add leading zero if needed
	dayToRunWithLeadingZero := fmt.Sprintf("%02s", dayToRun)
	fmt.Println("Running day:", dayToRunWithLeadingZero)

	inputFileName := fmt.Sprintf("C:/dev/advent-of-code/2025/input/%s.txt", dayToRunWithLeadingZero)
	fmt.Println("Using input file:", inputFileName)

	data, err := os.ReadFile(inputFileName)
	utils.Check(err)

	input := utils.SplitLines(data)

	timeStart := time.Now()
	switch dayToRunWithLeadingZero {
	case "01":
		day01.Run(input)
	case "02":
		day02.Run(input)
	case "03":
		day03.Run(input)
	case "04":
		day04.Run(input)
	default:
		fmt.Println("Day not implemented")
	}
	fmt.Println("Time taken:", time.Since(timeStart).String())
}
