package day04

type Input []string

func (input Input) Part1() (count int) {
	_, count = newGridAndCount(convertToRuneSlice(input))
	return
}

func (input Input) Part2() (count int) {
	grid := convertToRuneSlice(input)
	for {
		newGrid, newCount := newGridAndCount(grid)
		if newCount == 0 {
			break // Reached steady state
		}
		grid = newGrid
		count += newCount
	}
	return
}

func newGridAndCount(grid [][]rune) (newGrid [][]rune, count int) {
	height := len(grid)
	width := len(grid[0])
	newGrid = copy2dSlice(grid)
	for h := range height {
		for w := range width {
			if grid[h][w] != '@' {
				continue // Only interested in locations with rolls of paper
			}
			neighbourCount := 0
			for ho := -1; ho <= 1; ho++ {
				for wo := -1; wo <= 1; wo++ {
					if ho == 0 && wo == 0 {
						continue // This is the target cell, not a neighbour
					}
					if h+ho < 0 || h+ho >= height || w+wo < 0 || w+wo >= width {
						continue // We are out of bounds, so no neighbour here
					}
					if grid[h+ho][w+wo] == '@' {
						neighbourCount++
					}
				}
			}
			if neighbourCount < 4 {
				count++
				newGrid[h][w] = 'x'
			}
		}
	}
	return
}

func copy2dSlice(grid [][]rune) (gridCopy [][]rune) {
	for _, row := range grid {
		rowCopy := make([]rune, len(row))
		copy(rowCopy, row)
		gridCopy = append(gridCopy, rowCopy)
	}
	return
}

func convertToRuneSlice(input Input) (grid [][]rune) {
	for idx := range len(input) {
		grid = append(grid, []rune(input[idx]))
	}
	return
}
