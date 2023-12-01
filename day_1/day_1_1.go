package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func calculateNumberFromString(input string) int {
	var firstSymbol rune
	var lastSymbol rune
	for _, symbol := range input {
		if unicode.IsDigit(symbol) {
			lastSymbol = symbol
			if firstSymbol == 0 {
				firstSymbol = symbol
			}
		}
	}
	outString := string(firstSymbol) + string(lastSymbol)

	number, err := strconv.Atoi(outString)
	if err != nil {
		fmt.Println("Error converting string to number:", err)
	}

	return number
}

func main() {
	file, err := os.Open("day_1_1_input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	// Close the file when the function exits
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	result := 0

	// Iterate through each line
	for scanner.Scan() {
		result += calculateNumberFromString(scanner.Text())

	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Result: %d\n", result)
}
