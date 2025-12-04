package day04

import (
	"fmt"
	"strings"
)

// Receive the input as a slice of strings, one per line
func Run(input []string) {
	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}

func part1(input []string) int {
	_, count := helper(input)
	return count
}

func part2(input []string) int {
	currentCount := 0
	currentGrid := input
	for {
		newGrid, count := helper(currentGrid)
		if count == 0 {
			break // No more changes to find
		}
		currentGrid = newGrid
		currentCount += count
		for _, x := range newGrid {
			fmt.Println(x)
		}
		fmt.Println()
	}
	return currentCount
}

func helper(input []string) ([]string, int) {
	width := len(input[0])
	height := len(input)

	count := 0

	// Initialise a blank grid of dots
	var newGrid [][]byte
	for range height {
		newGrid = append(newGrid, []byte(strings.Repeat(".", width)))
	}

	for h := range height {
		for w := range width {
			// Only interested in locations with rolls of paper
			if input[h][w] != '@' {
				continue
			}

			countOfNeighbours := 0

			// Top left
			if h != 0 && w != 0 && input[h-1][w-1] == '@' {
				countOfNeighbours++
			}

			// Top
			if h != 0 && input[h-1][w] == '@' {
				countOfNeighbours++
			}

			// Top right
			if h != 0 && w != width-1 && input[h-1][w+1] == '@' {
				countOfNeighbours++
			}

			// Right
			if w != width-1 && input[h][w+1] == '@' {
				countOfNeighbours++
			}

			// Bottom right
			if h != height-1 && w != width-1 && input[h+1][w+1] == '@' {
				countOfNeighbours++
			}

			// Bottom
			if h != height-1 && input[h+1][w] == '@' {
				countOfNeighbours++
			}

			// Bottom left
			if h != height-1 && w != 0 && input[h+1][w-1] == '@' {
				countOfNeighbours++
			}

			// Left
			if w != 0 && input[h][w-1] == '@' {
				countOfNeighbours++
			}

			if countOfNeighbours < 4 {
				count++
				newGrid[h][w] = 'x'
			} else {
				newGrid[h][w] = '@'
			}
		}
	}

	var stringNewGrid []string
	// Convert [][]byte back to []string
	for _, bytearr := range newGrid {
		stringNewGrid = append(stringNewGrid, string(bytearr))
	}

	return stringNewGrid, count
}
