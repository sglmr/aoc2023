package main

import (
	"fmt"
	"github.com/sglmr/aoc2023/internal/read"
	"log"
	"strconv"
	"strings"
)

var digitsOne = map[string]int{
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
}

var digitsTwo = map[string]int{
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {

	// Read the example file
	example, err := read.ReadFile("./01/example.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Read the example file
	exampleTwo, err := read.ReadFile("./01/example2.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Read the input file
	input, err := read.ReadFile("./01/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Run example part 01
	fmt.Println("example part 1: ", processRows(example, digitsOne))

	// Run input part 01
	fmt.Println("part 01: ", processRows(input, digitsOne))

	// Run example part 02
	fmt.Println("example part 2: ", processRows(exampleTwo, digitsTwo))

	// Run input part 02
	fmt.Println("part 02: ", processRows(input, digitsTwo))

}

func processRows(data []string, digits map[string]int) int {
	intSum := 0
	for _, row := range data {
		mi := findMin(row, digits)
		ma := findMax(row, digits)

		// Concatenate min+max to get the value as a string for the row.
		stringNum := fmt.Sprintf("%v%v", mi, ma)

		// Convert the string number to an integer
		i, err := strconv.Atoi(stringNum)
		if err != nil {
			log.Fatal(err)
		}

		// Add the integers up
		intSum = intSum + i
	}

	return intSum
}

// Find the maximum digit in a text string
func findMax(text string, digits map[string]int) int {

	// Iterate backwards through the string to perform
	// a reverse index search on the string
	for i := len(text); i >= 0; i-- {
		t := text[i:]
		// Iterate over every digit to see if it's in the substring
		for k, v := range digits {
			l := strings.Index(t, k)
			// continue if l = -1
			if l < 0 {
				continue
			}
			// found a match
			return v
		}
	}
	return -1
}

// Find the minimum digit in a text string
func findMin(text string, digits map[string]int) int {

	for i := 0; i <= len(text); i++ {
		t := text[:i]

		// Iterate though each record in the digits map
		for k, v := range digits {
			l := strings.Index(t, k)

			// Try the next value if the value wasn't found
			if l == -1 {
				continue
			}
			return v
		}

	}

	return -1
}
