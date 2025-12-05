package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

type PuzzleInput struct {
	ranges []Range
	ids    []int
}

type Answer struct {
	part1Result int
	part2Result int
}

func main() {
	filePath := "day5/inventory.txt"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	input := PuzzleInput{
		ranges: make([]Range, 0),
		ids:    make([]int, 0),
	}

	parsingRanges := true

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			parsingRanges = false
			continue
		}

		if parsingRanges {
			parts := strings.Split(line, "-")
			start, _ := strconv.Atoi(parts[0])
			end, _ := strconv.Atoi(parts[1])
			input.ranges = append(input.ranges, Range{start, end})
		} else {
			id, _ := strconv.Atoi(line)
			input.ids = append(input.ids, id)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	answer := Answer{0, 0}

	answer.part1(input)
	answer.part2(input)

	fmt.Println(answer.part1Result)
	fmt.Println(answer.part2Result)
}

func (sol *Answer) part1(data PuzzleInput) {

	for _, id := range data.ids {
		isFresh := false
		for _, r := range data.ranges {
			if id >= r.start && id <= r.end {
				isFresh = true
				break
			}
		}
		if isFresh {
			sol.part1Result++
		}
	}
}

func (sol *Answer) part2(data PuzzleInput) {

	ranges := make([]Range, len(data.ranges))
	copy(ranges, data.ranges)

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})

	if len(ranges) == 0 {
		return
	}

	currentStart := ranges[0].start
	currentEnd := ranges[0].end
	totalFresh := 0

	for i := 1; i < len(ranges); i++ {
		nextRange := ranges[i]

		if nextRange.start <= currentEnd+1 {
			if nextRange.end > currentEnd {
				currentEnd = nextRange.end
			}
		} else {
			totalFresh += currentEnd - currentStart + 1

			currentStart = nextRange.start
			currentEnd = nextRange.end
		}
	}

	totalFresh += currentEnd - currentStart + 1

	sol.part2Result = totalFresh
}
