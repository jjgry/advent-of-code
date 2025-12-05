package day05

import (
	"strconv"
	"strings"

	"github.com/jjgry/advent-of-code/2025/utils"
)

type Input []string

func (input Input) Part1() (count int) {
	ids, availables := splitIdsAndAvailable(input)
	for _, available := range availables {
		for _, id := range ids {
			if available >= id.start && available <= id.end {
				count++
				break
			}
		}
	}
	return
}

func (input Input) Part2() (count int) {
	ranges, _ := splitIdsAndAvailable(input)
	for foundMerge := true; foundMerge; {
		foundMerge = false
		// Loop over all combinations of i and j and merge wherever possible
		for i := 0; i < len(ranges); i++ {
			for j := 0; j < len(ranges); j++ {
				if i == j {
					continue
				}
				if ranges[j].start <= ranges[i].start && ranges[j].end >= ranges[i].end {
					// J emcompasses I, remove I
					ranges = append(ranges[:i], ranges[i+1:]...)
					foundMerge = true
				} else if ranges[i].end >= ranges[j].start && ranges[i].start <= ranges[j].start && ranges[i].end <= ranges[j].end {
					// End of I overlaps with start of J, merge the two
					ranges = append(ranges, Range{ranges[i].start, ranges[j].end})
					if i < j {
						ranges = append(ranges[:i], ranges[i+1:]...)
						ranges = append(ranges[:j-1], ranges[j:]...)
					} else {
						ranges = append(ranges[:j], ranges[j+1:]...)
						ranges = append(ranges[:i-1], ranges[i:]...)
					}
					foundMerge = true
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

func splitIdsAndAvailable(input []string) (ranges []Range, available []int) {
	var splitIdx int
	for idx, line := range input {
		if line == "" {
			splitIdx = idx
			break
		}
	}
	// Extract ranges
	for _, id := range input[:splitIdx] {
		vals := strings.Split(id, "-")
		startRange, err := strconv.Atoi(vals[0])
		utils.Check(err)
		endRange, err := strconv.Atoi(vals[1])
		utils.Check(err)
		ranges = append(ranges, Range{startRange, endRange})
	}
	// Extract available
	for _, strVal := range input[splitIdx+1:] {
		intVal, err := strconv.Atoi(strVal)
		utils.Check(err)
		available = append(available, intVal)
	}
	return

}
