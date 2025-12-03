package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			continue
		}

		fmt.Println(line)
	}
}
