package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Dial int

func main() {
	input, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	dial := Dial(50)

	numTimesAtZero := 0

	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			continue
		}
		val := parseLineToValue(line)
		dial = dial.Move(val, &numTimesAtZero)
	}

	fmt.Println(numTimesAtZero)
}

func (d Dial) Move(val int, counter *int) Dial {
	current := int(d)

	// Discount the double counted left start at 0
	if current == 0 && val < 0 {
		*counter--
	}

	current += val
	for current > 99 {
		*counter++
		current -= 100
	}
	for current < 0 {
		*counter++
		current += 100
	}
	if current == 0 {
		*counter++
	}

	// Discounts landing on 0 going right
	if current == 0 && val > 0 {
		*counter--
	}

	return Dial(current)
}

func parseLineToValue(in string) int {
	isNegative := false
	if strings.HasPrefix(in, "L") {
		isNegative = true
	}

	asInt, err := strconv.Atoi(in[1:])
	if err != nil {
		panic(err)
	}

	if isNegative {
		asInt = -asInt
	}

	return asInt
}
