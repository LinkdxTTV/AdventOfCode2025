package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type point struct {
	x float64
	y float64
}

type pair struct {
	p1   *point
	p2   *point
	area int
}

type boundary struct {
	fixedDim    float64
	varyDimLow  float64
	varyDimHigh float64
}

func main() {
	input, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	allPoints := []point{}

	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			continue
		}

		coords := strings.Split(line, ",")
		x, err := strconv.Atoi(coords[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(coords[1])
		if err != nil {
			panic(err)
		}

		allPoints = append(allPoints, point{float64(x), float64(y)})
	}

	allPairs := []pair{}

	for i := 0; i < len(allPoints); i++ {
		for j := i + 1; j < len(allPoints); j++ {
			p1 := allPoints[i]
			p2 := allPoints[j]
			allPairs = append(allPairs, pair{
				p1:   &p1,
				p2:   &p2,
				area: int(calculateArea(p1, p2)),
			})
		}
	}

	sort.Slice(allPairs, func(i, j int) bool {
		return allPairs[i].area > allPairs[j].area
	})

	fmt.Println("Part 1: ", allPairs[0].p1, allPairs[0].p2, "Area", allPairs[0].area)

	// Part 2... The solution is somewhere inside allPairs, but we need to select the largest area that has all red/green tiles inside.
	// The largest rectangle completely within the loop can be traced without crossing a boundary..
	// Lets form some boundaries
	horizontalBoundaries := []boundary{}
	verticalBoundaries := []boundary{}

	for i := range len(allPoints) - 1 {
		p1 := allPoints[i]
		p2 := allPoints[i+1]

		addBoundaries(p1, p2, &horizontalBoundaries, &verticalBoundaries)
	}
	// Get the last one to complete the loop
	addBoundaries(allPoints[len(allPoints)-1], allPoints[0], &horizontalBoundaries, &verticalBoundaries)

	// fmt.Println(horizontalBoundaries, verticalBoundaries)

	// Iterate across the biggest areas and disqualify them one at a time
	for _, pair := range allPairs {
		horiz, vert := GenerateSquareBoundariesFromPointsButSlightlySmaller(*pair.p1, *pair.p2)
		cross := false
		for _, h := range horiz {
			for _, v := range verticalBoundaries {
				if DoBoundariesCross(h, v) {
					cross = true
					break
				}
			}
			if cross {
				break
			}
		}
		if cross {
			continue
		}
		for _, v := range vert {
			for _, h := range horizontalBoundaries {
				if DoBoundariesCross(h, v) {
					cross = true
					break
				}
			}
			if cross {
				break
			}
		}
		if cross {
			continue
		}
		fmt.Println("Part 2: ", *pair.p1, *pair.p2, pair.area)
		return
	}
}

func GenerateSquareBoundariesFromPointsButSlightlySmaller(p1, p2 point) ([]boundary, []boundary) {
	horizontal := []boundary{
		boundary{
			fixedDim:    p1.y,
			varyDimLow:  min(p1.x, p2.x),
			varyDimHigh: max(p1.x, p2.x),
		},
		boundary{
			fixedDim:    p2.y,
			varyDimLow:  min(p1.x, p2.x),
			varyDimHigh: max(p1.x, p2.x),
		},
	}

	vertical := []boundary{
		boundary{
			fixedDim:    p1.x,
			varyDimLow:  min(p1.y, p2.y),
			varyDimHigh: max(p1.y, p2.y),
		},
		boundary{
			fixedDim:    p2.x,
			varyDimLow:  min(p1.y, p2.y),
			varyDimHigh: max(p1.y, p2.y),
		},
	}

	// Shrink the boundaries ever so slightly
	if horizontal[0].fixedDim > horizontal[1].fixedDim {
		horizontal[0].fixedDim -= 0.1
		horizontal[1].fixedDim += 0.1
	} else {
		horizontal[0].fixedDim += 0.1
		horizontal[1].fixedDim -= 0.1
	}

	if vertical[0].fixedDim > vertical[1].fixedDim {
		vertical[0].fixedDim -= 0.1
		vertical[1].fixedDim += 0.1
	} else {
		vertical[0].fixedDim += 0.1
		vertical[1].fixedDim -= 0.1
	}

	horizontal[0].varyDimLow += 0.1
	horizontal[0].varyDimHigh -= 0.1
	horizontal[1].varyDimLow += 0.1
	horizontal[1].varyDimHigh -= 0.1
	vertical[0].varyDimLow += 0.1
	vertical[0].varyDimHigh -= 0.1
	vertical[1].varyDimLow += 0.1
	vertical[1].varyDimHigh -= 0.1

	return horizontal, vertical
}

func DoBoundariesCross(b1, b2 boundary) bool {
	return b1.fixedDim < b2.varyDimHigh && b1.fixedDim > b2.varyDimLow && b2.fixedDim < b1.varyDimHigh && b2.fixedDim > b1.varyDimLow
}

func addBoundaries(p1, p2 point, horizontalBoundaries, verticalBoundaries *[]boundary) {

	if p1.x == p2.x {
		*verticalBoundaries = append(*verticalBoundaries, boundary{
			fixedDim:    p1.x,
			varyDimLow:  min(p1.y, p2.y),
			varyDimHigh: max(p1.y, p2.y),
		})
	} else if p1.y == p2.y {
		*horizontalBoundaries = append(*horizontalBoundaries, boundary{
			fixedDim:    p1.y,
			varyDimLow:  min(p1.x, p2.x),
			varyDimHigh: max(p1.x, p2.x),
		})
	}
}

func abs(in float64) float64 {
	if in < 0 {
		return -in
	}
	return in
}

func calculateArea(p1, p2 point) float64 {
	return (abs(p1.x-p2.x) + 1) * (abs(p1.y-p2.y) + 1)
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b float64) float64 {
	if a < b {
		return b
	}
	return a
}
