package day02

import (
	"strconv"
	"strings"

	"github.com/jjgry/advent-of-code/2025/utils"
)

type Input []string

func (ranges Input) Part1() int {
	var invalidNumbers []int

	for _, arange := range ranges {
		rangeEnds := strings.Split(arange, "-")
		rangeStartStr := rangeEnds[0]
		rangeEndStr := rangeEnds[1]

		rangeStart, err := strconv.Atoi(rangeStartStr)
		utils.Check(err)

		rangeEnd, err := strconv.Atoi(rangeEndStr)
		utils.Check(err)

		for candidate := rangeStart; candidate <= rangeEnd; candidate++ {
			candidateString := strconv.Itoa(candidate)

			if len(candidateString)%2 != 0 {
				// Candidate is an odd number of digits so can't be two repeated values
				continue
			}

			if candidateString[:len(candidateString)/2] == candidateString[len(candidateString)/2:] {
				// First half and second half are the same, so add this to invalid numbers
				invalidNumbers = append(invalidNumbers, candidate)
			}
		}
	}

	sum := 0
	for _, invalidNumber := range invalidNumbers {
		sum += invalidNumber
	}
	return sum
}

func (ranges Input) Part2() int {
	var invalidNumbers []int

	for _, arange := range ranges {
		rangeEnds := strings.Split(arange, "-")
		rangeStartStr := rangeEnds[0]
		rangeEndStr := rangeEnds[1]

		rangeStart, err := strconv.Atoi(rangeStartStr)
		utils.Check(err)

		rangeEnd, err := strconv.Atoi(rangeEndStr)
		utils.Check(err)

	CandidateLoop:
		for candidate := rangeStart; candidate <= rangeEnd; candidate++ {
			candidateString := strconv.Itoa(candidate)

		HowManyPartsLoop:
			for howManyParts := 2; howManyParts <= len(candidateString); howManyParts++ {
				if len(candidateString)%howManyParts != 0 {
					// Candidate cannot be evenly split into 'howManyParts'
					continue
				}

				var parts []string
				partSize := len(candidateString) / howManyParts
				for i := 0; i < howManyParts; i++ {
					parts = append(parts, candidateString[i*partSize:i*partSize+partSize])
				}

				for i := 0; i < len(parts)-1; i++ {
					if parts[i] != parts[i+1] {
						// The component parts are not all repeated so move onto the next
						continue HowManyPartsLoop
					}
				}

				// We've split into 'howManyParts' segments and all of these are equal
				invalidNumbers = append(invalidNumbers, candidate)
				// Don't try anything else with this candidate
				continue CandidateLoop
			}

		}
	}
	sum := 0
	for _, invalidNumber := range invalidNumbers {
		sum += invalidNumber

	}
	return sum
}
