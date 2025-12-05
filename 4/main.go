package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type vec struct {
	dx int
	dy int
}

type point struct {
	x int
	y int
}

func (p point) move(v vec) point {
	return point{
		x: p.x + v.dx,
		y: p.y + v.dy,
	}
}

var allDirections = []vec{{-1, 1}, {0, 1}, {1, 1}, {-1, 0}, {1, 0}, {-1, -1}, {0, -1}, {1, -1}}

type grid []string

func (g grid) CharAt(p point) (string, error) {
	height := len(g)
	width := len(g[0])

	if p.x < 0 || p.x >= width || p.y < 0 || p.y >= height {
		return "", fmt.Errorf("out of bounds")
	}

	return string(g[p.y][p.x]), nil
}

func main() {
	input, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	grid := grid{}
	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			continue
		}

		grid = append(grid, line)
	}

	accessiblePapers := Part2(grid)

	fmt.Print(accessiblePapers)
}

func Part1(grid grid) int {
	accessiblePapers := 0

	for y, row := range grid {
		for x, _ := range row {
			currentPoint := point{x: x, y: y}
			if thing, err := grid.CharAt(currentPoint); err == nil && thing == "@" {
				adjacentRolls := 0
				for _, v := range allDirections {
					if nearbyThing, err := grid.CharAt(currentPoint.move(v)); err == nil && nearbyThing == "@" {
						adjacentRolls++
					}
				}
				if adjacentRolls < 4 {
					accessiblePapers++
				}
			}
		}
	}
	return accessiblePapers
}

func Part2(grid grid) int {
	totalAccessiblePapers := 0

	accessiblePapersThisRound := -1

	for accessiblePapersThisRound != 0 {
		totalPaperRolls := 0
		gridCopy := grid.copy()

		accessiblePapersThisRound = 0

		for y, row := range grid {
			for x, _ := range row {
				currentPoint := point{x: x, y: y}
				if thing, err := grid.CharAt(currentPoint); err == nil && thing == "@" {
					totalPaperRolls++
					adjacentRolls := 0
					for _, v := range allDirections {
						if nearbyThing, err := grid.CharAt(currentPoint.move(v)); err == nil && nearbyThing == "@" {
							adjacentRolls++
						}
					}
					if adjacentRolls < 4 {
						accessiblePapersThisRound++
						gridCopy.replace(currentPoint, ".")
					}
				}
			}
		}

		fmt.Printf("Total Paper Rolls: %d \n", totalPaperRolls)
		totalAccessiblePapers += accessiblePapersThisRound
		grid.print()
		time.Sleep(80 * time.Millisecond)
		grid = gridCopy
	}

	return totalAccessiblePapers
}

func (g grid) copy() grid {
	out := grid{}
	for _, row := range g {
		out = append(out, row)
	}
	return out
}

func (g grid) replace(p point, replacement string) {
	g[p.y] = g[p.y][:p.x] + replacement + g[p.y][p.x+1:]
}

func (g grid) print() {
	for _, row := range g {
		fmt.Println(row)
	}
}
