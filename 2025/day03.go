package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Receive the input as a slice of strings, one per line
func day03(input []string) {
	fmt.Println("Part 1:", day03part1(input))
	fmt.Println("Part 2:", day03part2(input))
}

func day03part1(banks []string) int {
	rollingSum := 0
	for _, bank := range banks {
		maxValueInBank := 0
		length := len(bank)
		for firstLetterIndex := 0; firstLetterIndex < length-1; firstLetterIndex++ {
			for secondLetterIndex := firstLetterIndex + 1; secondLetterIndex < length; secondLetterIndex++ {
				firstNumber, err := strconv.Atoi(string(bank[firstLetterIndex]))
				check(err)
				secondNumber, err := strconv.Atoi(string(bank[secondLetterIndex]))
				check(err)
				value := 10*firstNumber + secondNumber
				if value > maxValueInBank {
					maxValueInBank = value
				}
			}
		}
		rollingSum += maxValueInBank
	}
	return rollingSum
}

func day03part2Slow(banks []string) int {
	rollingSum := 0
	for _, bank := range banks {
		maxValueInBank := 0
		length := len(bank)
		for firstLetterIndex := 0; firstLetterIndex < length-11; firstLetterIndex++ {
			fmt.Println("Checked nth first letter: ", firstLetterIndex)
			for secondLetterIndex := firstLetterIndex + 1; secondLetterIndex < length-10; secondLetterIndex++ {
				fmt.Println("Checked nth second letter: ", secondLetterIndex)
				for thirdLetterIndex := secondLetterIndex + 1; thirdLetterIndex < length-9; thirdLetterIndex++ {
					fmt.Println("Checked nth third letter: ", thirdLetterIndex)
					for fourthLetterIndex := thirdLetterIndex + 1; fourthLetterIndex < length-8; fourthLetterIndex++ {
						fmt.Println("Checked nth fourth letter: ", fourthLetterIndex)
						for fifthLetterIndex := fourthLetterIndex + 1; fifthLetterIndex < length-7; fifthLetterIndex++ {
							fmt.Println("Checked nth fifth letter: ", fifthLetterIndex)
							for sixthLetterIndex := fifthLetterIndex + 1; sixthLetterIndex < length-6; sixthLetterIndex++ {
								fmt.Println("Checked nth sixth letter: ", sixthLetterIndex)
								for seventhLetterIndex := sixthLetterIndex + 1; seventhLetterIndex < length-5; seventhLetterIndex++ {
									fmt.Println("Checked nth seventh letter: ", seventhLetterIndex)
									for eighthLetterIndex := seventhLetterIndex + 1; eighthLetterIndex < length-4; eighthLetterIndex++ {
										fmt.Println("Checked nth eighth letter: ", eighthLetterIndex)
										for ninthLetterIndex := eighthLetterIndex + 1; ninthLetterIndex < length-3; ninthLetterIndex++ {
											for tenthLetterIndex := ninthLetterIndex + 1; tenthLetterIndex < length-2; tenthLetterIndex++ {
												for eleventhLetterIndex := tenthLetterIndex + 1; eleventhLetterIndex < length-1; eleventhLetterIndex++ {
													for twelvthLetterIndex := eleventhLetterIndex + 1; twelvthLetterIndex < length; twelvthLetterIndex++ {
														var sb strings.Builder
														sb.WriteByte(bank[firstLetterIndex])
														sb.WriteByte(bank[secondLetterIndex])
														sb.WriteByte(bank[thirdLetterIndex])
														sb.WriteByte(bank[fourthLetterIndex])
														sb.WriteByte(bank[fifthLetterIndex])
														sb.WriteByte(bank[sixthLetterIndex])
														sb.WriteByte(bank[seventhLetterIndex])
														sb.WriteByte(bank[eighthLetterIndex])
														sb.WriteByte(bank[ninthLetterIndex])
														sb.WriteByte(bank[tenthLetterIndex])

														sb.WriteByte(bank[eleventhLetterIndex])
														sb.WriteByte(bank[twelvthLetterIndex])

														value, err := strconv.Atoi(sb.String())
														check(err)

														if value > maxValueInBank {
															maxValueInBank = value
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
		fmt.Println("Max value: ", maxValueInBank)
		rollingSum += maxValueInBank
	}
	return rollingSum
}

func day03part2(banks []string) int {
	rollingSum := 0.0
	for _, bank := range banks {
		// Convert the string into an array of ints
		var bankInts []int
		for idx, val := range bank {
			intVal, err := strconv.Atoi(string(val))
			check(err)
			if idx == 0 {
				bankInts = []int{intVal}
			} else {
				bankInts = append(bankInts, intVal)
			}
		}

		numbers := findNextNumber(12, bankInts, 0)

		// Sum up our numbers to calculate the joltage
		// Why doesn't Go have inbuilt integer exponentiation!
		total := 0.0
		for idx, val := range numbers {
			total += math.Pow10(12-idx-1) * float64(val)
		}

		// Add this joltage to the total
		rollingSum += total
	}
	return int(rollingSum)
}

func findNextNumber(numberOfBatteriesToFind int, bankValue []int, indexToFind int) []int {
	// Work out how many batteries we need to leave in the remainer for our recursion to find
	numberThatMustRemain := numberOfBatteriesToFind - 1

	// Find the first largest number
	var selectedIndex int
TargetNumberLoop:
	for targetNumber := 9; targetNumber >= 1; targetNumber-- {
		for idx := 0; idx < len(bankValue)-numberThatMustRemain; idx++ {
			if bankValue[idx] == targetNumber {
				// Found our first number
				selectedIndex = idx
				break TargetNumberLoop
			}
		}
	}

	returnVal := []int{bankValue[selectedIndex]}

	// If we're looking for the last number, return an array with just that value
	if numberThatMustRemain == 0 {
		return returnVal
	}

	// Otherwise find the rest of the array and concatenate it together
	remainder := findNextNumber(numberOfBatteriesToFind-1, bankValue[selectedIndex+1:], indexToFind+1)
	return append(returnVal, remainder...)
}
