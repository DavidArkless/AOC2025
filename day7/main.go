package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Answer struct {
	part1Result int
	part2Result int
}

type Pos struct {
	row int
	col int
}

func main() {

	filePath := "day7/manifold.txt"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var grid [][]rune

	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	beams := make(map[int]struct{})

	for i := 0; i < len(grid[0]); i++ {
		if grid[0][i] == 'S' {
			beams[i] = struct{}{}
		}
	}

	answer := Answer{0, 0}

	for _, row := range grid {
		for j, cell := range row {
			if cell == '^' {
				if _, ok := beams[j]; ok {
					answer.part1(beams, j) // TODO: Boundary checks
				}
			}
		}
	}

	startCol := -1
	for c := 0; c < len(grid[0]); c++ {
		if grid[0][c] == 'S' {
			startCol = c
			break
		}
	}

	timelineCounts := make(map[int]int)
	timelineCounts[startCol] = 1

	totalTimelines := 0
	cols := len(grid[0])

	for _, row := range grid {
		nextRowCounts := make(map[int]int)

		for c, count := range timelineCounts {

			if c < 0 || c >= cols {

				totalTimelines += count
				continue
			}

			cell := row[c]

			if cell == '^' {

				if c-1 < 0 {
					totalTimelines += count
				} else {
					nextRowCounts[c-1] += count
				}

				if c+1 >= cols {
					totalTimelines += count
				} else {
					nextRowCounts[c+1] += count
				}
			} else {

				nextRowCounts[c] += count
			}
		}

		timelineCounts = nextRowCounts
	}

	for _, count := range timelineCounts {
		totalTimelines += count
	}

	answer.part2Result = totalTimelines

	fmt.Println(answer.part1Result)
	fmt.Println(answer.part2Result)

}

func (sol *Answer) part1(beams map[int]struct{}, col int) {
	delete(beams, col)
	sol.part1Result++
	beams[col+1] = struct{}{} // boundary checks
	beams[col-1] = struct{}{} // Boundary checks
}
