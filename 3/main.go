package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const NumBatteriesNeededPart1 = 2
const NumBatteriesNeededPart2 = 12

func main() {
	input, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	joltages := []int{}

	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			continue
		}

		intSlice := convertLineToIntSlice(line)
		joltages = append(joltages, getMaxJoltageFromBankPart2(intSlice, NumBatteriesNeededPart2))
	}

	// fmt.Println(joltages)

	sum := 0
	for _, jolt := range joltages {
		sum += jolt
	}

	fmt.Println(sum)
}

func convertLineToIntSlice(in string) []int {
	out := []int{}
	for _, char := range in {
		asInt, err := strconv.Atoi(string(char))
		if err != nil {
			panic(err)
		}
		out = append(out, asInt)
	}
	return out
}

func getMaxJoltageFromBankPart1(intSlice []int) int {
	// For a slice of len n, grab the biggest number from the first n-1 numbers, call that index m. Then grab the biggest number from m+1 -> end
	firstMax := slices.Max(intSlice[:len(intSlice)-1])

	firstIndex := slices.Index(intSlice, firstMax)

	secondMax := slices.Max(intSlice[firstIndex+1:])

	outputNumber := 10*firstMax + secondMax
	return outputNumber
}

// ^
// This function is generically solved by part 2 below

func getMaxJoltageFromBankPart2(intSlice []int, numNeed int) int {
	// If we have to grab 12 digits, then we just have to alter how much slice we expose
	if numNeed > len(intSlice) {
		panic("oh god u need more batteries than u have")
	}

	// Keeping the outnumber as a string for now will make it easier to append
	outNumber := ""
	// The bank has 15 batteries, so the first biggest number comes from the first 4. We can dynamically alter how many numbers we expose
	// based on how many numbers we have left
	maxLength := len(intSlice)
	nextIndex := 0
	allowance := maxLength - numNeed + 1

	// Basically in this case we have an "allowance" of 15-12+1 = 4, so on first iteration we can look at the first 4 numbers.
	// If we ever take the non 0th number from out considered slice, then our allowance is smaller..
	// Ex. If we took the 1st number, then for all subsequent searches, we can only consider a max of 3 numbers, otherwise we run out of batteries
	for i := 0; i < numNeed; i++ {
		consideredSlice := intSlice[nextIndex : nextIndex+allowance]
		max := slices.Max(consideredSlice)
		index := slices.Index(consideredSlice, max)

		nextIndex += index + 1
		allowance -= index

		outNumber += fmt.Sprintf("%d", max)
	}

	asInt, err := strconv.Atoi(outNumber)
	if err != nil {
		panic(err)
	}
	return asInt
}
