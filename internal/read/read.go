package read

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// Reads an input file and returns a slice of strings from the file contents
func ReadFile(filePath string) ([]string, error) {
	var fileData []string

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
	for scanner.Scan() {
		row := scanner.Text()
		fileData = append(fileData, strings.TrimSpace(row))
	}

	return fileData, nil
}
