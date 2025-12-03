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
	for _, bank := range banks {
		// Convert the string into an array of ints
		bankInts := make([]int, len(bank))
		for idx, val := range bank {
			intVal, err := strconv.Atoi(string(val))
			check(err)
			bankInts[idx] = intVal
		}

		selectedBatteries := recusivelySelectBatteries(numberOfBatteries, bankInts)

		// Sum up our numbers to calculate the joltage
		// Why doesn't Go have inbuilt integer exponentiation!
		joltage := 0
		for idx, val := range selectedBatteries {
			joltage += int(math.Pow10(numberOfBatteries-idx-1)) * val
		}

		rollingSum += joltage
	}
	return rollingSum
}

func recusivelySelectBatteries(numberOfBatteriesToFind int, bankValue []int) []int {
	var selectedIndex int
TargetNumberLoop:
	for targetNumber := 9; targetNumber >= 1; targetNumber-- {
		for idx := 0; idx <= len(bankValue)-numberOfBatteriesToFind; idx++ {
			if bankValue[idx] == targetNumber {
				// Found our first number
				selectedIndex = idx
				break TargetNumberLoop
			}
		}
	}

	currentSelection := []int{bankValue[selectedIndex]}

	// If we're looking for the last number, return an array with just that value
	if numberOfBatteriesToFind == 1 {
		return currentSelection
	}

	// Otherwise find the rest of the array and concatenate it together
	recursiveSelection := recusivelySelectBatteries(numberOfBatteriesToFind-1, bankValue[selectedIndex+1:])
	return append(currentSelection, recursiveSelection...)
}
