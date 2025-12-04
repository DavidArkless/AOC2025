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

func main() {
	filePath := "day4/map.txt"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var grid [][]rune

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		paddedLine := "." + line + "."

		if len(grid) == 0 {
			topBorder := make([]rune, len(paddedLine))
			for i := range topBorder {
				topBorder[i] = '.'
			}
			grid = append(grid, topBorder)
		}

		grid = append(grid, []rune(paddedLine))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if len(grid) > 0 {
		width := len(grid[0])
		bottomBorder := make([]rune, width)
		for i := range bottomBorder {
			bottomBorder[i] = '.'
		}
		grid = append(grid, bottomBorder)
	}

	answer := Answer{0, 0}
	answer.part1(grid)
	answer.part2(grid)

	fmt.Println(answer.part1Result)
	fmt.Println(answer.part2Result)

}

func (sol *Answer) part1(grid [][]rune) {
	for i, row := range grid {
		for j, cell := range row {
			if cell != '@' {
				continue
			}
			numAdjacentRolls := getNumAdjacentRollCells(grid, i, j)

			if numAdjacentRolls < 4 {
				sol.part1Result++

			}
		}
	}

}

func (sol *Answer) part2(grid [][]rune) {

	for {
		change := false

		for i, row := range grid {
			for j, cell := range row {
				if cell != '@' {
					continue
				}
				numAdjacentRolls := getNumAdjacentRollCells(grid, i, j)

				if numAdjacentRolls < 4 {
					grid[i][j] = '.'
					sol.part2Result++
					change = true

				}
			}
		}

		if !change {
			break
		}
	}

}

func getNumAdjacentRollCells(grid [][]rune, row int, col int) int {
	offsets := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	count := 0
	for _, offset := range offsets {

		if grid[row+offset[0]][col+offset[1]] == '@' {
			count++
		}
	}

	return count
}
