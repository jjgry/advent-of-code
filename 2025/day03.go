package main

import (
	"fmt"
	"math"
	"strconv"
)

// Receive the input as a slice of strings, one per line
func day03(input []string) {
	fmt.Println("Part 1:", day03part1(input))
	fmt.Println("Part 2:", day03part2(input))
}

func day03part1(banks []string) int {
	return day3fn(banks, 2)
}

func day03part2(banks []string) int {
	return day3fn(banks, 12)
}

func day3fn(banks []string, numberOfBatteries int) int {
	rollingSum := 0
	for _, bankStr := range banks {
		// Convert to ints now to make life easier later
		bankInts := stringToIntSlice(bankStr)

		// Find the batteries we need to use
		selectedBatteries := selectBatteries(numberOfBatteries, bankInts)

		// Sum up our battery joltages to calculate the joltage
		// Why doesn't Go have inbuilt integer exponentiation!
		joltage := 0
		for idx, val := range selectedBatteries {
			joltage += int(math.Pow10(numberOfBatteries-idx-1)) * val
		}

		rollingSum += joltage
	}
	return rollingSum
}

func selectBatteries(numBatteriesToFind int, bank []int) []int {
	selectedIdx := findNextBatteryIdx(numBatteriesToFind, bank)
	selectedBattery := []int{bank[selectedIdx]}

	if numBatteriesToFind == 1 {
		return selectedBattery
	} else {
		return append(selectedBattery, selectBatteries(numBatteriesToFind-1, bank[selectedIdx+1:])...)
	}
}

func findNextBatteryIdx(numBatteriesToFind int, bank []int) int {
	// Only search a subsection of the bank (we need to leave enough batteries for later steps)
	bankToSearch := bank[:len(bank)-numBatteriesToFind+1]
	for target := 9; target >= 1; target-- {
		for idx, joltage := range bankToSearch {
			if joltage == target {
				return idx
			}
		}
	}
	panic("Could not find a next battery")
}

func stringToIntSlice(str string) []int {
	intSlice := make([]int, len(str))
	for idx, char := range str {
		intVal, err := strconv.Atoi(string(char))
		check(err)
		intSlice[idx] = intVal
	}
	return intSlice
}
