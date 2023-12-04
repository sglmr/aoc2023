package main

import (
	"fmt"
	"github.com/sglmr/aoc2023/internal/read"
	"log"
	"strconv"
	"strings"
)

func main() {

	// Read the example file
	example, err := read.ReadFile("./02/example.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Read the input file
	// Read the example file
	input, err := read.ReadFile("./02/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Run example part 01
	fmt.Println("example part 1: ", processRows(example, 12, 13, 14))

	// Run part 01
	fmt.Println("Part 1: ", processRows(input, 12, 13, 14))

}

func processRows(data []string, red, green, blue int) int {
	var gameIDSum int

	for i, row := range data {
		// Parse out the game play from each row of the original data.
		gameIDSum += checkGame(row, i+1, red, green, blue)
	}

	return gameIDSum
}

func checkGame(row string, gameID, red, green, blue int) int {

	// Parse the game into it's individual setsSlice
	sets := strings.SplitAfter(row, ": ")[1]
	setsSlice := strings.Split(sets, "; ")

	// For each set, check its validity.
	for _, set := range setsSlice {
		// If it's not a valid set, we'll add up the game ID
		if !validSet(set, red, green, blue) {
			return 0
		}
	}

	return gameID
}

// Parse the set into each color's count and check its validity.
// Returns true for a valid set and false for an invalid set
func validSet(set string, red, green, blue int) bool {

	// Split the set on ", " to get the parts and loop over them.
	for _, numColor := range strings.Split(set, ", ") {
		numColor = strings.TrimSpace(numColor)
		countString := strings.Split(numColor, " ")[0]
		color := strings.Split(numColor, " ")[1]

		count, err := strconv.Atoi(countString)
		if err != nil {
			log.Fatal(err)
		}

		if color == "red" && count > red {
			return false
		}
		if color == "green" && count > green {
			return false
		}
		if color == "blue" && count > blue {
			return false
		}
	}

	return true
}
