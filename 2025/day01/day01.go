package day01

import (
	"fmt"
	"strconv"

	"github.com/jjgry/advent-of-code/2025/utils"
)

// Receive the input as a slice of strings, one per line
func Run(input []string) {
	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}

func part1(input []string) int {
	position := 50
	countOfZeroes := 0

	for _, inputLine := range input {
		direction := inputLine[0:1]
		stepsString := inputLine[1:]
		steps, err := strconv.Atoi(stepsString)
		utils.Check(err)

		if direction == "R" {
			position += steps
		}

		if direction == "L" {
			position -= steps
		}

		for position < 0 {
			position += 100
		}

		for position > 99 {
			position -= 100
		}

		if position == 0 {
			countOfZeroes++
		}

	}

	return countOfZeroes
}

func part2(input []string) int {
	position := 50
	countOfZeroes := 0

	for _, inputLine := range input {
		direction := inputLine[0:1]
		stepsString := inputLine[1:]
		steps, err := strconv.Atoi(stepsString)
		utils.Check(err)

		if direction == "R" {
			for steps > 0 {
				position += 1
				steps--
				for position < 0 {
					position += 100
				}

				for position > 99 {
					position -= 100
				}

				if position == 0 {
					countOfZeroes++
				}
			}
		}

		if direction == "L" {
			for steps > 0 {
				position -= 1
				steps--
				for position < 0 {
					position += 100
				}

				for position > 99 {
					position -= 100
				}

				if position == 0 {
					countOfZeroes++
				}
			}
		}

	}

	return countOfZeroes
}
