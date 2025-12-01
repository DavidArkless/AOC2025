package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	filePath := "day1/rotations.txt"
	part1Pos := 50
	part2Pos := 50
	maxPosition := 99
	minPosition := 0
	var part1Count int
	var part2Count int

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		direction, amount := parseInstruction(scanner.Text())

		part1Pos, part1Count = part1(direction, amount, part1Pos, maxPosition, minPosition, part1Count)
		part2Pos, part2Count = part2(direction, amount, part2Pos, maxPosition, minPosition, part2Count)

	}

	fmt.Println(part1Pos, part1Count)
	fmt.Println(part2Pos, part2Count)

}

func parseInstruction(instruction string) (direction string, amount int) {
	direction = instruction[:1]
	amount, err := strconv.Atoi(instruction[1:])
	if err != nil {
		log.Fatal(err)
	}
	return direction, amount

}

func part1(direction string, amount int, position int, maxPosition int, minPosition int, count int) (int, int) {
	divisor := maxPosition - minPosition + 1
	if direction == "R" {
		position += amount
	}
	if direction == "L" {
		position -= amount
	}

	if position > maxPosition || position < minPosition {
		position = ((position % divisor) + divisor) % divisor

	}

	if position == 0 {
		count++
	}

	return position, count

}

func part2(direction string, amount int, position int, maxPosition int, minPosition int, count int) (int, int) {
	divisor := maxPosition - minPosition + 1

	if direction == "R" {
		count += (position + amount) / divisor
		position += amount
	}

	if direction == "L" {
		distanceToZero := position
		if distanceToZero == 0 {
			distanceToZero = divisor
		}

		if amount >= distanceToZero {
			count++
			count += (amount - distanceToZero) / divisor
		}
		position -= amount
	}

	position = ((position % divisor) + divisor) % divisor

	return position, count
}
