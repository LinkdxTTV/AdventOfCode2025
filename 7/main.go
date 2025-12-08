package main

import (
	"fmt"
	"os"
	"strings"
)

type point struct {
	x int
	y int
}

type graph []string

func main() {
	input, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	graph := graph{}

	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			continue
		}

		graph = append(graph, line)
	}

	maxWidth := len(graph[0])
	part1Splits := 0
	// We can go along the graph in the y direction downwards and populate changes to the row below before we get to it

	// Part 2: Lets keep track of the number of lines in every spot in the array
	thisRowParticles := make([]int, maxWidth)

	for y, row := range graph {
		nextRowParticles := make([]int, maxWidth)
		if y == len(graph)-1 {
			// Last row do nothing
			continue
		}
		for x, char := range row {
			pointBelow := point{x, y + 1}
			if string(char) == "S" {
				if graph.charAt(pointBelow) == "." {
					nextRowParticles[x] += 1
					graph.alterPoint(pointBelow, "|")
				}
			}
			if string(char) == "|" {
				nextRowParticles[x] += thisRowParticles[x]
				if graph.charAt(pointBelow) == "." {
					graph.alterPoint(pointBelow, "|")
				}
				if graph.charAt(pointBelow) == "^" {
					part1Splits++
					if x > 0 {
						pointBottomLeft := point{x - 1, y + 1}
						graph.alterPoint(pointBottomLeft, "|")
						nextRowParticles[x-1] += thisRowParticles[x]

					}
					if x < maxWidth-1 {
						pointBottomRight := point{x + 1, y + 1}
						graph.alterPoint(pointBottomRight, "|")
						nextRowParticles[x+1] += thisRowParticles[x]
					}
				}
			}
		}
		thisRowParticles = nextRowParticles
	}
	graph.print()
	fmt.Println("Part 1:", part1Splits)

	// Part 2, sum up our last row
	part2Sum := 0
	for _, particles := range thisRowParticles {
		part2Sum += particles
	}

	fmt.Println("Part 2:", part2Sum)
}

func (g graph) charAt(p point) string {
	return string(g[p.y][p.x])
}

func (g graph) alterPoint(p point, newChar string) {
	existingRow := g[p.y]
	newRow := existingRow[:p.x] + newChar + existingRow[p.x+1:]
	g[p.y] = newRow
}

func (g graph) print() {
	for _, row := range g {
		fmt.Println(row)
	}
}
