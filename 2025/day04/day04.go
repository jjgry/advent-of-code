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
	_, count := getNewGridAndCount(input)
	return count
}

func part2(input []string) int {
	count := 0
	grid := input
	for {
		newGrid, newCount := getNewGridAndCount(grid)
		if newCount == 0 {
			break // No more changes to find
		}
		grid = newGrid
		count += newCount
	}
	return count
}

func getNewGridAndCount(input []string) ([]string, int) {
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
			if input[h][w] != '@' {
				continue // Only interested in locations with rolls of paper
			}

			countOfNeighbours := 0

			if h != 0 && w != 0 && input[h-1][w-1] == '@' { // Top left
				countOfNeighbours++
			}
			if h != 0 && input[h-1][w] == '@' { // Top
				countOfNeighbours++
			}
			if h != 0 && w != width-1 && input[h-1][w+1] == '@' { // Top right
				countOfNeighbours++
			}
			if w != width-1 && input[h][w+1] == '@' { // Right
				countOfNeighbours++
			}
			if h != height-1 && w != width-1 && input[h+1][w+1] == '@' { // Bottom right
				countOfNeighbours++
			}
			if h != height-1 && input[h+1][w] == '@' { // Bottom
				countOfNeighbours++
			}
			if h != height-1 && w != 0 && input[h+1][w-1] == '@' { // Bottom left
				countOfNeighbours++
			}
			if w != 0 && input[h][w-1] == '@' { // Left
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
