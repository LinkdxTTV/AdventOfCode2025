package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type freshRange struct {
	lower int
	upper int
}

func main() {
	input, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	allRanges := []freshRange{}
	parsingRanges := true

	totalFresh := 0

	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			parsingRanges = false
			continue
		}

		if parsingRanges {
			afterSplit := strings.Split(line, "-")
			lower, err := strconv.Atoi(afterSplit[0])
			if err != nil {
				panic(err)
			}
			upper, err := strconv.Atoi(afterSplit[1])
			if err != nil {
				panic(err)
			}

			allRanges = append(allRanges, freshRange{lower: lower, upper: upper})
			continue
		}

		// Parsing numbers then
		asNum, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		for _, freshRange := range allRanges {
			if freshRange.includes(asNum) {
				totalFresh++
				break
			}
		}
	}

	fmt.Println(totalFresh)

	// Part 2: Find total number of numbers in all arrays.. TBH 189 arrays is not that many arrays..
	// The algorithm for this problem is to sort and then check for overlaps and merge.

	sort.SliceStable(allRanges, func(i, j int) bool {
		if allRanges[i].lower == allRanges[j].lower {
			return allRanges[i].upper < allRanges[j].upper
		}
		return allRanges[i].lower < allRanges[j].lower
	})

	mergedRanges := []freshRange{allRanges[0]}

	for i, testRange := range allRanges {
		if i == 0 {
			continue // We already handled the first one
		}

		// Compare to last range
		compareRange := mergedRanges[len(mergedRanges)-1]

		// No overlap
		if testRange.lower > compareRange.upper {
			mergedRanges = append(mergedRanges, testRange)
			continue
		}

		// Update existing
		if testRange.upper > compareRange.upper {
			mergedRanges[len(mergedRanges)-1].upper = testRange.upper
		}
	}

	total := 0
	for _, r := range mergedRanges {
		total += r.upper - r.lower + 1
	}

	fmt.Println(total)
}

func (f freshRange) includes(num int) bool {
	return num >= f.lower && num <= f.upper
}
