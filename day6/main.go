package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Answer struct {
	part1Result int
	part2Result int
}

func main() {
	filePath := "day6/homework.txt"

	answer := Answer{0, 0}

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var rawGrid []string
	scanner := bufio.NewScanner(file)
	maxLineLength := 0

	for scanner.Scan() {
		line := scanner.Text()
		rawGrid = append(rawGrid, line)
		if len(line) > maxLineLength {
			maxLineLength = len(line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i, line := range rawGrid {
		rawGrid[i] = PadRight(line, " ", maxLineLength)
	}

	var part1Grid [][]string
	for _, line := range rawGrid {
		part1Grid = append(part1Grid, strings.Fields(line))
	}

	if len(part1Grid) > 0 {
		cols := len(part1Grid[0])
		for i := 0; i < cols; i++ {
			col := make([]string, 0)
			for _, row := range part1Grid {
				if i < len(row) {
					col = append(col, row[i])
				}
			}
			answer.part1Result += answer.calculatePart1(col)
		}
	}

	answer.part2Result = answer.solvePart2(rawGrid, maxLineLength)

	fmt.Println("Part 1:", answer.part1Result)
	fmt.Println("Part 2:", answer.part2Result)
}

func PadRight(str, pad string, length int) string {
	if len(str) >= length {
		return str
	}
	return str + strings.Repeat(pad, length-len(str))
}

func (sol *Answer) calculatePart1(calculations []string) int {
	if len(calculations) == 0 {
		return 0
	}
	lastIndex := len(calculations) - 1
	operation := calculations[lastIndex]
	var result int

	startIdx := 0
	if len(calculations) > 1 {
		val, _ := strconv.Atoi(calculations[0])
		result = val
		startIdx = 1
	}

	for i := startIdx; i < lastIndex; i++ {
		number, err := strconv.Atoi(calculations[i])
		if err != nil {
			continue
		}
		if operation == "+" {
			result += number
		}
		if operation == "*" {
			result *= number
		}
	}
	return result
}

func (sol *Answer) solvePart2(lines []string, maxWidth int) int {
	grandTotal := 0
	height := len(lines)

	var currentBlockNumbers []int
	var currentBlockOp string

	for x := 0; x < maxWidth; x++ {

		isSeparator := true
		for y := 0; y < height; y++ {
			if getChar(lines, x, y) != ' ' {
				isSeparator = false
				break
			}
		}

		if isSeparator {
			if len(currentBlockNumbers) > 0 {
				grandTotal += calculateBlock(currentBlockNumbers, currentBlockOp)
			}
			currentBlockNumbers = []int{}
			currentBlockOp = ""
			continue
		}

		numStr := ""

		for y := 0; y < height; y++ {
			char := getChar(lines, x, y)

			if y == height-1 {
				if char == '+' || char == '*' {
					currentBlockOp = string(char)
				}
			}

			if char >= '0' && char <= '9' {
				numStr += string(char)
			}
		}

		if numStr != "" {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				log.Fatal(err)
			}
			currentBlockNumbers = append(currentBlockNumbers, num)
		}
	}

	if len(currentBlockNumbers) > 0 {
		grandTotal += calculateBlock(currentBlockNumbers, currentBlockOp)
	}

	return grandTotal
}

func getChar(grid []string, x, y int) byte {
	if y < 0 || y >= len(grid) {
		return ' '
	}
	if x < 0 || x >= len(grid[y]) {
		return ' '
	}
	return grid[y][x]
}

func calculateBlock(nums []int, op string) int {
	if len(nums) == 0 {
		return 0
	}
	result := nums[0]

	for i := 1; i < len(nums); i++ {
		if op == "+" {
			result += nums[i]
		} else if op == "*" {
			result *= nums[i]
		}
	}
	return result
}
