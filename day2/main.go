package main

import (
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
	filePath := "day2/ids.txt"
	solution := Answer{0, 0}
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	text := string(file)
	text = strings.TrimSpace(text)
	ids := strings.Split(text, ",")

	for _, id := range ids {
		startEnd := strings.Split(id, "-")
		start, err := strconv.Atoi(startEnd[0])
		if err != nil {
			log.Fatal(err)
		}
		end, err := strconv.Atoi(startEnd[1])
		if err != nil {
			log.Fatal(err)
		}

		for i := start; i <= end; i++ {
			solution.part1(i)
			solution.part2(i)

		}

	}

	fmt.Println(solution.part1Result)
	fmt.Println(solution.part2Result)

}

func (sol *Answer) part1(i int) {
	id := strconv.Itoa(i)
	if len(id)%2 == 1 {
		return
	}

	sliceLength := len(id) / 2
	if id[:sliceLength] == id[sliceLength:] {
		sol.part1Result += i
	}
}

func (sol *Answer) part2(i int) {
	id := strconv.Itoa(i)
	length := len(id)

	for chunkSize := 1; chunkSize < length; chunkSize++ {

		if length%chunkSize != 0 {
			continue
		}

		pattern := id[:chunkSize]

		repeats := length / chunkSize

		reconstructed := strings.Repeat(pattern, repeats)

		if reconstructed == id {
			sol.part2Result += i
			return
		}
	}

}

// I looked into some faster mathematics-based solutions so the below is not my own solution,
// and these are much better and have led to me understanding some interesting properties of the problem.
// Being able to create a repeated pattern from a number is possible using integer exponentiation.
// E.g. if you want to check if 1212 is a valid pair, you can divide it by 100 and check if the remainder is 12.
// There was another method as well, that is more for generating these numbers, but if you want to create 123123, for
// example, you could do this by doing 123 * 1001, so you could create a solution where we generate the possible combinations.

func (sol *Answer) part1Fast(i int) {
	digits := getDigitCount(i)

	// If an odd number of digits, it cannot be split evenly
	if digits%2 != 0 {
		return
	}

	// Calculate divisor to split the number in half
	// e.g. for 1212 (4 digits), we need 10^2 = 100
	halfDigits := digits / 2
	divisor := intPow10(halfDigits)

	left := i / divisor  // 1212 / 100 = 12
	right := i % divisor // 1212 % 100 = 12

	if left == right {
		sol.part1Result += i
	}
}

func (sol *Answer) part2Fast(i int) {
	digits := getDigitCount(i)

	// We only check up to half the length because a pattern must repeat at least twice to be valid.
	limit := digits / 2

	for chunkSize := 1; chunkSize <= limit; chunkSize++ {
		// If chunk doesn't divide evenly, skip it
		if digits%chunkSize != 0 {
			continue
		}

		// 1. Get the pattern from the top of the number
		// e.g. 123123 (6 digits), chunk 3. shift = 3.
		// 123123 / 1000 = 123
		shift := digits - chunkSize
		pattern := i / intPow10(shift)

		reconstructed := 0
		repeats := digits / chunkSize
		multiplier := intPow10(chunkSize)

		for k := 0; k < repeats; k++ {
			reconstructed = (reconstructed * multiplier) + pattern
		}

		if reconstructed == i {
			sol.part2Result += i
			return
		}
	}
}

func getDigitCount(n int) int {
	if n == 0 {
		return 1
	}
	count := 0
	for n > 0 {
		n /= 10
		count++
	}
	return count
}

func intPow10(n int) int {
	res := 1
	for i := 0; i < n; i++ {
		res *= 10
	}
	return res
}
