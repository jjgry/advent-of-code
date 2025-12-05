package day05

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jjgry/advent-of-code/2025/utils"
)

// Receive the input as a slice of strings, one per line
func Run(input []string) {
	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}

func part1(input []string) (count int) {
	ids, availables := splitIdsAndAvailable(input)
	for _, available := range availables {
		for _, id := range ids {
			available, err := strconv.Atoi(available)
			utils.Check(err)
			if available >= id.start && available <= id.end {
				count++
				break
			}
		}
	}
	return
}

func part2(input []string) (count int) {
	ranges, _ := splitIdsAndAvailable(input)
	// Loop until we no longer find any merges to make changes
	for found := true; found; {
		found = false
		// Loop over all combinations where i < j
		for i := 0; i < len(ranges)-1; i++ {
			for j := i + 1; j < len(ranges); j++ {
				if ranges[j].start <= ranges[i].start && ranges[j].end >= ranges[i].end {
					// Remove I
					ranges = append(ranges[:i], ranges[i+1:]...)
					found = true
				} else if ranges[i].start <= ranges[j].start && ranges[i].end >= ranges[j].end {
					// Remove J
					ranges = append(ranges[:j], ranges[j+1:]...)
					found = true
				} else if ranges[i].end >= ranges[j].start && ranges[i].start <= ranges[j].start && ranges[i].end <= ranges[j].end {
					// Create a new element and remove I and J
					ranges = append(ranges, Range{ranges[i].start, ranges[j].end})
					ranges = append(ranges[:i], ranges[i+1:]...)
					ranges = append(ranges[:j-1], ranges[j:]...)
					found = true
				} else if ranges[j].end >= ranges[i].start && ranges[j].start <= ranges[i].start && ranges[j].end <= ranges[i].end {
					// Create a new element and remove I and J
					ranges = append(ranges, Range{ranges[j].start, ranges[i].end})
					ranges = append(ranges[:i], ranges[i+1:]...)
					ranges = append(ranges[:j-1], ranges[j:]...)
					found = true
				}
			}
		}
	}
	for _, arange := range ranges {
		count += arange.end - arange.start + 1
	}
	return count
}

type Range struct {
	start int
	end   int
}

func splitIdsAndAvailable(input []string) (ranges []Range, availableLines []string) {
	for i, line := range input {
		if line == "" {
			return extractRanges(input[:i]), input[i+1:]
		}
	}
	return
}

func extractRanges(idLines []string) (ranges []Range) {
	for _, id := range idLines {
		vals := strings.Split(id, "-")
		startRange, err := strconv.Atoi(vals[0])
		utils.Check(err)
		endRange, err := strconv.Atoi(vals[1])
		utils.Check(err)
		ranges = append(ranges, Range{startRange, endRange})
	}
	return
}
