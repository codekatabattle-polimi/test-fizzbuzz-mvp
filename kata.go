package main

import (
	"fmt"
	"log"
	"os"
)

func kataSolution(input string) string {
	// Write your code here
	return input
}

func main() {
	// Read the input from input.txt file
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Print the solution to standard output
	_, err = fmt.Print(kataSolution(string(file)))
	if err != nil {
		log.Fatal(err)
	}
}
