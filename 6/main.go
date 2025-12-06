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

	allNumbers := [][]int{}
	operations := []string{}

	part2Lines := []string{}

	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			continue
		}

		// Gonna pad the end with an extra whitespace because I use it as a signal to complete the computation
		part2Lines = append(part2Lines, line+" ")

		nums := []int{}

		lineSplit := strings.Split(line, " ")

		if lineSplit[0] == "*" || lineSplit[0] == "+" {
			for _, op := range lineSplit {
				if op != "" {
					operations = append(operations, op)
				}
			}
			continue
		}

		for _, numStr := range lineSplit {
			if numStr == "" {
				continue
			}
			asInt, err := strconv.Atoi(numStr)
			if err != nil {
				panic(err)
			}
			nums = append(nums, asInt)
		}
		allNumbers = append(allNumbers, nums)
	}

	total := 0
	for i := 0; i < len(allNumbers[0]); i++ {
		column := 0
		if operations[i] == "*" {
			column = 1
		}

		for _, nums := range allNumbers {
			if operations[i] == "+" {
				column += nums[i]
			}
			if operations[i] == "*" {
				column = column * nums[i]
			}
		}
		total += column
	}

	fmt.Println("Part 1", total)

	// Part 2
	// Its probably easier to work off the original string and move upwards, so lets just kind of start over
	// I needed to disable removing trailing whitespace in my editor for this...

	part2Total := 0
	total = 0
	var operator string
	// We can double for loop, iterate forward from the left (addition and multiplication are commutative) and then upwards, building the number as we go
	for x := 0; x < len(part2Lines[0]); x++ {
		numString := ""
		for y := len(part2Lines) - 1; y >= 0; y-- {
			if y == len(part2Lines)-1 {
				if string(part2Lines[y][x]) != " " {
					operator = string(part2Lines[y][x])
					if operator == "*" {
						total = 1
					}
				}
				continue
			}
			if string(part2Lines[y][x]) != " " {
				numString = string(part2Lines[y][x]) + numString
			}
		}
		if numString != "" {
			asNum, err := strconv.Atoi(numString)
			if err != nil {
				panic(err)
			}
			if operator == "*" {
				total = total * asNum
			} else {
				total += asNum
			}
		} else {
			part2Total += total
			total = 0
		}
	}

	fmt.Println("Part 2", part2Total)
}
