package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

// Token represents a lexical token
type Token int

var eof = rune(0)

func isDigit(ch rune) bool {
	return unicode.IsDigit(ch)
}

func main() {

	// Try the example
	fmt.Println("Trying example...")
	fmt.Println("\tWant: ", 142)
	runFile("./01/example.txt")

	fmt.Println("Trying part 1...")
	runFile("./01/input.txt")
}

func runFile(filePath string) {
	// Read the input file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	// Defer function to close the file and handle
	// any outstanding errors associated with opening the file.
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Create a scanner to read the file.
	scanner := bufio.NewScanner(file)

	// Iterate through the lines of the file.
	intSum := 0
	for scanner.Scan() {
		// Read the line of text
		line := scanner.Text()
		// Convert the line of text to an integer
		i, err := intsFromString(line)
		if err != nil {
			log.Fatal(err)
		}
		// Sum the integer
		intSum = intSum + i
	}

	// Print the final number
	fmt.Println("\tGot: ", intSum)
}

func intsFromString(text string) (int, error) {
	intStr := ""
	last := ""

	// Iterate through the string of text.
	for _, r := range text {
		// If the rune isn't a digit, continue
		if !unicode.IsDigit(r) {
			continue
		}
		// If it's the first one, save it
		if len(intStr) == 0 {
			intStr = string(r)
		}

		last = string(r)
	}

	// Concatenate the first and last digits
	intStr = intStr + last

	switch len(intStr) {
	case 0:
		return 0, nil
	default:
		return strconv.Atoi(intStr)
	}
}
