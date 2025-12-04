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

	filePath := "day3/battery.txt"
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	solution := Answer{0, 0}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		solution.part1(scanner.Text())
		solution.part2(scanner.Text())

	}

	fmt.Println(solution.part1Result)
	fmt.Println(solution.part2Result)

}

func (sol *Answer) part1(line string) {

	line = strings.TrimSpace(line)
	largestFirstNumber := 0
	index := 0
	for i := 0; i < len(line)-1; i++ {

		num := int(line[i] - '0')
		if num == 9 {
			largestFirstNumber = 9
			index = i
			break
		}
		if num > largestFirstNumber {
			largestFirstNumber = num
			index = i
		}
	}

	secondLargest := 0
	for i := index + 1; i < len(line); i++ {
		num, err := strconv.Atoi(string(line[i]))
		if err != nil {
			log.Fatal(err)
		}
		if num > secondLargest {
			secondLargest = num
		}
	}

	sol.part1Result += largestFirstNumber*10 + secondLargest

}

func (sol *Answer) part2(line string) {
	line = strings.TrimSpace(line)
	currentIndex := -1

	resultString := ""

	for i := 0; i < 12; i++ {
		remainingNeeded := 11 - i
		searchLimit := len(line) - 1 - remainingNeeded

		bestDigit := -1
		bestIndex := -1
		for j := currentIndex + 1; j <= searchLimit; j++ {
			digit := int(line[j] - '0')
			if digit == 9 {
				bestDigit = 9
				bestIndex = j
				break
			}
			if digit > bestDigit {
				bestDigit = digit
				bestIndex = j
			}
		}
		currentIndex = bestIndex
		resultString += strconv.Itoa(bestDigit)
	}
	val, _ := strconv.ParseInt(resultString, 10, 64)
	sol.part2Result += int(val)
}
