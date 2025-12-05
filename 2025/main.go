package main

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
	"time"

	"github.com/jjgry/advent-of-code/2025/day01"
	"github.com/jjgry/advent-of-code/2025/day02"
	"github.com/jjgry/advent-of-code/2025/day03"
	"github.com/jjgry/advent-of-code/2025/day04"
	"github.com/jjgry/advent-of-code/2025/day05"
	"github.com/jjgry/advent-of-code/2025/day06"
	"github.com/jjgry/advent-of-code/2025/day07"
	"github.com/jjgry/advent-of-code/2025/day08"
	"github.com/jjgry/advent-of-code/2025/day09"
	"github.com/jjgry/advent-of-code/2025/day10"
	"github.com/jjgry/advent-of-code/2025/day11"
	"github.com/jjgry/advent-of-code/2025/day12"
	"github.com/jjgry/advent-of-code/2025/utils"
)

type DayRunner interface {
	Part1() int
	Part2() int
}

func runAdventOfCode2025() {
	var dayToRun string

	// Use the current day as default
	now := time.Now()
	if now.Month() != time.December || now.Day() < 1 || now.Day() > 12 {
		fmt.Println("Today isn't a valid date! Defaulting to December 1st.")
		dayToRun = "1"
	} else {
		dayToRun = strconv.Itoa(now.Day())
	}

	//
	//
	// For convenience, you can hardcode the day here
	// dayToRun = "1"
	//
	//

	// Add leading zero if needed
	dayToRunWithLeadingZero := fmt.Sprintf("%02s", dayToRun)

	inputFileName := fmt.Sprintf("C:/dev/advent-of-code/2025/input/%s.txt", dayToRunWithLeadingZero)

	data, err := os.ReadFile(inputFileName)
	utils.Check(err)
	input := utils.SplitLines(data)

	var runner DayRunner
	switch dayToRunWithLeadingZero {
	case "01":
		runner = day01.Input(input)
	case "02":
		runner = day02.Input(input)
	case "03":
		runner = day03.Input(input)
	case "04":
		runner = day04.Input(input)
	case "05":
		runner = day05.Input(input)
	case "06":
		runner = day06.Input(input)
	case "07":
		runner = day07.Input(input)
	case "08":
		runner = day08.Input(input)
	case "09":
		runner = day09.Input(input)
	case "10":
		runner = day10.Input(input)
	case "11":
		runner = day11.Input(input)
	case "12":
		runner = day12.Input(input)
	default:
		panic("Day not implemented")
	}

	part1Start := time.Now()
	part1Result := runner.Part1()
	part1Duration := time.Since(part1Start).String()

	part2Start := time.Now()
	part2Result := runner.Part2()
	part2Duration := time.Since(part2Start).String()

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintf(w, "Day %v\tResult\tTime taken\n", dayToRunWithLeadingZero)
	fmt.Fprintf(w, "Part 1:\t%v\t%v\n", part1Result, part1Duration)
	fmt.Fprintf(w, "Part 2:\t%v\t%v\n", part2Result, part2Duration)
	w.Flush()
}

func main() {
	runAdventOfCode2025()
}
