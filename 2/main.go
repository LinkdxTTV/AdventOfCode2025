package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	sum := 0

	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			continue
		}

		for numRange := range strings.SplitSeq(line, ",") {
			parsed := parseRange(numRange)

			for i := parsed.low; i <= parsed.high; i++ {
				if !IdIsValidPart2(i) {
					sum += i
				}
			}
		}
	}

	fmt.Println(sum)
}

func IdIsValidPart1(num int) bool {
	numAsString := fmt.Sprintf("%d", num)
	// Odd lengthed numbers like 1, 101, 10101, are always valid
	if len(numAsString)%2 != 0 {
		return true
	}

	if numAsString[:len(numAsString)/2] == numAsString[len(numAsString)/2:] {
		return false
	}
	return true
}

func IdIsValidPart2(num int) bool {
	numAsString := fmt.Sprintf("%d", num)
	if len(numAsString) == 1 {
		return true
	}
	for i := 2; i <= len(numAsString); i++ {
		// Could not possibly fail if its not evenly divisible
		if len(numAsString)%i != 0 {
			continue
		}

		// Check every possibility
		repetition := true
		subLength := len(numAsString) / i
		for j := 0; j < i-1; j++ {
			if numAsString[j*subLength:(j+1)*subLength] != numAsString[(j+1)*subLength:(j+2)*subLength] {
				repetition = false
				break
			}
		}
		if repetition {
			fmt.Println(num)
			return false
		}
	}
	return true
}

type parsedRange struct {
	low  int
	high int
}

func parseRange(in string) parsedRange {
	split := strings.Split(in, "-")
	low, err := strconv.Atoi(split[0])
	if err != nil {
		panic(err)
	}
	high, err := strconv.Atoi(split[1])
	if err != nil {
		panic(err)
	}
	return parsedRange{low: low, high: high}
}
