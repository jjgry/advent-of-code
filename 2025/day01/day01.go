package day01

import (
	"strconv"

	"github.com/jjgry/advent-of-code/2025/utils"
)

type Input []string

func (input Input) Part1() int {
	return calculateCountOfZeroes(input, false)
}

func (input Input) Part2() int {
	return calculateCountOfZeroes(input, true)
}

func calculateCountOfZeroes(input []string, countAllZeroes bool) int {
	position := 50
	countOfZeroes := 0

	for _, inputLine := range input {
		direction := inputLine[0:1]
		stepsString := inputLine[1:]
		steps, err := strconv.Atoi(stepsString)
		utils.Check(err)

		stepMultiplier := 1
		if direction == "L" {
			stepMultiplier = -1
		}

		step := steps
		if countAllZeroes {
			step = 1 // If counting all zeroes, move one step at a time
		}

		for steps > 0 {
			// Move the position by 'step'
			position = position + step*stepMultiplier

			// Don't overflow 0 or 99
			for position < 0 {
				position += 100
			}
			for position > 99 {
				position -= 100
			}

			steps -= step
			if position == 0 {
				countOfZeroes++
			}
		}
	}
	return countOfZeroes
}
