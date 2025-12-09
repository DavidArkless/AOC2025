package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Box struct {
	x int
	y int
	z int
}

type BoxDistance struct {
	aIndex   int
	bIndex   int
	distance float64
}

func main() {
	filePath := "day8/box.txt"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	junctionBoxes := make([]Box, 0)
	for scanner.Scan() {
		line := scanner.Text()

		coords := strings.Split(line, ",")
		box := Box{}
		box.x, _ = strconv.Atoi(coords[0])
		box.y, _ = strconv.Atoi(coords[1])
		box.z, _ = strconv.Atoi(coords[2])
		junctionBoxes = append(junctionBoxes, box)

	}

	distanceMap := make([]BoxDistance, 0)
	for i, box := range junctionBoxes {
		for j := i + 1; j < len(junctionBoxes); j++ {
			distance := getDistance(box, junctionBoxes[j])
			distanceMap = append(distanceMap, BoxDistance{i, j, distance})
		}
	}

	sort.Slice(distanceMap, func(i, j int) bool {
		return distanceMap[i].distance < distanceMap[j].distance
	})

	connections := make(map[int][]int)
	const numberOfPairs = 1000
	for i := 0; i < numberOfPairs; i++ {
		pair := distanceMap[i]
		connections[pair.aIndex] = append(connections[pair.aIndex], pair.bIndex)
		connections[pair.bIndex] = append(connections[pair.bIndex], pair.aIndex)
	}

	visited := make(map[int]bool)
	var circuitSizes []int

	for i := 0; i < len(junctionBoxes); i++ {

		if visited[i] {
			continue
		}
		currentCircuitSize := 0

		queue := []int{i}
		visited[i] = true

		for len(queue) > 0 {
			currentBoxId := queue[0]
			queue = queue[1:]

			currentCircuitSize++

			neighbours := connections[currentBoxId]
			for _, neighbourId := range neighbours {
				if !visited[neighbourId] {
					queue = append(queue, neighbourId)
					visited[neighbourId] = true
				}
			}
		}
		circuitSizes = append(circuitSizes, currentCircuitSize)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(circuitSizes)))

	if len(circuitSizes) >= 3 {
		result := circuitSizes[0] * circuitSizes[1] * circuitSizes[2]
		fmt.Println(result)
	}

	// Part2

	groupIDs := make([]int, len(junctionBoxes))
	for i := 0; i < len(junctionBoxes); i++ {
		groupIDs[i] = i
	}

	groupsRemaining := len(junctionBoxes)

	for _, pair := range distanceMap {

		idA := groupIDs[pair.aIndex]
		idB := groupIDs[pair.bIndex]

		if idA == idB {
			continue
		}

		for i := 0; i < len(groupIDs); i++ {
			if groupIDs[i] == idB {
				groupIDs[i] = idA
			}
		}

		groupsRemaining--

		if groupsRemaining == 1 {

			boxA := junctionBoxes[pair.aIndex]
			boxB := junctionBoxes[pair.bIndex]

			fmt.Println(boxA.x * boxB.x)
			break
		}
	}

}

func getDistance(box1, box2 Box) float64 {
	x := box1.x - box2.x
	y := box1.y - box2.y
	z := box1.z - box2.z

	distance := math.Sqrt(float64(x*x + y*y + z*z))
	return distance

}
