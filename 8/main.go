package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
	z int
}

type junctionPair struct {
	p1              *point
	p2              *point
	distanceSquared float64
}

const (
	connectionsToMakeSample = 10
	connectionsToMakePart1  = 1000
)

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
		z, err := strconv.Atoi(coords[2])
		if err != nil {
			panic(err)
		}
		allPoints = append(allPoints, point{x, y, z})
	}

	connectionsToMake := connectionsToMakeSample
	if len(allPoints) > 100 {
		connectionsToMake = connectionsToMakePart1
	}

	allPairs := []junctionPair{}

	// ok can we avoid brute forcing this.? Theres only 1000 choose 2 ~ 500k calculations..?
	for i := 0; i < len(allPoints); i++ {
		for j := i + 1; j < len(allPoints); j++ {
			p1 := allPoints[i]
			p2 := allPoints[j]
			allPairs = append(allPairs, junctionPair{
				p1:              &p1,
				p2:              &p2,
				distanceSquared: math.Pow(float64(p1.x-p2.x), 2) + math.Pow(float64(p1.y-p2.y), 2) + math.Pow(float64(p1.z-p2.z), 2),
			})
		}
	}

	sort.Slice(allPairs, func(i, j int) bool {
		return allPairs[i].distanceSquared < allPairs[j].distanceSquared
	})

	connections := []map[point]bool{}

	for i := range connectionsToMake {
		pair := allPairs[i]
		inConns := []int{}
		for i, conn := range connections {
			_, ok := conn[*pair.p1]
			if ok {
				inConns = append(inConns, i)
				conn[*pair.p2] = true
				continue
			}
			_, ok = conn[*pair.p2]
			if ok {
				inConns = append(inConns, i)
				conn[*pair.p1] = true
				continue
			}
		}
		// Wasnt in any existing conn, make a new one
		if len(inConns) == 0 {
			connections = append(connections, map[point]bool{
				*pair.p1: true,
				*pair.p2: true,
			})
		}
		if len(inConns) == 2 {
			// Need to combine two larger conns
			superConn := map[point]bool{}
			for _, idx := range inConns {
				for k, v := range connections[idx] {
					superConn[k] = v
				}
			}
			newConns := []map[point]bool{}
			for i, conn := range connections {
				if i != inConns[0] && i != inConns[1] {
					newConns = append(newConns, conn)
				}
			}
			newConns = append(newConns, superConn)
			connections = newConns
		}
	}

	// Then sort connections by sizes
	sort.Slice(connections, func(i, j int) bool {
		return len(connections[i]) > len(connections[j])
	})

	product := 1
	// Multiply 3 largest circuits
	for i := range 3 {
		product = product * len(connections[i])
	}

	fmt.Println("Part 1: ", product)

	// Part 2... // Quick copy and paste but in an infinite for loop
	i := 0
	for {
		pair := allPairs[i]
		inConns := []int{}
		for i, conn := range connections {
			_, ok := conn[*pair.p1]
			if ok {
				inConns = append(inConns, i)
				conn[*pair.p2] = true
				continue
			}
			_, ok = conn[*pair.p2]
			if ok {
				inConns = append(inConns, i)
				conn[*pair.p1] = true
				continue
			}
		}
		// Wasnt in any existing conn, make a new one
		if len(inConns) == 0 {
			connections = append(connections, map[point]bool{
				*pair.p1: true,
				*pair.p2: true,
			})
		}
		if len(inConns) == 2 {
			// Need to combine two larger conns
			superConn := map[point]bool{}
			for _, idx := range inConns {
				for k, v := range connections[idx] {
					superConn[k] = v
				}
			}
			newConns := []map[point]bool{}
			for i, conn := range connections {
				if i != inConns[0] && i != inConns[1] {
					newConns = append(newConns, conn)
				}
			}
			newConns = append(newConns, superConn)
			connections = newConns
		}
		if len(connections) == 1 && len(connections[0]) == len(allPoints) {
			fmt.Printf("All junctions connected after %d connections \n", i+1)
			fmt.Println("Part 2: ", pair.p1.x*pair.p2.x)
			return
		}
		i++
	}
}
