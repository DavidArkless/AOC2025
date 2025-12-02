package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type solutionData struct {
	minPosition int
	maxPosition int
	position    int
	count       int
}

func main() {
	filePath := "day1/rotations.txt"

	part1Solution := solutionData{
		minPosition: 0,
		maxPosition: 99,
		position:    50,
		count:       0,
	}

	part2Solution := solutionData{
		minPosition: 0,
		maxPosition: 99,
		position:    50,
		count:       0,
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		direction, amount := parseInstruction(scanner.Text())

		part1Solution.part1(direction, amount)
		part2Solution.part2(direction, amount)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1 -> position=%d, zero_hits=%d\n", part1Solution.position, part1Solution.count)
	fmt.Printf("Part 2 -> position=%d, zero_crossings=%d\n", part2Solution.position, part2Solution.count)

}

func parseInstruction(instruction string) (direction string, amount int) {
	s := strings.TrimSpace(instruction)
	if len(s) < 2 {
		log.Fatalf("invalid instruction: %q", instruction)
	}
	direction = s[:1]
	if direction != "R" && direction != "L" {
		log.Fatalf("invalid direction in instruction: %q", instruction)
	}
	var err error
	amount, err = strconv.Atoi(s[1:])
	if err != nil {
		log.Fatal(err)
	}
	return direction, amount

}

func (sol *solutionData) part1(direction string, amount int) {
	divisor := sol.maxPosition - sol.minPosition + 1
	if direction == "R" {
		sol.position += amount
	}
	if direction == "L" {
		sol.position -= amount
	}

	// Always normalize to keep invariant explicit
	sol.position = ((sol.position % divisor) + divisor) % divisor

	if sol.position == 0 {
		sol.count++
	}

}

func (sol *solutionData) part2(direction string, amount int) {
	divisor := sol.maxPosition - sol.minPosition + 1

	if direction == "R" {
		sol.count += (sol.position + amount) / divisor
		sol.position += amount
	}

	if direction == "L" {
		distanceToZero := sol.position
		if distanceToZero == 0 {
			distanceToZero = divisor
		}

		if amount >= distanceToZero {
			sol.count++
			sol.count += (amount - distanceToZero) / divisor
		}
		sol.position -= amount
	}

	sol.position = ((sol.position % divisor) + divisor) % divisor

}
